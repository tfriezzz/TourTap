package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tfriezzz/tourtap/internal/auth"
	"github.com/tfriezzz/tourtap/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hi from login")
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User         User   `json:"user"`
		Token        string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil || !match {
		respondWithError(w, http.StatusUnauthorized, "incorrect email or password", err)
		return
	}

	JWTToken, err := auth.MakeJWT(
		user.ID,
		cfg.jwtSecret,
		time.Hour,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create access JWT", err)
		return
	}

	refreshToken := auth.MakeRefreshToken()
	if _, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	}); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not save refresh token", err)
		return
	}

	sseJWT, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create SSE-JWT", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "ssh_auth",
		Value:    sseJWT,
		Path:     "/events",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
			Name:      user.Name,
		},
		Token:        JWTToken,
		RefreshToken: refreshToken,
	})
}
