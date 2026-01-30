// Package templates handles rendering/layout
package templates

import (
	"embed"
	"log"
	"net/http"
	"text/template"
)

//go:embed *.html
var templatesFS embed.FS

var tmpl *template.Template

func Load() error {
	tmpl = template.Must(template.ParseFS(templatesFS, "*.html"))
	log.Printf("Parsed templates: %v", tmpl.DefinedTemplates())
	return nil
}

func RenderTemplate(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.Write([]byte(`
// <!DOCTYPE html>
// <html>
// <head><title>Test</title></head>
// <body><h1>hi from server</h1><p>can you see me?</p></body>
// </html>
// `))
// }
