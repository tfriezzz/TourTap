package main

import (
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerToursGet(w http.ResponseWriter, r *http.Request) {
	type Tour struct {
		ID        int32     `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		BasePrice string    `json:"base_price"`
	}

	tours, err := cfg.db.GetAllTours(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get tours", err)
		return
	}

	response := make([]Tour, len(tours))
	for i, tour := range tours {
		response[i] = Tour{
			ID:        tour.ID,
			Name:      tour.Name,
			CreatedAt: tour.CreatedAt,
			UpdatedAt: tour.UpdatedAt,
			BasePrice: tour.BasePrice,
		}
	}

	respondWithJSON(w, http.StatusOK, response)
}
