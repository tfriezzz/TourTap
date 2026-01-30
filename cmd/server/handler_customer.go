package main

import (
	"net/http"

	"github.com/tfriezzz/tourtap/internal/database"
	"github.com/tfriezzz/tourtap/internal/templates"
)

func (cfg *apiConfig) handlerCustomer(w http.ResponseWriter, r *http.Request) {
	tours, err := cfg.db.GetAllTours(r.Context())
	if err != nil {
		templates.RenderTemplate(w, "base.html", struct{ Error string }{Error: "failed to load tours: " + err.Error()})
		return
	}

	message := r.URL.Query().Get("message")
	errorMsg := r.URL.Query().Get("error")

	data := struct {
		Tours   []database.Tour // Your Tour model
		Message string
		Error   string
	}{
		Tours:   tours,
		Message: message,
		Error:   errorMsg,
	}

	templates.RenderTemplate(w, "base.html", data)
}
