package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB

	username = "root"
	password = "P%40ssw0rd"
	host     = "localhost"
	port     = "15432"
	database = "db"
)

func init() {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	var error error
	Client, error = sql.Open("postgres", dataSourceName)
	if error != nil {
		panic(error)
	}

	if error = Client.Ping(); error != nil {
		panic(error)
	}
}
