package routes

import (
	"soal2/controllers"
	"soal2/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/login", controllers.LoginHandler)
	r.GET("/protected", utils.AuthMiddleware(), controllers.ProtectedHandler)
}
