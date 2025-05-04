package controllers

import (
	"fmt"
	"net/http"
	"readmanga-api-auth/adapter/presenters"
	"readmanga-api-auth/application"
	"readmanga-api-auth/domain"
	"readmanga-api-auth/internal/auth"

	"gorm.io/gorm"
)

type userController struct {
	userUseCase application.UserApplication
}

func NewUserController(usecase application.UserApplication) userController {
	return userController{userUseCase: usecase}
}

func (u *userController) CreateUser(ctx *presenters.Context) {
	var user *domain.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": "Dados inválidos",
		})
		return
	}

	_, err := u.userUseCase.GetUserParams(map[string]string{
		"email": user.Email,
	})

	if err == nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": fmt.Sprintf("E-mail '%s' já é existente.", user.Email),
		})
		return
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  0,
			"message": "Erro ao verificar e-mail.",
		})
		return
	}

	result, err := u.userUseCase.Register(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": "Não foi possivel criar usuário.",
		})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  1,
		"message": fmt.Sprintf("Welcome to %s", user.Nickname),
		"data":    result.ID,
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

func (u *userController) LoginUser(ctx *presenters.Context) {
	var user *domain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": "Dados inválidos",
		})
		return
	}
	query, err := u.userUseCase.GetUserParams(map[string]string{
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  0,
			"message": fmt.Sprintf("Login de %s é inválido!", user.Email),
		})
		return
	}
	token, err := auth.GenerateJWT(query.Email, query.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": fmt.Sprintf("Não foi possível fazer login com '%s'", query.Email),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  1,
		"message": fmt.Sprintf(""),
		"data":    token,
	})
}
