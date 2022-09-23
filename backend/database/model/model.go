package model

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iamhabbeboy/elastic-app/database"
	"github.com/iamhabbeboy/elastic-app/ent"
	"github.com/spf13/viper"
)

var (
	dbConn *ent.Client

	ctx context.Context
)

func Init() {
	setupConf()
	client, err := database.DBConnection()
	if err != nil {
		panic(err)
	}
	dbConn = client
}

func setupConf() {
	viper.SetConfigFile("app.yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

}
