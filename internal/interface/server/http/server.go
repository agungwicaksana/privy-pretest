package http

import (
	"github.com/agungwicaksana/privy-pretest/internal/interface/container"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func StartHttpService(container *container.Container) {
	echo := echo.New()
	echo.Validator = &CustomValidator{validator: validator.New()}
	SetupMiddleware(echo)
	handler := SetupHandlers(container).Validate()
	router := SetupRouter(echo, handler)
	port := ":3000"
	router.Logger.Fatal(router.Start(port))
}
