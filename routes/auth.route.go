package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/controllers"
)

var auth *gin.RouterGroup

func AuthRoutes(router *gin.RouterGroup) {
	auth = router.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", controllers.Logout)
		auth.POST("/refresh", controllers.RefreshToken)
	}
}
