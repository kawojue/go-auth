package routes

import (
	"github.com/gin-gonic/gin"
	rate "github.com/kawojue/gin-ratelimiter"
	"github.com/kawojue/go-auth/controllers"
)

var auth *gin.RouterGroup

func AuthRoutes(router *gin.RouterGroup) {
	limiterConfig := rate.LimiterConfig{
		MaxAttempts: 5,
		TimerArray:  []int{10, 8, 20},
		Message:     "Too many request.",
	}

	limiter := rate.CreateLimiter(&limiterConfig)

	auth = router.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST(
			"/login",
			rate.RateLimiter(limiter, &limiterConfig),
			controllers.Login,
		)
		auth.POST("/logout", controllers.Logout)
		auth.POST("/refresh", controllers.RefreshToken)
		auth.POST("/req-otp", controllers.ForgotPassword)
	}
}
