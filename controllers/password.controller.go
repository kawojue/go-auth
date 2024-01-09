package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-auth/utils"
)

func ForgotPassword(ctx *gin.Context) {
	var (
		err  error
		user *models.Users
		body structs.ForgotPassword
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Invalid JSON.")
		return
	}

	email := strings.ToLower(strings.TrimSpace(body.Email))
	if err = configs.DB.Where("email = ?", email).First(&user).Error; err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Account not found.")
		return
	}

	totp := utils.GenOTP(5)

	totpRecord := models.TOTP{
		Otp:       totp.Otp,
		OtpExpiry: totp.Otp_Expiry,
		UserID:    user.ID,
	}

	if err = configs.DB.Create(&totpRecord).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed to save OTP in the database.")
		return
	}

	helpers.SendMail([]string{user.Email}, fmt.Sprintf("Subject: One-time passcode\r\n\r\nOTP: %s", totp.Otp))

	helpers.SendSuccess(ctx, http.StatusOK, "An OTP has been sent to your email.", nil)
}
