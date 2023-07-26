package logic

import (
	"go-architecture/app/cmd/server"
	"go-architecture/app/domains/v1/modules/ping/model"
)

type PingHandler struct {
	PingRepo model.Repo
	Server   *server.Server
}

func NewPingHandler(server *server.Server, pingRepo model.Repo) model.Logic {
	return &PingHandler{
		PingRepo: pingRepo,
		Server:   server,
	}
}
