package routes

import (
	"readmanga-api-auth/adapter/controllers"
	"readmanga-api-auth/adapter/presenters"
	"readmanga-api-auth/application"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, app application.UserApplication) {
	router.POST("/user/auth/create-user", func(c *gin.Context) {
		ctx := &presenters.Context{C: c}
		controller := controllers.NewUserController(app)
		controller.CreateUser(ctx)
	})

	router.POST("/user/auth/login", func(c *gin.Context) {
		ctx := &presenters.Context{C: c}
		controller := controllers.NewUserController(app)
		controller.LoginUser(ctx)
	})
}
