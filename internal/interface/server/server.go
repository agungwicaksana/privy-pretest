package server

import (
	"github.com/agungwicaksana/privy-pretest/internal/interface/container"
	"github.com/agungwicaksana/privy-pretest/internal/interface/server/http"
)

func StartService(container *container.Container) {
	http.StartHttpService(container)
}
