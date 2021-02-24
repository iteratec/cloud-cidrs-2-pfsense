package handlers

import (
	"github.com/iteratec/cloud-cidrs-2-pfsense/internal/app/cloud-cidrs-2-pfsense/service/aws"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CloudCidrsApi struct {}

func (api CloudCidrsApi) FetchAwsCidrs(ctx echo.Context) error  {
	cidrs, err := aws.FetchAwsCidrs()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, cidrs)
}

func (api CloudCidrsApi) FetchGcpCidrs(ctx echo.Context) error  {
	return ctx.String(http.StatusOK, "FetchGcpCidrs")
}
