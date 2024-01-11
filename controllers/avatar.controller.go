package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
)

func UploadAvatar(ctx *gin.Context) {
	var (
		err  error
		user models.Users
	)

	user_id, exists := ctx.Get("user_id")
	if !exists {
		helpers.SendError(ctx, http.StatusUnauthorized, "Something went wrong.")
		return
	}

	if err = configs.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Account not found.")
		return
	}

	// file, handler, err := ctx.Request.FormFile("avatar")

}

func DeleteAvatar(ctx *gin.Context) {

}
