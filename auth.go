package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthCtx struct {
	key string
	rt  *RuntimeCtx
}

type UserInfo struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

func (ac AuthCtx) NewToken(username string, passwordHash string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": passwordHash,
	})

	tokenStr, err := token.SignedString([]byte(ac.key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (ac AuthCtx) ReadToken(tokenStr string) (UserInfo, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(ac.key), nil
	})
	if err != nil || !token.Valid {
		return UserInfo{}, err
	}

	return UserInfo{
		Username: claims["username"].(string),
		Password: claims["password"].(string),
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

func (ac AuthCtx) HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")

	var userInfo UserInfo

	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	if !IsValidEmail(userInfo.Email) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid email format"))
		return
	}

	rt := ac.rt
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

	token, err := ac.NewToken(username, string(passHash))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to generate new token"))
		return
	}

	payload := UserInfo{
		Email:    email,
		Username: username,
		Password: string(passHash),
		Token:    token,
	}
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(response)
}
