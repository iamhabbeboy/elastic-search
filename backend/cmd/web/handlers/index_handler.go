package handlers

import (
	"net/http"

	"github.com/iamhabbeboy/elastic-app/service"
	"github.com/labstack/echo/v4"
)

var (
	r map[string]interface{}
)

func IndexHandler(c echo.Context) error {
	ec := service.NewElasticSearch()
	err := ec.GetInfo()
	if err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}
	return c.JSON(http.StatusOK, "We are good to go")
}
