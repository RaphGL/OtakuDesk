package main

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/raphgl/otakudesk/middleware"
	_ "modernc.org/sqlite"
)

//go:embed tables.sql
var schemaSQL string

type RuntimeCtx struct {
	dbConn *sql.DB
}

func eprintln(err any) {
	fmt.Fprintln(os.Stderr, "Error:", err)
}

func initRuntime() (RuntimeCtx, error) {
	var rtErr error

	const PATH_ENV = "OTAKUDESK_PATH"
	rtPath := os.Getenv(PATH_ENV)
	if rtPath == "" {
		pathErr := errors.New("No lookup path has been set in " + PATH_ENV)
		rtErr = errors.Join(rtErr, pathErr)
	}

	db, err := sql.Open("sqlite", path.Join(rtPath, "otakudesk.sqlite"))
	if err != nil {
		rtErr = errors.Join(rtErr, err)
	}

	if err = db.Ping(); err != nil {
		rtErr = errors.Join(rtErr, err)
	} else {
		fmt.Println("Established connection with Sqlite database.")
	}

	_, err = db.Exec(schemaSQL)
	if err != nil {
		rtErr = errors.Join(rtErr, err)
	}

	return RuntimeCtx{
		dbConn: db,
	}, rtErr
}

func (rt *RuntimeCtx) destroy() {
	rt.dbConn.Close()
}

func main() {
	mux := http.NewServeMux()
	rt, err := initRuntime()
	if err != nil {
		eprintln(err)
		return
	}
	defer rt.destroy()

	auth := AuthCtx{
		key: "test",
		rt:  &rt,
	}

	defaultMw := middleware.New(middleware.CORS)
	// authorizedMw := defaultMw.Extend(auth.AuthMiddleware)

	mux.Handle("POST /login", defaultMw.ThenFunc(auth.HandleLogin))
	mux.Handle("POST /register", defaultMw.ThenFunc(auth.HandleRegister))
	mux.Handle("POST /logout", defaultMw.ThenFunc(HandleLogout))

	// endpoint pinged to check is user is authenticated
	mux.Handle("GET /is-auth", defaultMw.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		userInfo, err := auth.IsAuthenticated(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("User is not authenticated."))
			return
		}
		defer r.Body.Close()

		userJSON, err := json.Marshal(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(userJSON)
	}))

	http.ListenAndServe(":8080", mux)
}
