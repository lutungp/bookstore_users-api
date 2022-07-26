package users_db

import (
	"database/sql"
	"fmt"
	"github/lutungp/bookstore_users-api/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	username = config.Config("DB_USERNAME")
	password = config.Config("DB_PASSWORD")
	host     = config.Config("DB_HOST")
	schema   = config.Config("DB_SCHEMA")
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
