package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthCtx struct {
	key string
	rt  *RuntimeCtx
}

type JWTContents struct {
	Username string `json:"username"`
}

type UserInfo struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// generates a new JWT token
func (ac AuthCtx) NewToken(c JWTContents) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": c.Username,
	})

	tokenStr, err := token.SignedString([]byte(ac.key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// retrieves and decodes JWT token
func (ac AuthCtx) ReadToken(tokenStr string) (JWTContents, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(ac.key), nil
	})
	if err != nil || !token.Valid {
		return JWTContents{}, err
	}

	return JWTContents{
		Username: claims["username"].(string),
	}, nil
}

func (ac AuthCtx) IsValidToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return []byte(ac.key), nil
	})
	if err != nil {
		return false
	}

	return token.Valid
}

func IsValidEmail(addr string) bool {
	_, err := mail.ParseAddress(addr)
	return err == nil
}

// registers a new user
func (ac AuthCtx) HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")

	var userInfo UserInfo

	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Failed to decode JSON"))
		return
	}
	defer r.Body.Close()

	rt := ac.rt
	row := rt.dbConn.QueryRow("SELECT username FROM users WHERE username = ?;", userInfo.Username)
	var username string
	row.Scan(&username)
	// don't bother continuing if the username is already in use
	if userInfo.Username == username {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Username is already in use."))
		return
	}

	// create a new user in the database
	if !IsValidEmail(userInfo.Email) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid email format"))
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to generate password hash"))
		eprintln(err)
		return
	}

	_, err = rt.dbConn.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?);", userInfo.Username, userInfo.Email, passHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database error"))
		eprintln(err)
		return
	}

	w.Write([]byte("User registered"))
}

// logs in an existing users and gives them a JWT token
func (ac AuthCtx) HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userInfo UserInfo

	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	rt := ac.rt
	res := rt.dbConn.QueryRow("SELECT email, username, password FROM users WHERE username = ?;", userInfo.Username)
	var username, email string
	var passHash []byte
	if res.Scan(&email, &username, &passHash) == sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database error"))
		eprintln(err)
		return
	}

	err = bcrypt.CompareHashAndPassword(passHash, []byte(userInfo.Password))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Incorrect username or password"))
		return
	}

	token, err := ac.NewToken(JWTContents{
		Username: username,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to generate new token"))
		return
	}

	payload := UserInfo{
		Email:    email,
		Username: username,
	}
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	cookie := http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: time.Now().Add(364 * 24 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
	w.Write(response)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "auth_token",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, &cookie)
}

func (ac AuthCtx) IsAuthenticated(r *http.Request) (JWTContents, error) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return JWTContents{}, err
	}

	if err = cookie.Valid(); err != nil {
		return JWTContents{}, err
	}

	if len(cookie.Value) == 0 {
		return JWTContents{}, errors.New("no cookie has been set yet")
	}

	authInfo, err := ac.ReadToken(cookie.Value)
	if err != nil {
		return JWTContents{}, err
	}

	return authInfo, nil
}

// prevents non authenticated users from accessing endpoint
func (ac AuthCtx) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := ac.IsAuthenticated(r); err == nil {
			next.ServeHTTP(w, r)
		}
	})
}
