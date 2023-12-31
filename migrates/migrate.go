package main

import (
	"github.com/kawojue/go-auth/db"
	"github.com/kawojue/go-auth/models"
	initenv "github.com/kawojue/init-env"
)

func init() {
	initenv.LoadEnv("../.env")
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.Users{})
}
