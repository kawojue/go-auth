package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/structs"
	initenv "github.com/kawojue/go-initenv"
)

var secretKey []byte = []byte(initenv.GetEnv("JWT_SECRET", ""))

func VerifyAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		access_token, err := ctx.Cookie("access_token")
		if err != nil {
			helpers.SendError(ctx, http.StatusUnauthorized, "Access denied.")
			return
		}

		token, err := jwt.ParseWithClaims(
			access_token, &structs.Claims{}, func(t *jwt.Token,
			) (interface{}, error) {
				return secretKey, nil
			})

		if err != nil {
			helpers.SendError(ctx, http.StatusUnauthorized, "Access denied.")
			return
		}

		claims, ok := token.Claims.(*structs.Claims)
		if !ok || !token.Valid {
			helpers.SendError(ctx, http.StatusForbidden, "Access denied.")
			return
		}

		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}
