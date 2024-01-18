package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create task",
	})
}

func UpdateTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update task",
	})
}

func DeleteTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete success",
	})
}

func ShowTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Show success",
	})
}

func ShowAllTasksHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Show all tasks success",
	})
}
