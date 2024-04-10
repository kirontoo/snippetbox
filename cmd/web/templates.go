package main

import (
	"text/template"
	"path/filepath"

	"github.com/kirontoo/snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get a slice of all filepaths that match the pattern 
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		// parse files into a template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// add template set to the map, using the name of the page as the key
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}

	return cache, nil
}
