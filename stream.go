package main

import "net/http"

type StreamCtx struct {
	rt *RuntimeCtx
}

func (StreamCtx) StreamAnime(w http.ResponseWriter, r *http.Request) {

}

func (StreamCtx) StreamManga(w http.ResponseWriter, r *http.Request) {

}
