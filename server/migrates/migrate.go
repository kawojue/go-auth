package main

import (
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/models"
	initenv "github.com/kawojue/go-initenv"
)

func init() {
	initenv.LoadEnv("../.env")
	configs.ConnectDB()
}

func main() {
	if err := configs.DB.AutoMigrate(&models.Users{}, &models.TOTP{}, &models.Profiles{}, &models.Avatars{}); err != nil {
		panic(err)
	}
}
