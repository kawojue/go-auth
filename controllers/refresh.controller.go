package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-initenv"
)

// on it

var secretKey []byte = []byte(initenv.GetEnv("JWT_SECRET", ""))

func RefreshToken(ctx *gin.Context) {
	refresh_token, err := ctx.Cookie("refresh_token")
	if err != nil {
		helpers.SendError(ctx, http.StatusUnauthorized, "Access denied.")
		return
	}

	token, err := jwt.ParseWithClaims(refresh_token, &structs.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		helpers.SendError(ctx, http.StatusForbidden, "Access denied.")
		return
	}

	claims, ok := token.Claims.(*structs.Claims)
	if !ok || !token.Valid {
		helpers.SendError(ctx, http.StatusForbidden, "Access denied.")
		return
	}

	access_token_exp := time.Now().Add(1 * time.Hour)
	access_token_claims := &structs.Claims{
		Username: claims.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: access_token_exp.Unix(),
		},
	}
	new_access_token_claims := jwt.NewWithClaims(jwt.SigningMethodHS256, access_token_claims)
	access_token, err := new_access_token_claims.SignedString(secretKey)

	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error generating new access token.")
		return
	}

	ctx.SetCookie("access_token", access_token, int(time.Until(access_token_exp)), "/", "localhost", false, true)

	helpers.SendSuccess(ctx, http.StatusOK, "Access token generated.", nil)
}
