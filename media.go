package main

import (
	"net/http"
)

type MediaCtx struct {
	rt *RuntimeCtx
}

type MediaType int

const (
	MediaTypeAnime = iota
	MediaTypeManga
)

type ListMediaResp struct{}

// reads all media from lookupPath/{anime,manga} and stores it into database for faster queries
func (mc MediaCtx) cacheMediaListToDB() {}

// lists the anime/manga the user has in lookupPath
func (mc MediaCtx) ListMedia(mtyp MediaType) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

type MediaInfoResp struct{}

// retrieve media info from external database and store it locally for faster lookup
func (mc MediaCtx) cacheMediaInfoToDB(mtyp MediaType, name string) {}

// returns info about manga either from cache or from external API like anilist or myanimelist
func (MediaCtx) GetMediaInfo(mtyp MediaType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
