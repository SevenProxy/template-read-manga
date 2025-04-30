package main

import (
	"readmanga-api-auth/infra/database"
	"readmanga-api-auth/infra/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	server := gin.Default()

	api := server.Group("/api/v1")
	{
		routes.RegisterUserRoutes(api)
	}

	server.Run(":3001")
}
