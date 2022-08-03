package http

import "github.com/agungwicaksana/privy-pretest/internal/interface/container"

type handler struct {
	cakeHandler *cakeHandler
}

func SetupHandlers(container *container.Container) *handler {
	return &handler{
		cakeHandler: NewCakeHandler().SetCakeService(container.CakeService).Validate(),
	}
}

func (h *handler) Validate() *handler {
	return h
}
