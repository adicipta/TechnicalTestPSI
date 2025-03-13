package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	SessionExp int64  `json:"session_exp"`
	jwt.RegisteredClaims
}
