package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRouter(server *echo.Echo, handler *handler) *echo.Echo {
	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Service is healthy")
	})

	cakes := server.Group("/cakes")
	cakes.GET("", handler.cakeHandler.Find)
	cakes.POST("", handler.cakeHandler.Save)

	return server
}
