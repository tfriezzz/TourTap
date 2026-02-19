package main

import (
	"net/http"

	"github.com/tfriezzz/tourtap/internal/auth"
)

func (cfg *apiConfig) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
			return
		}

		if _, err := auth.ValidateJWT(token, cfg.jwtSecret); err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
