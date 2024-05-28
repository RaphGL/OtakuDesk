// todo: add logger
package main

import (
	"database/sql"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"

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

func initRuntime(lookupPath string) (RuntimeCtx, error) {
	var rtErr error

	const PATH_ENV = "OTAKUDESK_PATH"

	db, err := sql.Open("sqlite", path.Join(lookupPath, "otakudesk.sqlite"))
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
	// --- flag parsing
	homeDir, err := os.UserHomeDir()
	if err != nil {
		eprintln(err)
		return
	}

	lookupPath := flag.String("path", filepath.Join(homeDir, "OtakuDesk"), "Lookup path for content delivered by the server")
	flag.Parse()

	// --- http server configuration
	mux := http.NewServeMux()
	rt, err := initRuntime(*lookupPath)
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

	mux.Handle("POST /login", defaultMw.ThenFunc(auth.HandleLogin))
	mux.Handle("POST /register", defaultMw.ThenFunc(auth.HandleRegister))
	mux.Handle("POST /logout", defaultMw.ThenFunc(HandleLogout))
	// endpoint pinged to check is user is authenticated
	mux.Handle("GET /is-auth", defaultMw.ThenFunc(auth.CheckSessionValidity))

	// authorizedMw := defaultMw.Extend(auth.AuthMiddleware)

	http.ListenAndServe(":8080", mux)
}
