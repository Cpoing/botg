package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /blog/view/{id}", dynamic.ThenFunc(app.blogView))
	mux.Handle("GET /blog/create", dynamic.ThenFunc(app.blogCreate))
	mux.Handle("POST /blog/create", dynamic.ThenFunc(app.blogCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
