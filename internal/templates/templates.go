// Package templates handles rendering/layout
package templates

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

func Load() error {
	tmpl = template.Must(template.ParseGlob("internal/templates/*.html"))
	return nil
}

func Render(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}
