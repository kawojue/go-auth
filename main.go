package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gologger "github.com/kawojue/gin-gologger"
	"github.com/kawojue/go-auth/db"
	"github.com/kawojue/go-auth/routes"
	initenv "github.com/kawojue/go-initenv"
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	initenv.LoadEnv()
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	PORT := initenv.GetEnv("PORT", "8080")

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))

	router.Use(gologger.Logger(gin.ReleaseMode))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Go Auth",
		})
	})

	routes.AuthRoutes(router)

	router.Run(fmt.Sprintf(":%s", PORT))
}
