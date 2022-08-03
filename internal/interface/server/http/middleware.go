package http

import (
	"net/http"

	"github.com/agungwicaksana/privy-pretest/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(server *echo.Echo) {
	// server.Use(middleware.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders:     []string{"Accept", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))
	server.HTTPErrorHandler = errorHandler
}

func errorHandler(err error, c echo.Context) {
	if c.Get("error-handled") != nil {
		return
	}

	c.Set("error-handled", true)

	if c.Get("RC") == nil {
		c.Set("RC", response.STATUS_GENERAL_ERROR)
	}

	code := c.Get("RC").(string)
	resp := response.CreateResponse(code, err.Error(), nil)
	err = c.JSON(http.StatusOK, resp)
}
