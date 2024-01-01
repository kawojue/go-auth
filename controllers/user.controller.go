package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/db"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
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
		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
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

func Login(ctx *gin.Context) {

}
