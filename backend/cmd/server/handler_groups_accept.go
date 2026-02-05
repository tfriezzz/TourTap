package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/tfriezzz/tourtap/internal/auth"
	"github.com/tfriezzz/tourtap/internal/pubsub"
)

func (cfg *apiConfig) handlerGroupsAccept(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	if _, err := auth.ValidateJWT(token, cfg.jwtSecret); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	groupsIDString := r.PathValue("groupID")
	groupID, err := uuid.Parse(groupsIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid group ID", err)
		return
	}

	group, err := cfg.db.GroupStatusAccepted(r.Context(), groupID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not accept group", err)
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

	if err := pubsub.Publish("group_accepted", payload); err != nil {
		log.Printf("could not publish group_accepted: %v", err)
	}

	respondWithJSON(w, http.StatusOK, payload)

	// log.Printf("tour accepted, sending email with payment details to: %s\n", group.Email)
}
