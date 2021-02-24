package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CloudCidrsApi struct {}

func (api CloudCidrsApi) FetchAwsCidrs(ctx echo.Context) error  {
	return ctx.String(http.StatusOK, "FetchAwsCidrs")
}

func (api CloudCidrsApi) FetchGcpCidrs(ctx echo.Context) error  {
	return ctx.String(http.StatusOK, "FetchGcpCidrs")
}
