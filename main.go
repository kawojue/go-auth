package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gologger "github.com/kawojue/gin-gologger"
	"github.com/kawojue/go-auth/db"
	initenv "github.com/kawojue/init-env"
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
		AllowMethods:     []string{"GET", "POST"},
	}))

	router.Use(gologger.Logger("dev"))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Go Auth",
		})
	})

	router.Run(fmt.Sprintf(":%s", PORT))
}
