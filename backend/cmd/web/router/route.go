package router

import (
	"github.com/iamhabbeboy/elastic-app/cmd/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DefineRoute() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthstatus", handler.HealthCheckHandler)

	return e
}
