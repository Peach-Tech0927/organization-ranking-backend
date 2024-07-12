package main

import (
	"organization-ranking-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.Run(":8080")
}