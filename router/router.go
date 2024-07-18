package router

import (
	"organization-ranking-backend/controllers"
	"organization-ranking-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")
	auth := public.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
	public.GET("/organizations-result",controllers.GetOrganizationsRanking)

	protected := router.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	return router
}