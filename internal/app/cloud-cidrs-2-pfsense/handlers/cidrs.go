package handlers

import (
	"github.com/iteratec/cloud-cidrs-2-pfsense/internal/app/cloud-cidrs-2-pfsense/api"
	"github.com/iteratec/cloud-cidrs-2-pfsense/internal/app/cloud-cidrs-2-pfsense/service/aws"
	"github.com/iteratec/cloud-cidrs-2-pfsense/internal/app/cloud-cidrs-2-pfsense/service/gcp"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CloudCidrsApi struct {}

func (api CloudCidrsApi) FetchAwsCidrs(ctx echo.Context, params api.FetchAwsCidrsParams) error  {
	cidrs, err := aws.FetchAwsCidrs(params.Regions)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, cidrs)
}

func (api CloudCidrsApi) FetchGcpCidrs(ctx echo.Context) error  {
	cidrs, err := gcp.FetchGcpCidrs()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, cidrs)
}
