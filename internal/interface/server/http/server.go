package http

import (
	"github.com/agungwicaksana/privy-pretest/internal/interface/container"
)

func StartHttpService(container *container.Container) {
	router := SetupRouter(SetupHandlers(container).Validate())
	port := ":3000"
	router.Logger.Fatal(router.Start(port))
}
