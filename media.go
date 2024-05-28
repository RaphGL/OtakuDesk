package main

import "net/http"

type MediaCtx struct {
	rt *RuntimeCtx
}

type MediaType int

const (
	MediaTypeAnime = iota
	MediaTypeManga
)

func (MediaCtx) ListMedia(mtyp MediaType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (MediaCtx) GetMediaInfo(mtyp MediaType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
