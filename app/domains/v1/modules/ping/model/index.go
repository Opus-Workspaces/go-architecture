package model

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-architecture/app/models"
)

type Ping struct {
	Message string `json:"message"`
}

type APIs interface {
	Ping(ctx echo.Context) error
}

type Logic interface {
	Ping(ctx context.Context) (*models.ResponseSucceed, *models.ResponseError)
}

type Repo interface {
	Ping(ctx context.Context) (string, error)
}
