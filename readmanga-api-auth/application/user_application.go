package application

import (
	"readmanga-api-auth/domain"
	"readmanga-api-auth/infra/repository"
)

type UserApplication interface {
	Register(user *domain.User) (*domain.User, error)
	GetUserId(id int) (*domain.User, error)
}

type userApplication struct {
	repo repository.UserRepository
}

func NewUserApplication(repo repository.UserRepository) UserApplication {
	return &userApplication{repo}
}

func (app *userApplication) Register(user *domain.User) (*domain.User, error) {
	queryParams, err := app.repo.FindUser(map[string]interface{}{"email": user.Email})
	if err != nil || queryParams != nil {
		return nil, err
	}

	if err := app.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *userApplication) GetUserId(id int) (*domain.User, error) {
	return app.repo.GetById(id)
}
