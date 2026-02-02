// Package auth
package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tfriezzz/tourtap/internal/database"
)

func HashPassword(password string) (string, error) {
	params := argon2id.DefaultParams
	hash, err := argon2id.CreateHash(password, params)
	// fmt.Printf("from auth: %v\n", hash)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	now := jwt.NewNumericDate(time.Now().UTC())
	expiration := jwt.NewNumericDate(now.Add(expiresIn))
	jwtID := uuid.New().String()
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  now,
		ExpiresAt: expiration,
		Subject:   userID.String(),
		ID:        jwtID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := jwt.RegisteredClaims{}

	keyFunc := func(*jwt.Token) (any, error) {
		returnToken := []byte(tokenSecret)

		return returnToken, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)
	if err != nil {
		return uuid.UUID{}, err
	}

	strID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != string(TokenTypeAccess) {
		return uuid.Nil, errors.New("invalid user")
	}

	id, err := uuid.Parse(strID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid user ID: %w", err)
	}

	return id, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	header := headers["Authorization"]
	for _, headerString := range header {
		splitString := strings.Split(headerString, " ")
		tokenString := strings.TrimSpace(splitString[1])
		return tokenString, nil
	}
	return "", fmt.Errorf("missing string")
}

func MakeRefreshToken() string {
	token := make([]byte, 32)
	rand.Read(token)
	return hex.EncodeToString(token)
}

func GetAPIKey(headers http.Header) (string, error) {
	header := headers["Authorization"]
	for _, string := range header {
		splitString := strings.Split(string, " ")
		apiString := strings.TrimSpace(splitString[1])
		return apiString, nil
	}
	return "", fmt.Errorf("missing string")
}

// func RJWTokenMaker(JWTString string, user database.User) (string, string, error) {
// 	JWTtoken, err := MakeJWT(user.ID, JWTString, time.Hour)
// 	if err != nil {
// 		return "", "", err
// 	}

	refreshToken, err := MakeRefreshToken()
	if err != nil {
		return "", "", err
	}
	return JWTtoken, refreshToken, nil
}

func RefreshTokenToDatabase(cfg *apiConfig, r *http.Request, refreshToken string, user uuid.UUID) error {
	returnTokenParams := database.CreateRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	_, err := cfg.DB.CreateRefreshToken(r.Context(), returnTokenParams)
	if err != nil {
		return err
	}
	return nil
}

func DBUserToUserResponse(u database.User, cfg *apiConfig, JWTtoken string, refreshToken string) userResponse {
	return userResponse{
		ID:           u.ID,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.CreatedAt,
		Email:        u.Email,
		Token:        JWTtoken,
		RefreshToken: refreshToken,
		IsChirpyRed:  u.IsChirpyRed,
	}
}
