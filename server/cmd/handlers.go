package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"api/internal/models"
	"api/internal/validator"
)

type blogCreateForm struct {
	Title   string
	Content string
	validator.Validator
}

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
		return
	}

	flash := app.sessionManager.PopString(r.Context(), "flash")

	fmt.Fprintf(w, "%+v", blog, flash)
}

func (app *application) blogCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new blog"))
}

func (app *application) blogCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	form := blogCreateForm{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")

	if !form.Valid() {
		fmt.Fprint(w, form.FieldErrors)
		return
	}

	id, err := app.blogs.Insert(form.Title, form.Content)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Blog successfully created")

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
