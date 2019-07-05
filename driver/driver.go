package driver

import (
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"

	_ "github.com/lib/pq"
)

const driver = "postgres"

type dbConfig struct {
	Db            string
	Host          string
	Port          int
	User          string
	Password      string
	Params        string
	MongoHost     string
	MongoPort     string
	MongoUser     string
	MongoPass     string
	MongoDatabase string
	SlackToken    string
}

var dbc dbConfig

func ConnectDB() *sql.DB {
	err := envconfig.Process("PG", &dbc)
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", dbc.User, dbc.Password, dbc.Host, dbc.Port, dbc.Db)
	//db, err := sqlx.Connect(driver, dsn)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
