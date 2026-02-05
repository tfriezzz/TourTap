package main

import (
	"encoding/json"
	"net/http"

	"github.com/tfriezzz/tourtap/internal/auth"
	"github.com/tfriezzz/tourtap/internal/database"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not hash password", err)
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		HashedPassword: hashedPassword,
		Email:          params.Email,
		Name:           params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create user", err)
		return
	}

	// JWTtoken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	// if err != nil {
	// 	log.Printf("RJWTokenMaker returned err: %v", err)
	// }

	// if err := auth.RefreshTokenToDatabase(cfg, r, refreshToken, user.ID); err != nil {
	// 	log.Printf("refreshTokenToDatabase returned err: %v", err)
	// }

	// userResponse := auth.DBUserToUserResponse(user, cfg, JWTtoken, refreshToken)

	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
			Name:      user.Name,
		},
	})
}
