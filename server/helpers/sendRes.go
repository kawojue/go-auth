package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"success": true,
		"data":    data,
		"message": message,
	})
}

func SendError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   message,
	})
}

var SOMETHING_WENT_WRONG = func(ctx *gin.Context) {
	SendError(ctx, http.StatusInternalServerError, "Something went wrong")
}

var ACCOUNT_NOT_FOUND = func(ctx *gin.Context) {
	SendError(ctx, http.StatusNotFound, "Account not found")
}

var INVALID_JSON = func(ctx *gin.Context) {
	SendError(ctx, http.StatusInternalServerError, "Error binding Data")
}

var FAILED_TO_SAVE_OTP = func(ctx *gin.Context) {
	SendError(ctx, http.StatusInternalServerError, "Failed to save OTP in the database")
}

var ALL_FIELDS_REQUIRED = func(ctx *gin.Context) {
	SendError(ctx, http.StatusBadRequest, "All fields are required")
}

var FORBIDDEN_ACESS_DENIED = func(ctx *gin.Context) {
	SendError(ctx, http.StatusForbidden, "Access denied")
}
var UNAUTHORIZED_ACESS_DENIED = func(ctx *gin.Context) {
	SendError(ctx, http.StatusUnauthorized, "Access denied")
}
