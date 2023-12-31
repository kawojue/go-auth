package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-auth/utils"
	gobcrypt "github.com/kawojue/go-bcrypt"
	"gorm.io/gorm"
)

func SignUp(ctx *gin.Context) {
	var (
		err  error
		body structs.SignUp
		user models.Users
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Invalid request body.")
		return
	}

	password := body.Password
	email := strings.TrimSpace(body.Email)
	username := strings.ToLower(strings.TrimSpace(body.Username))

	if len(password) == 0 || len(email) == 0 || len(username) == 0 {
		helpers.SendError(ctx, http.StatusBadRequest, "All fields are required.")
		return
	}

	if len([]byte(password)) > 72 || len([]byte(password)) < 8 {
		helpers.SendError(ctx, http.StatusBadRequest, "Password is too short or too long.")
		return
	}

	if !utils.UsernameRegex.MatchString(username) {
		helpers.SendError(ctx, http.StatusBadRequest, "Invalid Username")
		return
	}

	if !utils.EmailRegex.MatchString(email) {
		helpers.SendError(ctx, http.StatusBadRequest, "Invalid Email.")
		return
	}

	if err = configs.DB.Where("email = ?", email).First(&user).Error; err != gorm.ErrRecordNotFound {
		helpers.SendError(ctx, http.StatusConflict, "Email already exists.")
		return
	}

	if err = configs.DB.Where("username = ?", username).First(&user).Error; err != gorm.ErrRecordNotFound {
		helpers.SendError(ctx, http.StatusBadRequest, "Username already exists.")
		return
	}

	pswd := gobcrypt.HashPassword(password)

	user = models.Users{
		Email:    email,
		Password: pswd,
		Username: username,
	}

	if err = configs.DB.Create(&user).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SendSuccess(ctx, http.StatusOK, "Account has been created successfully.", nil)
}

func Login(ctx *gin.Context) {
	var (
		err  error
		user models.Users
		body structs.Login
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Invalid request body.")
		return
	}

	var userId string = strings.ToLower(strings.TrimSpace(body.UserId))

	if len(body.Password) == 0 || len(userId) == 0 {
		helpers.SendError(ctx, http.StatusBadRequest, "All fields are required.")
		return
	}

	if utils.EmailRegex.MatchString(userId) {
		if err = configs.DB.Where("email = ?", userId).First(&user).Error; err != nil {
			helpers.SendError(ctx, http.StatusNotFound, "Invalid email or password.")
			return
		}
	} else {
		if err = configs.DB.Where("username = ?", userId).First(&user).Error; err != nil {
			helpers.SendError(ctx, http.StatusNotFound, "Invalid username or password.")
			return
		}
	}

	isMatch := gobcrypt.CompareHashAndPassword(user.Password, body.Password)
	if !isMatch {
		helpers.SendError(ctx, http.StatusUnauthorized, "Incorrect password.")
		return
	}

	utils.GenTokens(ctx, user.Username, user.ID.String())

	helpers.SendSuccess(ctx, http.StatusOK, "Login successful.", nil)
}

func clearCookies(ctx *gin.Context, cookieNames []string) {
	for _, cookieName := range cookieNames {
		ctx.SetCookie(cookieName, "", -1, "/", "localhost", false, true)
	}
	ctx.Status(http.StatusNoContent)
}

func Logout(ctx *gin.Context) {
	refresh_token, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}

	cookieNames := []string{"refresh_token", "access_token"}

	if err := configs.DB.Model(&models.Users{}).Where("refresh_token = ?", refresh_token).Update("refresh_token", "").Error; err != nil {
		clearCookies(ctx, cookieNames)
		return
	}

	clearCookies(ctx, cookieNames)
}
