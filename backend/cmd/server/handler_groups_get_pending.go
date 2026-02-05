package main

import "net/http"

func (cfg *apiConfig) handlerGroupsGetPending(w http.ResponseWriter, r *http.Request) {
	groups, err := cfg.db.GetGroupsPending(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get groups", err)
		return
	}

	payload := make([]Group, len(groups))

	for i, group := range groups {
		payload[i] = Group{
			ID:        group.ID,
			CreatedAt: group.CreatedAt,
			UpdatedAt: group.UpdatedAt,
			Email:     group.Email,
			Name:      group.Name,
			Pax:       group.Pax,
			Status:    group.Status,
		}
	}

	respondWithJSON(w, http.StatusOK, payload)
}
