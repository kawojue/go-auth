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
		err            error
		existingAvatar models.Avatars
		user           models.Users
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
		panic(err)
	}
	defer avatar.Close()

	file, err := utils.HandleFile(ctx, 5<<20, header, avatar, "jpg", "png")
	if err != nil {
		panic(err)
	}

	helpers.UploadS3(ctx, file.FileName, file.FileBytes)
	avatar_url := helpers.GetS3(file.FileName)

	if err = configs.DB.Where("user_id = ?", user.ID.String()).First(&existingAvatar).Error; err != nil {
		avatarRecord := models.Avatars{
			Url:    avatar_url,
			Path:   file.FileName,
			Type:   file.FileExtension,
			UserId: user.ID,
		}

		if err = configs.DB.Create(&avatarRecord).Error; err != nil {
			helpers.SOMETHING_WENT_WRONG(ctx)
			return
		}
	} else {
		if err = configs.DB.Where("user_id = ?", user.ID.String()).Updates(&models.Avatars{
			Url:  avatar_url,
			Path: file.FileName,
			Type: file.FileExtension,
		}).Error; err != nil {
			helpers.SOMETHING_WENT_WRONG(ctx)
			return
		}
	}

	helpers.SendSuccess(ctx, http.StatusOK, "", map[string]string{
		"path": file.FileName,
		"url":  avatar_url,
	})
}

func DeleteAvatar(ctx *gin.Context) {

}
