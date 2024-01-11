package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/helpers"
)

func LimitRequestBodySize(maxSize int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.Request.ParseMultipartForm(maxSize); err != nil {
			helpers.SendError(ctx, http.StatusRequestEntityTooLarge, "File too large")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
