package http

import (
	"github.com/agungwicaksana/privy-pretest/internal/interface/container"
	"github.com/labstack/echo/v4"
)

func StartHttpService(container *container.Container) {
	echo := echo.New()
	SetupMiddleware(echo)
	handler := SetupHandlers(container).Validate()
	router := SetupRouter(echo, handler)
	port := ":3000"
	router.Logger.Fatal(router.Start(port))
}
