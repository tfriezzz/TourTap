package main

import "net/http"

func (cfg *apiConfig) handlerGroupsReset(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Msg string `json:"msg"`
	}

	if err := cfg.db.DeleteAllGroups(r.Context()); err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot reset groups: %v", err)
	}

	respondWithJSON(w, http.StatusAccepted, response{Msg: "group successfully reset"})
}
