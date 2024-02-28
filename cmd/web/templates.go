package main

import (
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"pythononline.usmanbala.com/internal/models"
	"pythononline.usmanbala.com/ui"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type templateData struct {
	Progress        *models.Progress
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	Url             string
}

func humanDate(t time.Time) string {
	// Return the empty string if time has the zero value.
	if t.IsZero() {
		return ""
	}
	// Convert the time to UTC before formatting it.
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	pages := []string{}

	err := fs.WalkDir(ui.Files, "html/pages", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".html") {
			pages = append(pages, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	// Loop through the page filepaths one-by-one.
	for _, page := range pages {

		name := filepath.Base(page)
		// Parse the base template file into a template set.

		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map as normal...
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}
