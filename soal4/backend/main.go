package main

import (
	"soal4/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", handlers.GetUser)
	r.Run(":8080")
}
