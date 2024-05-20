package main

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	_ "modernc.org/sqlite"
	"net/http"
	"os"
	"path"
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

	mux.HandleFunc("POST /login", auth.HandleLogin)
	mux.HandleFunc("POST /register", auth.HandleRegister)

	http.ListenAndServe(":8080", mux)
}
