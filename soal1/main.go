package main

import (
	"soal1/routes"
)

func main() {
	r := routes.Router()
	r.Run(":8080")
}
