package main

import (
	"fmt"
	"net/http"

	"github.com/iamhabbeboy/elastic-app/database/model"
	"github.com/iamhabbeboy/elastic-app/service"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("app.yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	model.Init()

	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		ec := service.NewElasticSearch()
		_ = ec.GetInfo()

		return c.String(http.StatusOK, "Hello world")
	})
	fmt.Println("Started Listening on :3000 ......")
	e.Logger.Fatal(e.Start(":3000"))
}
