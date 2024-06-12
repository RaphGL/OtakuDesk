package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type MediaCtx struct {
	rt *RuntimeCtx
}

type MediaType int

const (
	MediaTypeAnime = iota
	MediaTypeManga
)

// reads all media from lookupPath/{anime,manga} and stores it into database for faster queries
func (mc MediaCtx) cacheMediaListToDB() error {
	rt := mc.rt

	// todo: stop duplicating cache in db every time function runs
	cacheMedia := func(mtyp MediaType, ch chan<- error) {
		var query, mediaPath string

		switch mtyp {
		case MediaTypeAnime:
			query = "INSERT OR IGNORE INTO anime (name, path) VALUES (?, ?);"
			mediaPath = rt.AnimePath
		case MediaTypeManga:
			query = "INSERT OR IGNORE INTO manga (name, path) VALUES (?, ?);"
			mediaPath = rt.MangaPath
		}

		dir, err := os.ReadDir(mediaPath)
		if err != nil {
			ch <- err
			return
		}

		for _, content := range dir {
			contentPath := filepath.Join(mediaPath, content.Name())
			contentPath = strings.TrimPrefix(contentPath, rt.LookupPath)

			_, err := rt.DB.Exec(query, content.Name(), contentPath)
			if err != nil {
				ch <- err
				return
			}
		}

		ch <- nil
		return
	}

	medChan := make(chan error, 2)
	go cacheMedia(MediaTypeAnime, medChan)
	go cacheMedia(MediaTypeManga, medChan)

	var err error
	for i := 0; i < 2; i++ {
		err = errors.Join(err, <-medChan)
	}

	fmt.Println("Finished caching media")
	return err
}

// retrieve media info from external database and store it locally for faster lookup
func (mc MediaCtx) cacheMediaInfoToDB(mtyp MediaType, name string) {}

type ListMediaResp struct {
	Name string `json:"name"`
	Path string `json:"path"`

	// used for anime

	Episode       int `json:"episode,omitempty"`
	CurrEpisode   int `json:"curr_episode,omitempty"`
	TotalEpisodes int `json:"total_episodes,omitempty"`

	// used for manga

	Chapter       int `json:"chapter,omitempty"`
	CurrChapter   int `json:"curr_chapter,omitempty"`
	TotalChapters int `json:"total_chapters,omitempty"`
}

// lists the anime/manga the user has in lookupPath
func (mc MediaCtx) ListMedia(mtyp MediaType) http.Handler {
	rt := mc.rt

	var query string
	switch mtyp {
	case MediaTypeAnime:
		query = "SELECT name, path FROM anime;"
	case MediaTypeManga:
		query = "SELECT name, path FROM manga"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := rt.DB.Query(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer rows.Close()

		var mediaList []ListMediaResp
		for rows.Next() {
			var row ListMediaResp
			if err := rows.Scan(&row.Name, &row.Path); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			mediaList = append(mediaList, row)
		}

		resp, err := json.Marshal(mediaList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(resp)
	})
}

type MediaInfoResp struct{}

// returns info about anime/manga either from cache or from external API like anilist or myanimelist
func (MediaCtx) GetMediaInfo(mtyp MediaType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
