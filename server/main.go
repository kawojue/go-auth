package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gologger "github.com/kawojue/gin-gologger"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/controllers"
	"github.com/kawojue/go-auth/routes"
	initenv "github.com/kawojue/go-initenv"
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	initenv.LoadEnv()
	configs.ConnectDB()
}

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 7 << 20

	PORT := initenv.GetEnv("PORT", "8080")

	router.Use(gologger.Logger(gin.ReleaseMode))

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     cors.DefaultConfig().AllowHeaders,
	}))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Go Auth",
		})
	})

	routes.AuthRoutes(router.Group("/"))
	routes.FileRoutes(router.Group("/"))
	routes.PasswordRoutes(router.Group("/auth"))
	router.GET("/:username", controllers.GetUser)

	router.Run(fmt.Sprintf(":%s", PORT))
}
