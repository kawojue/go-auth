package helpers

import "github.com/gin-gonic/gin"

func SendSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"success": true,
		"data":    data,
	})
}

func SendError(ctx *gin.Context, statusCode int, msg string) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   msg,
	})
}
