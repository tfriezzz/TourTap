package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/tfriezzz/tourtap/internal/pubsub"
)

func (cfg *apiConfig) handlerPaymentWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			GroupID uuid.UUID `json:"group_id"`
		}
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}

	if params.Event != "group.paid" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	group, err := cfg.db.GroupStatusConfirmed(r.Context(), params.Data.GroupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "could not find group", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "could not update group", err)
		return
	}

	payload := Group{
		ID:              group.ID,
		CreatedAt:       group.CreatedAt,
		UpdatedAt:       group.UpdatedAt,
		Email:           group.Email,
		Name:            group.Name,
		Pax:             group.Pax,
		Status:          group.Status,
		RequestedTourID: group.RequestedTourID,
		RequestedDate:   group.RequestedDate,
		BookingID:       group.BookingID,
	}

	if err := pubsub.Publish("group_confirmed", payload); err != nil {
		respondWithError(w, http.StatusInternalServerError, "coult not publish group_confirmed", err)
	}

	respondWithJSON(w, http.StatusOK, payload)
}
