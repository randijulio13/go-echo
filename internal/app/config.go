package app

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func LoadEnv(e *echo.Echo) {
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}
}
