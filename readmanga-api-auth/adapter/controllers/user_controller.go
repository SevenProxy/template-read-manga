package controllers

import (
	"net/http"
	"readmanga-api-auth/adapter/presenters"
)

type userController struct {
}

func NewUserController() userController {
	return userController{}
}

func (user *userController) GetUser(ctx *presenters.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "hello world",
	})
}
