package utils

import (
	"net/http"
	"os"
	"soal2/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		if len(jwtSecret) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret is missing"})
			c.Abort()
			return
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.SessionExp {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		c.Set("user", gin.H{
			"id":          claims.ID,
			"username":    claims.Username,
			"session_exp": claims.SessionExp,
		})

		c.Next()
	}
}
