package main

import (
	"organization-ranking-backend/models"
	"organization-ranking-backend/router"

	"github.com/gin-contrib/cors"
)

func main() {
	models.ConnectDatabase()

	r := router.SetUpRouter()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))

	r.Run(":8080")
}