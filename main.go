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
	"path/filepath"

	"github.com/raphgl/otakudesk/middleware"
	_ "modernc.org/sqlite"
)

//go:embed tables.sql
var schemaSQL string

type RuntimeCtx struct {
	LookupPath string
	AnimePath  string
	MangaPath  string
	DB         *sql.DB
}

func eprintln(err any) {
	fmt.Fprintln(os.Stderr, "Error:", err)
}

func initRuntime(lookupPath string) (RuntimeCtx, error) {
	var rtErr error

	const PATH_ENV = "OTAKUDESK_PATH"

	db, err := sql.Open("sqlite", filepath.Join(lookupPath, "otakudesk.sqlite"))
	if err != nil {
		rtErr = errors.Join(rtErr, err)
	}
	db.SetMaxOpenConns(1)

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
		LookupPath: lookupPath,
		AnimePath:  filepath.Join(lookupPath, "anime"),
		MangaPath:  filepath.Join(lookupPath, "manga"),
		DB:         db,
	}, rtErr
}

func (rt *RuntimeCtx) destroy() {
	rt.DB.Close()
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
	media := MediaCtx{
		rt: &rt,
	}

	defaultMw := middleware.New(middleware.CORS)

	mux.Handle("POST /login", defaultMw.ThenFunc(auth.HandleLogin))
	mux.Handle("POST /register", defaultMw.ThenFunc(auth.HandleRegister))
	mux.Handle("POST /logout", defaultMw.ThenFunc(HandleLogout))
	// endpoint pinged to check is user is authenticated
	mux.Handle("GET /is-auth", defaultMw.ThenFunc(auth.CheckSessionValidity))
	mux.Handle("GET /animes", defaultMw.Then(media.ListMedia(MediaTypeAnime)))
	mux.Handle("GET /mangas", defaultMw.Then(media.ListMedia(MediaTypeManga)))

	// authorizedMw := defaultMw.Extend(auth.AuthMiddleware)

	// todo: recache if cache paths in database don't start with $OTAKUDESK_PATH
	if err := media.cacheMediaListToDB(); err != nil {
		eprintln(err)
	}

	fmt.Println("Running server on localhost:8080")
	http.ListenAndServe(":8080", mux)
}
