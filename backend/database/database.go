package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/iamhabbeboy/elastic-app/ent"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DBPostgres struct {
	client *ent.Client
}

// func DBConnection() (*ent.Client, error) {
// 	client, err := ent.Open("mysql", "elastic:secret@tcp(mysql:3306)/elastic")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return client, nil
// }

func (d *DBPostgres) GetDSN() string {
	DBPort, ok := viper.Get("DB.PORT").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	fmt.Printf("viper : %s = %s \n", "Database Port", DBPort)
	// user := "elastic"
	pass := "secret"
	host := "mysql"
	db := "elastic"
	port := "3306"
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		pass,
		host,
		port,
		db,
	) //os.Getenv("DB_USER"),
	//os.Getenv("DB_PASS"),
	//os.Getenv("DB_HOST"),
	//os.Getenv("DB_PORT"),
	//os.Getenv("DB_NAME"),

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
