package controllers

import (
	"fmt"
	"net/http"

	// "path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	// "github.com/kawojue/go-auth/utils"
)

func UploadAvatar(ctx *gin.Context) {
	var (
		err error
		// existingAvatar models.Avatars
		// body           structs.AvatarForm
		user models.Users
	)

	// if err = ctx.ShouldBind(&body); err != nil {
	// 	helpers.INVALID_JSON(ctx)
	// 	return
	// }

	user_id, exists := ctx.Get("user_id")
	if !exists {
		helpers.SOMETHING_WENT_WRONG(ctx)
		return
	}

	if err = configs.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		helpers.ACCOUNT_NOT_FOUND(ctx)
		return
	}

	if err := ctx.Request.ParseMultipartForm(5 << 20); err != nil {
		helpers.SendError(ctx, http.StatusRequestEntityTooLarge, "File too large")
		return
	}

	// header := body.Avatar

	form, err := ctx.MultipartForm()
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(form)

	// files := form.File["files"]

	// for _, file := range files {
	// 	fileName := filepath.Base(file.Filename)

	// 	avatar, err := utils.HandleFile(ctx, 5<<20, file.Header, file, "jpg", "png")
	// 	if err != nil {
	// 		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
	// 		return
	// 	}

	// }

	// helpers.UploadS3(ctx, avatar.FileName, avatar.FileBytes)
	// avatar_url := helpers.GetS3(avatar.FileName)

	// if err = configs.DB.Where("user_id = ?", user.ID.String()).First(&existingAvatar).Error; err != nil {
	// 	avatarRecord := models.Avatars{
	// 		Url:    avatar_url,
	// 		Path:   avatar.FileName,
	// 		Type:   avatar.FileExtension,
	// 		UserId: user.ID,
	// 	}

	// 	if err = configs.DB.Create(&avatarRecord).Error; err != nil {
	// 		helpers.SOMETHING_WENT_WRONG(ctx)
	// 		return
	// 	}
	// } else {
	// 	if err = configs.DB.Where("user_id = ?", user.ID.String()).Updates(&models.Avatars{
	// 		Url:  avatar_url,
	// 		Path: avatar.FileName,
	// 		Type: avatar.FileExtension,
	// 	}).Error; err != nil {
	// 		helpers.SOMETHING_WENT_WRONG(ctx)
	// 		return
	// 	}
	// }

	helpers.SendSuccess(ctx, http.StatusOK, "", map[string]string{
		// "path": avatar.FileName,
		// "url":  avatar_url,
	})
}

func DeleteAvatar(ctx *gin.Context) {

}
