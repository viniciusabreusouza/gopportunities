package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, status int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"message":   message,
		"errorCode": status,
	})
}

func SendSuccess(ctx *gin.Context, op string, status int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"message": fmt.Sprintf("%s successfully", op),
		"data":    data,
	})
}
