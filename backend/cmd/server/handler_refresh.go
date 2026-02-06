package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tfriezzz/tourtap/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	log.Print("hi from refresh")
	type response struct {
		Token string `json:"access_token"`
	}

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could not find token", err)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could not get user from refresh token", err)
		return
	}

	JWTToken, err := auth.MakeJWT(
		user.ID,
		cfg.jwtSecret,
		time.Hour,
	)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could not validate token", err)
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		Token: JWTToken,
	})
}
