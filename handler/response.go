package handler

import "github.com/gin-gonic/gin"

func sendResponseError(ctx *gin.Context, message string) {
	ctx.JSON(400, gin.H{
		"message": message,
	})
}

func sendResponseSuccess(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, gin.H{
		"message": message,
		"data":    data,
	})
}
