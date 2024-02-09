package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-auth/utils"
	gobcrypt "github.com/kawojue/go-bcrypt"
)

func ForgotPassword(ctx *gin.Context) {
	var (
		err          error
		user         models.Users
		existingTotp models.TOTP
		body         structs.ForgotPassword
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		helpers.INVALID_JSON(ctx)
		return
	}

	email := strings.ToLower(strings.TrimSpace(body.Email))
	if err = configs.DB.Where("email = ?", email).First(&user).Error; err != nil {
		helpers.ACCOUNT_NOT_FOUND(ctx)
		return
	}

	totp := utils.GenOTP(6)

	if err = configs.DB.Where("user_id", user.ID.String()).First(&existingTotp).Error; err != nil {
		totpRecord := models.TOTP{
			Otp:       totp.Otp,
			OtpExpiry: totp.Otp_Expiry,
			UserID:    user.ID,
		}
		if err = configs.DB.Create(&totpRecord).Error; err != nil {
			helpers.FAILED_TO_SAVE_OTP(ctx)
			return
		}
	} else {
		otp_exp := existingTotp.OtpExpiry
		if otp_exp != "" {
			parsedOTPExp, err := time.Parse("2006-01-02T15:04:05.999Z", existingTotp.OtpExpiry)
			if err != nil {
				panic(err)
			}

			remainingTime := parsedOTPExp.Sub(time.Now().UTC())

			if remainingTime > 0 {
				remainingMinutes := int(remainingTime.Minutes())

				helpers.SendError(ctx, http.StatusUnauthorized, fmt.Sprintf("Request after %d minutues.", remainingMinutes))
				return
			}
		}
		if err = configs.DB.Model(&existingTotp).Where("user_id", user.ID.String()).Updates(&models.TOTP{
			Otp:       totp.Otp,
			OtpExpiry: totp.Otp_Expiry,
		}).Error; err != nil {
			helpers.FAILED_TO_SAVE_OTP(ctx)
			return
		}
	}

	helpers.SendMail([]string{user.Email}, fmt.Sprintf("Subject: One-time passcode\r\n\r\nOTP: %s", totp.Otp))

	helpers.SendSuccess(ctx, http.StatusOK, "An OTP has been sent to your email.", nil)
}

func ResetPasword(ctx *gin.Context) {
	var (
		err          error
		user         models.Users
		existingTotp models.TOTP
		body         structs.ResetPassword
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		helpers.INVALID_JSON(ctx)
		return
	}

	otp, pswd, pswd2 := body.Otp, body.Pswd, body.Pswd2

	if otp == "" || pswd == "" || pswd2 == "" {
		helpers.ALL_FIELDS_REQUIRED(ctx)
		return
	}

	if pswd != pswd2 {
		helpers.SendError(ctx, http.StatusBadRequest, "Passwords not match")
		return
	}

	if len(pswd) < 8 {
		helpers.SendError(ctx, http.StatusBadRequest, "Password is too short.")
		return
	}

	if err = configs.DB.Where("otp = ?", otp).First(&existingTotp).Error; err != nil {
		helpers.SendError(ctx, http.StatusUnauthorized, "Incorrect OTP.")
		return
	}

	otp_expiry := existingTotp.OtpExpiry

	parsedOTPExp, err := time.Parse("2006-01-02T15:04:05.999Z", otp_expiry)
	if err != nil {
		helpers.SOMETHING_WENT_WRONG(ctx)
		return
	}

	otp_expired := parsedOTPExp.Before(time.Now().UTC())
	if otp_expired {
		configs.DB.Model(&models.TOTP{}).Where("otp = ?", otp).Update("otp", "")
		helpers.SendError(ctx, http.StatusForbidden, "OTP has expired.")
		return
	}

	if err = configs.DB.Where("id = ?", existingTotp.UserID.String()).First(&user).Error; err != nil {
		helpers.ACCOUNT_NOT_FOUND(ctx)
		return
	}

	hashPswd := gobcrypt.HashPassword(pswd)

	if err = configs.DB.Model(&user).Where("id = ?", user.ID.String()).Update("password", hashPswd).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error updating password.")
		return
	}

	updatedTOTP := map[string]interface{}{
		"Otp":       "",
		"OtpExpiry": "",
	}
	if err = configs.DB.Model(&existingTotp).Where("user_id = ?", user.ID.String()).Updates(updatedTOTP).Error; err != nil {
		helpers.SOMETHING_WENT_WRONG(ctx)
		return
	}

	helpers.SendSuccess(ctx, http.StatusOK, "You have successfully changed your password.", nil)
}
