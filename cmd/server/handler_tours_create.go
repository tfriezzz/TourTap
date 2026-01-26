package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tfriezzz/tourtap/internal/database"
)

type Tour struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BasePrice string    `json:"base_price"`
}

func (cfg *apiConfig) handlerToursCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name      string `json:"name"`
		BasePrice string `json:"base_price"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}

	newTourParams := database.CreateTourParams{
		Name:      params.Name,
		BasePrice: params.BasePrice,
	}

	tour, err := cfg.db.CreateTour(r.Context(), newTourParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create tour", err)
	}

	respondWithJSON(w, http.StatusCreated, Tour{
		ID:        tour.ID,
		Name:      tour.Name,
		CreatedAt: tour.CreatedAt,
		UpdatedAt: tour.UpdatedAt,
		BasePrice: tour.BasePrice,
	})
}
