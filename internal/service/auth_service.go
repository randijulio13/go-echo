package service

import (
	"encoding/json"
	"go-echo/internal/entity/model"
	"go-echo/internal/entity/request"
	"go-echo/internal/entity/response"
	"go-echo/internal/repository"
	"net/http"
)

type authService struct {
	repo repository.UserRepository
}

func (svc *authService) CreateUser(user *model.User) serviceResult {
	if err := svc.repo.CreateUser(user); err != nil {
		byteErr, _ := json.Marshal(err)
		var newError GormErr
		json.Unmarshal((byteErr), &newError)
		return serviceResult{
			Data: response.HttpResponse{
				Status:  "error",
				Message: newError.Message,
			},
			HttpStatusCode: http.StatusConflict,
		}
	}
	return serviceResult{
		Data: response.HttpResponse{
			Status: "success",
		},
		HttpStatusCode: http.StatusOK,
	}
}

func (svc *authService) AuthenticateUser(req *request.AuthRequest) serviceResult {

	return serviceResult{
		Data: response.HttpResponse{
			Status: "success",
		},
		HttpStatusCode: http.StatusOK,
	}
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}
