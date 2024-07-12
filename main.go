package main

import (
	"organization-ranking-backend/models"
	"organization-ranking-backend/router"
)

func main() {
	models.ConnectDatabase()

	r := router.SetUpRouter()
	r.Run(":8080")
}