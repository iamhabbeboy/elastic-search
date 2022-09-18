package model

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iamhabbeboy/elastic-app/database"
	"github.com/iamhabbeboy/elastic-app/ent"
)

var (
	dbConn *ent.Client

	ctx context.Context
)

func Init() {
	client, err := database.DBConnection()
	if err != nil {
		panic(err)
	}
	dbConn = client
}
