package router

import (
	"organization-ranking-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	protected := router.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	return router
}