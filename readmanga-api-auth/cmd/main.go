package main

import (
	"readmanga-api-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	api := server.Group("/api/v1")
	{
		routes.RegisterUserRoutes(api)
	}

	server.Run(":3001")
}
