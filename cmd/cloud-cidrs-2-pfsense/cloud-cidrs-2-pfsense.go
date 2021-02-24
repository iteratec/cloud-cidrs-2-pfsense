package main

import (
	"github.com/iteratec/cloud-cidrs-2-pfsense/internal/app/cloud-cidrs-2-pfsense/api"
	"github.com/iteratec/cloud-cidrs-2-pfsense/internal/app/cloud-cidrs-2-pfsense/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	sharedHandlers "github.com/iteratec/cloud-cidrs-2-pfsense/internal/pkg/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	registerRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func registerRoutes(e *echo.Echo) {
	registerGeneratedApi(e)
	e.GET("/swagger.json", sharedHandlers.GetSwaggerJsonHandler(api.GetSwagger))
}

func registerGeneratedApi(e *echo.Echo) {
	var cloudCidrsApi handlers.CloudCidrsApi
	api.RegisterHandlers(e, cloudCidrsApi)
}

