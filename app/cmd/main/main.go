package main

import (
	"go-architecture/app/cmd/server"
	"go-architecture/app/config"
	"go-architecture/app/domains/v1/modules"
)

func main() {
	cfg := config.LoadConfig()
	s := server.InitServer(cfg)

	modules.Modules(s)

	server.Run(s)
}
