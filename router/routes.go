package router

import (
	"silvanotfound/tasks/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/task", handler.ShowTaskHandler)
		v1.GET("/tasks", handler.ShowAllTasksHandler)
		v1.POST("/create", handler.CreateTaskHandler)
		v1.PUT("/update", handler.UpdateTaskHandler)
		v1.DELETE("/delete", handler.DeleteTaskHandler)
	}
}
