package app

import (
	"go-echo/internal/router"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoadRouter(e *echo.Echo, db *gorm.DB) {
	router.Routes(e, db)
}
