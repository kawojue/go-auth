package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-initenv"
)

func RefreshToken(ctx *gin.Context) {
	secretKey := []byte(initenv.GetEnv("JWT_SECRET", ""))

	refresh_token, err := ctx.Cookie("refresh_token")
	if err != nil {
		helpers.SendError(ctx, http.StatusUnauthorized, "Access denied.")
		return
	}

	fmt.Println(secretKey)

	token, err := jwt.ParseWithClaims(refresh_token, &structs.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		helpers.SendError(ctx, http.StatusForbidden, "Access denied.")
		return
	}

	claims, ok := token.Claims.(*structs.Claims)
	if !ok || !token.Valid || time.Now().Unix() > claims.ExpiresAt {
		helpers.SendError(ctx, http.StatusForbidden, "Access denied.")
		return
	}

	access_token_exp := time.Now().Add(2 * time.Hour).UTC()
	access_token_claims := &structs.Claims{
		ID:       claims.ID,
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

	ctx.SetCookie("access_token", access_token, int(time.Until(access_token_exp).Seconds()), "/", "localhost", false, true)

	helpers.SendSuccess(ctx, http.StatusOK, "New access token generated.", nil)
}
