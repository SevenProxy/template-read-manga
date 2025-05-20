package application

import (
	"errors"
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
	email, ok := params["email"]
	if !ok || email == "" {
		return nil, errors.New("email is required")
	}

	query, err := app.repo.FindUser(map[string]string{"email": params["email"]})
	if err != nil {
		return nil, err
	}

	if password, ok := params["password"]; ok && password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(query.Password), []byte(password))
		if err != nil {
			return nil, errors.New("invalid password")
		}
	} else {
		return nil, errors.New("Invalid password")
	}

	return query, nil
}
