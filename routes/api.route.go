package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/controllers"
)

func ApiRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/signup", controllers.SignUp)
	}
}
