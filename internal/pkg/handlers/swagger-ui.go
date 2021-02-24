package handlers

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
)

//Provides echo handler providing swagger-ui json.
func GetSwaggerJsonHandler(swaggerJsonFunc func() (*openapi3.Swagger, error)) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		swagger, err := swaggerJsonFunc()
		if err != nil {
			ctx.Logger().Error("Could not initialize swagger-json.")
		}
		jsonBytes, err := swagger.MarshalJSON()
		if err != nil {
			ctx.Logger().Error("Could not initialize swagger-json.")
		}
		return ctx.JSONBlob(http.StatusOK, jsonBytes)
	}
}
