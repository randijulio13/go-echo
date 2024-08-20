package router

import (
	"go-echo/internal/handler"

	"github.com/labstack/echo/v4"
)

func authRouter(e *echo.Echo, h handler.AuthHandler) {
	g := e.Group("/user")
	g.GET("", h.AuthenticateUser)
	g.POST("", h.CreateUser)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
}
