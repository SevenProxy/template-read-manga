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
		"data":    map[string]interface{}{"id": result.ID},
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
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": "Dados inválidos",
		})
		return
	}
	query, err := u.userUseCase.GetUserParams(map[string]string{
		"email":    input.Email,
		"password": input.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  0,
			"message": fmt.Sprintf("Credenciais inválidas"),
		})
		return
	}
	token, err := auth.GenerateJWT(query.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": fmt.Sprintf("Erro ao gerar token"),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  1,
		"message": fmt.Sprintf("Login realizado com sucesso"),
		"data":    token,
	})
}
