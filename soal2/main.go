package main

import (
	"log"
	"soal2/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	r := gin.Default()
	routes.RegisterRouter(r)
	r.Run(":8080")
}
