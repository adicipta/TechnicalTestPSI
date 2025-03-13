package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soal4/models"

	"github.com/gin-gonic/gin"
)

func fetchRandomUser(results, page int) (*models.RandomUserResponse, error) {
	url := fmt.Sprintf("https://randomuser.me/api?results=%d&page=%d", results, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userResponse models.RandomUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return nil, err
	}
	return &userResponse, nil
}

func GetUser(c *gin.Context) {
	results := 10
	page := 1
	userResponse, err := fetchRandomUser(results, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(userResponse.Results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user := userResponse.Results[0]
	formattedUser := models.UserResponse{
		Name:     fmt.Sprintf("%s, %s %s", user.Name.Title, user.Name.First, user.Name.Last),
		Location: fmt.Sprintf("%d, %s, %s, %s, %s", user.Location.Street.Number, user.Location.Street.Name, user.Location.City, user.Location.State, user.Location.Country),
		Email:    user.Email,
		Age:      user.Dob.Age,
		Phone:    user.Phone,
		Cell:     user.Cell,
		Picture:  []string{user.Picture.Large, user.Picture.Medium, user.Picture.Thumb},
	}

	c.JSON(http.StatusOK, formattedUser)
}
