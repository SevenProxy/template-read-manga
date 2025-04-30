package controllers

import (
	"net/http"
	"readmanga-api-auth/adapter/presenters"
	"readmanga-api-auth/application"
)

type userController struct {
	userUseCase application.UserApplication
}

func NewUserController(usecase application.UserApplication) userController {
	return userController{userUseCase: usecase}
}

func (u *userController) CreateUser(ctx *presenters.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "welcome",
		"data":    "aa",
	})
}

func (u *userController) GetUser(ctx *presenters.Context) {
	user, err := u.userUseCase.GetUserId(1)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "Usuário não foi encontrado.",
		})
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "welcome",
		"data":    user,
	})
}
