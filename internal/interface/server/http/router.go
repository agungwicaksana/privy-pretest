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
	cakes.GET("/:id", handler.cakeHandler.FindOne)
	cakes.PATCH("/:id", handler.cakeHandler.Update)
	cakes.DELETE("/:id", handler.cakeHandler.Delete)

	return server
}
