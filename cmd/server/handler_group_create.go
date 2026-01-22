package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tfriezzz/tourtap/internal/database"
)

type Group struct {
	ID              uuid.UUID            `json:"id"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	Email           string               `json:"email"`
	Name            string               `json:"name"`
	Pax             int32                `json:"pax"`
	Status          database.GroupStatus `json:"customer_status"`
	RequestedTourID int32                `json:"requested_tour_id"`
	RequestedDate   time.Time            `json:"requested_date"`
}

func (cfg *apiConfig) handlerGroupCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email           string    `json:"email"`
		Name            string    `json:"name"`
		Pax             int32     `json:"pax"`
		RequestedTourID int32     `json:"requested_tour_id"`
		RequestedDate   time.Time `json:"requested_date"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}

	newGroupParams := database.CreateCustomerParams{
		Email:           params.Email,
		Name:            params.Name,
		Pax:             params.Pax,
		RequestedTourID: params.RequestedTourID,
		RequestedDate:   params.RequestedDate,
	}

	group, err := cfg.db.CreateCustomer(r.Context(), newGroupParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create group", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Group{
		ID:              group.ID,
		CreatedAt:       group.CreatedAt,
		UpdatedAt:       group.UpdatedAt,
		Email:           group.Email,
		Name:            group.Name,
		Pax:             group.Pax,
		Status:          group.Status,
		RequestedTourID: group.RequestedTourID,
		RequestedDate:   group.RequestedDate,
	})
}
