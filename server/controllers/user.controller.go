package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-auth/utils"
	gobcrypt "github.com/kawojue/go-bcrypt"
	"github.com/kawojue/go-initenv"
	"gorm.io/gorm"
)

func SignUp(ctx *gin.Context) {
	var (
		err  error
		body structs.SignUp
		user models.Users
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		helpers.INVALID_JSON(ctx)
		return
	}

	password := body.Password
	password2 := body.Password2
	email := strings.TrimSpace(body.Email)
	username := strings.ToLower(strings.TrimSpace(body.Username))

	if len(password) == 0 || len(email) == 0 || len(username) == 0 {
		helpers.ALL_FIELDS_REQUIRED(ctx)
		return
	}

	if password != password2 {
		helpers.SendError(ctx, http.StatusBadRequest, "Passwords not match")
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
		helpers.INVALID_JSON(ctx)
		return
	}

	var userId string = strings.ToLower(strings.TrimSpace(body.UserId))

	if len(body.Password) == 0 || len(userId) == 0 {
		helpers.ALL_FIELDS_REQUIRED(ctx)
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

	helpers.SendSuccess(ctx, http.StatusOK, "Login successful.", map[string]string{
		"username": user.Username,
	})
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

func GetUser(ctx *gin.Context) {
	var (
		err             error
		username        string
		user            models.Users
		isAuthenticated bool = false
		data            map[string]interface{}
	)
	secretKey := []byte(initenv.GetEnv("JWT_SECRET", ""))
	access_token, _ := ctx.Cookie("access_token")
	usernameParam := ctx.Param("username")

	isValid, usernameAuth := validateToken(access_token, secretKey)

	if isValid && usernameAuth == usernameParam {
		username = usernameAuth
		isAuthenticated = true
	} else {
		username = usernameParam
	}

	if err = configs.DB.Where("username = ?", username).First(&user).Preload("Profiles").Error; err != nil {
		helpers.ACCOUNT_NOT_FOUND(ctx)
		return
	}

	if isAuthenticated {
		data = map[string]interface{}{
			"user": user,
		}
	} else {
		data = map[string]interface{}{
			"username": user.Username,
		}
	}

	helpers.SendSuccess(ctx, http.StatusOK, "", data)
}

func validateToken(tokenString string, secretKey []byte) (bool, string) {
	token, err := jwt.ParseWithClaims(tokenString, &structs.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, ""
	}

	claims, ok := token.Claims.(*structs.Claims)
	if !ok || !token.Valid || time.Now().Unix() > claims.ExpiresAt {
		return false, ""
	}

	return true, claims.Username
}
