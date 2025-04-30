package repository

import (
	"readmanga-api-auth/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetById(id int) (*domain.User, error)
	FindUser(params interface{}) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) FindUser(params interface{}) (*domain.User, error) {
	var user domain.User
	if err := repo.db.First(&user, params).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) Create(user *domain.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepository) GetById(id int) (*domain.User, error) {
	var user domain.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
