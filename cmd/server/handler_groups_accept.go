package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGroupsAccept(w http.ResponseWriter, r *http.Request) {
	groupsIDString := r.PathValue("groupID")
	groupID, err := uuid.Parse(groupsIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid group ID", err)
		return
	}

	group, err := cfg.db.GroupStatusAccepted(r.Context(), groupID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not accept group", err)
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
}
