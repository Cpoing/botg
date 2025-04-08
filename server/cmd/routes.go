package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /blog/view/{id}", app.blogView)
	mux.HandleFunc("GET /blog/create", app.blogCreate)
	mux.HandleFunc("POST /blog/create", app.blogCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
