package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/kirontoo/snippetbox/internal/models"
)

// Home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if the current URL path does not exactly match "/" then it is not the home page
	// return a 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.notFound(w)
	}

	err = ts.ExecuteTemplate(w, "base", snippets)
	if err != nil {
		app.serverError(w, err)
	}
}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", snippet)
	if err != nil {
		app.serverError(w, err)
	}
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// If it's not a POST method, send status code 405 and meesage
		// "Method Not Allowed"
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
