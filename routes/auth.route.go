package routes

import (
	"github.com/gin-gonic/gin"
	rate "github.com/kawojue/gin-ratelimiter"
	"github.com/kawojue/go-auth/controllers"
)

func AuthRoutes(router *gin.Engine) {
	limiterConfig := rate.LimiterConfig{
		MaxAttempts: 5,
		TimerArray:  []int{15, 14, 19},
		Message:     "Too many requests",
	}

	limiter := rate.CreateLimiter(&limiterConfig)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST(
			"/login",
			rate.RateLimiter(limiter, &limiterConfig),
			controllers.Login,
		)
		auth.POST("/logout", controllers.Logout)
		auth.POST("/refresh", controllers.RefreshToken)
	}
}
