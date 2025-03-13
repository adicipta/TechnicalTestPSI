package utils

import (
	"os"
	"soal2/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(username string) (string, time.Time, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	id := uuid.New().String()
	expTime := time.Now().Add(60 * time.Second)

	claims := models.Claims{
		ID:         id,
		Username:   username,
		SessionExp: expTime.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, expTime, err
}
