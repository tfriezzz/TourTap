package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tfriezzz/tourtap/internal/auth"
	"github.com/tfriezzz/tourtap/internal/database"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := User{}
	if err := decoder.Decode(&params); err != nil {
		fmt.Print(err)
	}
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Print(err)
	}
	userParams := database.CreateUserParams{
		Email: params.Email, HashedPassword: hashedPassword,
	}
	user, err := cfg.db.CreateUser(r.Context(), userParams)
	if err != nil {
		fmt.Printf("CreateUser call: %v\n", err)
	}

	JWTtoken, refreshToken, err := auth.RJWTokenMaker(cfg, user)
	if err != nil {
		log.Printf("RJWTokenMaker returned err: %v", err)
	}

	if err := auth.RefreshTokenToDatabase(cfg, r, refreshToken, user.ID); err != nil {
		log.Printf("refreshTokenToDatabase returned err: %v", err)
	}

	userResponse := auth.DBUserToUserResponse(user, cfg, JWTtoken, refreshToken)

	if err := respondWithJSON(w, 201, userResponse); err != nil {
		fmt.Print(err)
	}
}
