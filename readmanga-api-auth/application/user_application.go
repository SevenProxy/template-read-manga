package application

import (
	"fmt"
	"readmanga-api-auth/domain"
	"readmanga-api-auth/infra/repository"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type UserApplication interface {
	Register(user *domain.User) (*domain.User, error)
	GetUserId(id int) (*domain.User, error)
	GetUserParams(params map[string]string) (*domain.User, error)
}

type userApplication struct {
	repo repository.UserRepository
}

func NewUserApplication(repo repository.UserRepository) UserApplication {
	return &userApplication{repo}
}

func (app *userApplication) Register(user *domain.User) (*domain.User, error) {
	queryParams, err := app.repo.FindUser(map[string]interface{}{"email": user.Email})
	if err == nil && queryParams != nil {
		return nil, fmt.Errorf("usuário já existe")
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	if err := app.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *userApplication) GetUserId(id int) (*domain.User, error) {
	return app.repo.GetById(id)
}

func (app *userApplication) GetUserParams(params map[string]string) (*domain.User, error) {
	if password, ok := params["password"]; ok && password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		params["password"] = string(hashedPassword)
	}

	query, err := app.repo.FindUser(params)
	if err != nil {
		return nil, err
	}

	return query, nil
}
