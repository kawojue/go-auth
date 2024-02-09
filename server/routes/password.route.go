package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/controllers"
)

var password *gin.RouterGroup

func PasswordRoutes(router *gin.RouterGroup) {
	password = router.Group("/password")

	{
		password.POST("/req-otp", controllers.ForgotPassword)
		password.POST("/reset", controllers.ResetPasword)
	}
}
