package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/task", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Teste success",
			})
		})

		v1.GET("/tasks", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Teste success",
			})
		})

		v1.POST("/create", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Teste success",
			})
		})

		v1.PUT("/update", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Teste success",
			})
		})

		v1.DELETE("/delete", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Teste success",
			})
		})
	}
}
