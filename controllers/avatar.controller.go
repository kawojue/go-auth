package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	"github.com/kawojue/go-auth/utils"
)

func UploadAvatar(ctx *gin.Context) {
	var (
		err  error
		existingAvatar models.Avatars
		user models.Users
	)

	user_id, exists := ctx.Get("user_id")
	if !exists {
		helpers.SOMETHING_WENT_WRONG(ctx)
		return
	}

	if err = configs.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		helpers.ACCOUNT_NOT_FOUND(ctx)
		return
	}

	avatar, header, err := ctx.Request.FormFile("avatar")
	if err != nil {
		return
	}
	defer avatar.Close()

	file, err := utils.HandleFile(ctx, 5<<20, header, avatar, "jpg", "png")
	if err != nil {
		return
	}

	helpers.UploadS3(ctx, file.FileName, file.FileBytes)
	avatar_url := helpers.GetS3(file.FileName)

	if err = configs.DB.
}

func DeleteAvatar(ctx *gin.Context) {

}
