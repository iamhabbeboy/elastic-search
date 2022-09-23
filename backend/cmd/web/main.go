package main

import (
	"github.com/iamhabbeboy/elastic-app/cmd/web/router"
	"github.com/iamhabbeboy/elastic-app/database/model"
)

func main() {
	model.Init()

	e := router.DefineRoute()
	e.Logger.Fatal(e.Start(":3000"))
}
