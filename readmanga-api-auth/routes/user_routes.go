package routes

import (
	"readmanga-api-auth/adapter/controllers"
	"readmanga-api-auth/adapter/presenters"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/user/auth/create-user", func(c *gin.Context) {
		ctx := &presenters.Context{C: c}
		controller := controllers.NewUserController()
		controller.GetUser(ctx)
	})
}
