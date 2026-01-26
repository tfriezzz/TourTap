package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerToursGet(w http.ResponseWriter, r *http.Request) {
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
