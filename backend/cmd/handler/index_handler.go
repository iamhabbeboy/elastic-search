package handler

import (
	"fmt"
	"net/http"

	"github.com/iamhabbeboy/elastic-app/cmd/service"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	res, err := service.NewElasticSearch().Ping()
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())

	}
	return c.String(http.StatusOK, fmt.Sprintf("Elasticsearch version %s\n", res.Version.Number))
}
