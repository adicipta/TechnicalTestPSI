package controllers

import (
	"net/http"
	"soal2/utils"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	token, exp, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't generate token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  exp,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successful",
		"token":       token,
		"session_exp": exp.Unix(),
	})
}

func ProtectedHandler(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "Access granted",
		"user":    user,
	})
}
