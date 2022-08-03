package helpers

import (
	"github.com/agungwicaksana/privy-pretest/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Validate(c echo.Context, s interface{}) (err error) {
	ctx := c.Request().Context()

	if err = c.Bind(s); err != nil {
		log.Error(ctx, "Bind error", err.Error())
		c.Set("RC", response.STATUS_REQUEST_BINDING_FAILED)
		return
	}

	log.Info(ctx, "Incoming request", s)

	if err = c.Validate(s); err != nil {
		log.Error(ctx, "Validation error", err.Error())
		c.Set("RC", response.STATUS_REQUEST_VALIDATION_FAILED)
		return
	}

	return
}
