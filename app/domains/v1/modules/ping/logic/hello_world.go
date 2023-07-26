package logic

import (
	"context"
	"go-architecture/app/models"
	"net/http"
)

func (p PingHandler) Ping(ctx context.Context) (*models.ResponseSucceed, *models.ResponseError) {
	message, err := p.PingRepo.Ping(ctx)
	if err != nil {
		return nil, &models.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	return &models.ResponseSucceed{
		StatusCode: http.StatusOK,
		Data:       message,
	}, nil
}
