package main

import (
	"net/http"

	"github.com/iamhabbeboy/elastic-app/cmd/web/handlers"
	"github.com/iamhabbeboy/elastic-app/database/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	model.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.IndexHandler)
	http.ListenAndServe(":3000", r)
}
