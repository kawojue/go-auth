package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/db"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	"github.com/kawojue/go-auth/utils"
	gobcrypt "github.com/kawojue/go-bcrypt"
	"gorm.io/gorm"
)

type SignUpBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func SignUp(ctx *gin.Context) {
	var (
		err        error
		signUpBody SignUpBody
		user       models.Users
	)

	if err = ctx.ShouldBindJSON(&signUpBody); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Invalid request body.")
		return
	}

	password := signUpBody.Password
	email := strings.TrimSpace(signUpBody.Email)
	username := strings.ToLower(strings.TrimSpace(signUpBody.Email))

	if len(password) == 0 || len(email) == 0 || len(username) == 0 {
		helpers.SendError(ctx, http.StatusBadRequest, "All fields are required.")
		return
	}

	if len([]byte(password)) > 72 {
		helpers.SendError(ctx, http.StatusBadRequest, "Password should be at most 72 characters.")
		return
	}

	if len([]byte(password)) < 8 {
		helpers.SendError(ctx, http.StatusBadRequest, "Password should be at least 8 characters.")
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

	if err = db.DB.Where("email = ?", email).First(&user).Error; err != gorm.ErrRecordNotFound {
		helpers.SendError(ctx, http.StatusConflict, "Email already exists.")
		return
	}

	if err = db.DB.Where("username = ?", username).First(&user).Error; err != gorm.ErrRecordNotFound {
		helpers.SendError(ctx, http.StatusBadRequest, "Username already exists.")
		return
	}

	pswd := gobcrypt.HashPassword(password)

	user = models.Users{
		Email:    email,
		Password: pswd,
		Username: username,
	}

	if err = db.DB.Create(&user).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SendSuccess(ctx, http.StatusOK, map[string]string{
		"msg": "Account has been created successfully.",
	})
}

type LoginBody struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var (
		err            error
		recNotFoundErr error
		user           models.Users
		body           LoginBody
		userId         string = strings.ToLower(strings.TrimSpace(body.UserId))
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Invalid request body.")
		return
	}

	if len(body.Password) == 0 || len(userId) == 0 {
		helpers.SendError(ctx, http.StatusBadRequest, "All fields are required.")
		return
	}

	if utils.EmailRegex.MatchString(userId) {
		recNotFoundErr = db.DB.Where("email = ?", userId).First(&user).Error
	} else {
		recNotFoundErr = db.DB.Where("username = ?", userId).First(&user).Error
	}

	if recNotFoundErr == gorm.ErrRecordNotFound {
		helpers.SendError(ctx, http.StatusNotFound, "Invalid username or password.")
		return
	}

	isMatch := gobcrypt.VerifyPassword(user.Password, body.Password)
	if !isMatch {
		helpers.SendError(ctx, http.StatusUnauthorized, "Incorrect password.")
		return
	}

}
