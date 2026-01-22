package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tfriezzz/tourtap/internal/database"
)

type Customer struct {
	ID        uuid.UUID               `json:"id"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	Email     string                  `json:"email"`
	Name      string                  `json:"name"`
	Status    database.CustomerStatus `json:"customer_status"`
}

func (cfg *apiConfig) handlerCustomersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}

	newCustomerParams := database.CreateCustomerParams{
		Email: params.Email,
		Name:  params.Name,
	}

	customer, err := cfg.db.CreateCustomer(r.Context(), newCustomerParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create customer", err)
		log.Printf("could not create customer: %v\n", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Customer{
		ID:        customer.ID,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
		Email:     customer.Email,
		Name:      customer.Name,
		Status:    customer.Status,
	})
}
