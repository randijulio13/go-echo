package handler

import (
	"go-echo/internal/entity/model"
	"go-echo/internal/entity/request"
	"go-echo/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	svc service.AuthService
}

func (h *authHandler) AuthenticateUser(c echo.Context) error {
	req := new(request.AuthRequest)
	if err := c.Bind(req); err != nil {
		validationError := parseValidationError(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, validationError)
	}

	res := h.svc.AuthenticateUser(req)
	return c.JSON(res.HttpStatusCode, res.Data)
}

func (h *authHandler) CreateUser(c echo.Context) error {

	req := new(request.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return err
		// validationError := parseValidationError(err.Error())
		// return c.JSON(http.StatusBadRequest, validationError)
	}

	user := model.User{
		ID:       uuid.NewString(),
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	res := h.svc.CreateUser(&user)

	return c.JSON(res.HttpStatusCode, res.Data)
}

func NewAuthHandler(svc service.AuthService) AuthHandler {
	return &authHandler{svc: svc}
}
