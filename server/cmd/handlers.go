package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"api/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	blogs, err := app.blogs.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, blog := range blogs {
		fmt.Fprintf(w, "%+v\n", blog)
	}
}

func (app *application) blogView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	blog, err := app.blogs.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
	}

	fmt.Fprintf(w, "%+v", blog)
}

func (app *application) blogCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new blog"))
}

func (app *application) blogCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "blog 1"
	content := "Testing blog 1"

	id, err := app.blogs.Insert(title, content)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/view/%d", id), http.StatusSeeOther)
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
