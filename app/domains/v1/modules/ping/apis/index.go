package apis

import (
	"go-architecture/app/domains/v1/modules/ping/model"
)

type APIs struct {
	pingHandler model.Logic
}

func NewAPIs(pingHandler model.Logic) *APIs {
	return &APIs{
		pingHandler: pingHandler,
	}
}
