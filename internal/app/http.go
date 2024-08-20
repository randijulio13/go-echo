package app

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

var appServer *echo.Echo

func InitializeServer() *echo.Echo {
	appServer = echo.New()
	appServer.HideBanner = true
	return appServer
}

func StartServer() *echo.Echo {
	port := os.Getenv("APP_PORT")
	appServer.Start(fmt.Sprintf(":%s", port))
	return appServer
}
