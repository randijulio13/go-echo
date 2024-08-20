package handler

import (
	"go-echo/internal/entity/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func parseValidationError(errorString string) response.ValidationErrors {
	var errors []response.FieldError

	errorString = strings.TrimPrefix(errorString, "code=400, message=")
	parts := strings.Split(errorString, "\n")

	for _, part := range parts {
		if part != "" {
			keyValue := strings.SplitN(part, " Error:", 2)
			if len(keyValue) == 2 {
				fieldError := response.FieldError{
					Key:   strings.TrimSpace(keyValue[0]),
					Error: strings.TrimSpace(keyValue[1]),
				}
				errors = append(errors, fieldError)
			}
		}
	}

	return response.ValidationErrors{Errors: errors}
}

type AuthHandler interface {
	AuthenticateUser(echo.Context) error
	CreateUser(echo.Context) error
}
