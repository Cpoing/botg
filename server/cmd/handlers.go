package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Home"))
}

func (app *application) blogView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific blog with ID %d...", id)
	w.Write([]byte(msg))
}

func (app *application) blogCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new blog"))
}

func (app *application) blogCreatePost(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new blog..."))
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
