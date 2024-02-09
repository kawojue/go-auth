package middlewares

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/structs"
	initenv "github.com/kawojue/go-initenv"
)

func VerifyAuth() gin.HandlerFunc {
	secretKey := []byte(initenv.GetEnv("JWT_SECRET", ""))

	return func(ctx *gin.Context) {
		access_token, err := ctx.Cookie("access_token")
		if err != nil {
			helpers.UNAUTHORIZED_ACESS_DENIED(ctx)
			return
		}

		token, err := jwt.ParseWithClaims(
			access_token, &structs.Claims{}, func(t *jwt.Token,
			) (interface{}, error) {
				return secretKey, nil
			})

		if err != nil {
			helpers.UNAUTHORIZED_ACESS_DENIED(ctx)
			return
		}

		claims, ok := token.Claims.(*structs.Claims)
		if !ok || !token.Valid || time.Now().Unix() > claims.ExpiresAt {
			helpers.FORBIDDEN_ACESS_DENIED(ctx)
			return
		}

		ctx.Set("user_id", claims.ID)
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}
