package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/models"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-initenv"
)

func GenTokens(ctx *gin.Context, username string, id string) {
	access_token_exp := time.Now().Add(2 * time.Hour).UTC()
	refresh_token_exp := time.Now().Add(120 * 24 * time.Hour).UTC()

	access_token_claims := &structs.Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: access_token_exp.Unix(),
		},
	}
	refresh_token_claims := &structs.Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refresh_token_exp.Unix(),
		},
	}

	new_access_token_claims := jwt.NewWithClaims(jwt.SigningMethodHS256, access_token_claims)
	access_token, err := new_access_token_claims.SignedString([]byte(initenv.GetEnv("JWT_SECRET", "")))
	if err != nil {
		helpers.SOMETHING_WENT_WRONG(ctx)
		return
	}

	new_refresh_token_claims := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_token_claims)
	refresh_token, err := new_refresh_token_claims.SignedString([]byte(initenv.GetEnv("JWT_SECRET", "")))
	if err != nil {
		helpers.SOMETHING_WENT_WRONG(ctx)
		return
	}

	if err := configs.DB.Model(&models.Users{}).Where("username = ?", username).Update("refresh_token", refresh_token).Error; err != nil {
		helpers.SOMETHING_WENT_WRONG(ctx)
		return
	}

	ctx.SetCookie("access_token", access_token, int(time.Until(access_token_exp).Seconds()), "/", "localhost", false, false)
	ctx.SetCookie("refresh_token", refresh_token, int(time.Until(refresh_token_exp).Seconds()), "/", "localhost", false, true)
}
