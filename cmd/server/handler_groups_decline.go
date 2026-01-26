package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGroupsDecline(w http.ResponseWriter, r *http.Request) {
	groupsIDString := r.PathValue("groupID")
	groupID, err := uuid.Parse(groupsIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid group ID", err)
		return
	}

	group, err := cfg.db.GroupStatusDeclined(r.Context(), groupID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decline group", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Group{
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
	})

	log.Printf("tour declined, sending email to: %s\n", group.Email)
}
