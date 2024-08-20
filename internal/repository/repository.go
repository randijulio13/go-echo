package repository

import "go-echo/internal/entity/model"

type UserRepository interface {
	CreateUser(*model.User) error
	AuthenticateUser(*model.User) error
}
