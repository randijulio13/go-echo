package service

import (
	"go-echo/internal/entity/model"
	"go-echo/internal/entity/request"
	"go-echo/internal/entity/response"
)

type serviceResult struct {
	HttpStatusCode int
	Data           response.HttpResponse
}

type AuthService interface {
	CreateUser(*model.User) serviceResult
	AuthenticateUser(*request.AuthRequest) serviceResult
}

type GormErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}
