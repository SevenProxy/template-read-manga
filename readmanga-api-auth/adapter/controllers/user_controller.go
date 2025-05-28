package controllers

import (
	"fmt"
	"net/http"
	"readmanga-api-auth/adapter/middleware"
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
		"message": fmt.Sprintf("Welcome to %s", result.Nickname),
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
			"message": "Unauthorized: Dados inválidos",
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
			"message": "Unauthorized: Credenciais inválidas",
		})
		return
	}
	token, err := auth.GenerateJWT(query.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  0,
			"message": "Unauthorized: Erro ao gerar token",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  1,
		"message": "Login realizado com sucesso",
		"data":    token,
	})
}

func (u *userController) TokenLogin(ctx *presenters.Context) {
	email, ok := ctx.Value(middleware.UserEmailKey)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  0,
			"message": "Unauthorized: Usuário nao autorizado.",
		})
		return
	}

	query, err := u.userUseCase.GetUserParams(map[string]string{
		"email": email,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  0,
			"message": "Unauthorized: Usuário não autorizado.",
		})
		return
	}

	user := map[string]interface{}{
		"id":        query.ID,
		"nickname":  query.Nickname,
		"avatar":    query.Avatar,
		"createdAt": query.CreatedAt,
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": 1,
		"data":   user,
	})
}
