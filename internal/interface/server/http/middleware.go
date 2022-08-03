package http

import (
	"net/http"
	"time"

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

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	server.Use(middleware.RateLimiterWithConfig(config))

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
