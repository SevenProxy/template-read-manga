package routes

import (
	"readmanga-api-auth/adapter/controllers"
	"readmanga-api-auth/adapter/middleware"
	"readmanga-api-auth/adapter/presenters"
	"readmanga-api-auth/application"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, app application.UserApplication) {
	router.POST("/auth/create-user", func(c *gin.Context) {
		ctx := &presenters.Context{C: c}
		controller := controllers.NewUserController(app)
		controller.CreateUser(ctx)
	})

	router.POST("/auth/login", func(c *gin.Context) {
		ctx := &presenters.Context{C: c}
		controller := controllers.NewUserController(app)
		controller.LoginUser(ctx)
	})

	authRoutes := router.Group("/user")
	authRoutes.Use(middleware.AuthMiddlewareGin())
	{
		authRoutes.GET("/me", func(c *gin.Context) {
			ctx := &presenters.Context{C: c}
			controller := controllers.NewUserController(app)
			controller.TokenLogin(ctx)
		})
	}
}
