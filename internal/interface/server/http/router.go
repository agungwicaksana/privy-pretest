package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRouter(handler *handler) *echo.Echo {
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Service is healthy")
	})

	return server
}
