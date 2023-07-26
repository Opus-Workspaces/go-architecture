package modules

import (
	"go-architecture/app/cmd/server"
	"go-architecture/app/domains/v1/modules/ping"
	"go-architecture/app/routers"
)

func Modules(server *server.Server) {
	apiVersion := server.Echo.Group(routers.Ver1)
	// Using middleware if needed

	ping.RouterPing(server, apiVersion.Group(routers.Ping))
}
