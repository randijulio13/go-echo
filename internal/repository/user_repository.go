package repository

import (
	"go-echo/internal/entity/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) CreateUser(user *model.User) error {
	user.HashPassword()
	err := repo.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) AuthenticateUser(user *model.User) error {

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
