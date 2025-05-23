package main

import (
	"readmanga-api-auth/application"
	"readmanga-api-auth/config"
	"readmanga-api-auth/domain"
	"readmanga-api-auth/infra/database"
	"readmanga-api-auth/infra/repository"
	"readmanga-api-auth/infra/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := database.Connect()
	db.AutoMigrate(&domain.User{}, &domain.Notification{})

	userRepo := repository.NewUserRepository(db)
	userApp := application.NewUserApplication(userRepo)

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.GetEnv("FRONT_END")},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := server.Group("/api/v1")
	{
		routes.RegisterUserRoutes(api, userApp)
	}

	server.Run(":3001")
}
