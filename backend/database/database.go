package database

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/iamhabbeboy/elastic-app/ent"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DBPostgres struct {
	client *ent.Client
}

func (d *DBPostgres) GetDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.Get("DB.USER"),
		viper.Get("DB.PASS"),
		viper.Get("DB.HOST"),
		viper.Get("DB.PORT"),
		viper.Get("DB.NAME"),
	)
}
func (d *DBPostgres) Open() (*ent.Client, error) {
	return ent.Open(dialect.MySQL, d.GetDSN())
}

func DBConnection() (*ent.Client, error) {
	db := &DBPostgres{}
	client, err := db.Open()
	if err != nil {
		return nil, err
	}
	db.client = client
	db.RunMigration()
	defer client.Close()
	return client, nil
}

func (db *DBPostgres) RunMigration() {
	if err := db.client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
