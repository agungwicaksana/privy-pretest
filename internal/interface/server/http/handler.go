package http

import "github.com/agungwicaksana/privy-pretest/internal/interface/container"

type handler struct {
}

func SetupHandlers(container *container.Container) *handler {
	return &handler{}
}

func (h *handler) Validate() *handler {
	return h
}
