package apis

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *APIs) Ping(e echo.Context) error {
	res, err := a.pingHandler.Ping(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	return e.JSON(http.StatusOK, res)
}
