package cmd

import (
	"github.com/agungwicaksana/privy-pretest/internal/interface/container"
	"github.com/agungwicaksana/privy-pretest/internal/interface/server"
)

func Run() {
	container := container.Setup()
	server.StartService(container)
}
