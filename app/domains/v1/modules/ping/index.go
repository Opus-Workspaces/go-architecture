package ping

import (
	"github.com/labstack/echo/v4"
	"go-architecture/app/cmd/server"
	_api "go-architecture/app/domains/v1/modules/ping/apis"
	_logic "go-architecture/app/domains/v1/modules/ping/logic"
	_repo "go-architecture/app/domains/v1/modules/ping/repo"
)

const (
	helloWorld = ""
)

func RouterPing(s *server.Server, g *echo.Group) {

	r := _repo.NewRepo(s.DBMongo.DB)
	l := _logic.NewPingHandler(s, r)
	a := _api.NewAPIs(l)

	g.GET(helloWorld, a.Ping)

}
