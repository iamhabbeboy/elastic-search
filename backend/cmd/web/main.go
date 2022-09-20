package main

import (
	"fmt"
	"net/http"

	"github.com/iamhabbeboy/elastic-app/cmd/web/handlers"
	"github.com/iamhabbeboy/elastic-app/database/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/Users/solomon/Work/elastic-search/backend/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	model.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.IndexHandler)
	http.ListenAndServe(":3000", r)
}
