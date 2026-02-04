package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerFrontend(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}
	type response struct {
		Message string `json:"message"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		log.Printf("frontend_test: %v", r.Body)
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}
	log.Printf("hi from backend: %v", params.Email)

	respondWithJSON(w, http.StatusCreated, response{
		Message: fmt.Sprintf("hi from backend: %v", params.Email),
	})
}
