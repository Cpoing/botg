package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /blog/view/{id}", app.blogView)
	mux.HandleFunc("GET /blog/create", app.blogCreate)
	mux.HandleFunc("POST /blog/create", app.blogCreatePost)

	return mux
}
