package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", app.home)
	mux.HandleFunc("/blog/view/{id}", app.blogView)
	mux.HandleFunc("/blog/create", app.blogCreate)

	return mux
}
