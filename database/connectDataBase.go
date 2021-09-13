package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDataBase() *sql.DB {
	connectionString := "user=postgres dbname=postgres password=admin host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return db
}
