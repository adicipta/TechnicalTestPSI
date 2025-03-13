package routes

import (
	"soal1/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/checkout", controllers.CheckoutHandler)
	return r
}
