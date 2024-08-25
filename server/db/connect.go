package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connect() *sql.DB {
	connection_str := os.Getenv("DB_CONNECTION_STRING")
	database, err := sql.Open("postgres", connection_str)
	if err != nil {
		log.Fatalf(os.Stderr.Name(), "failed to open db %s", err)
	}
	return database
}

func get_db() *sql.DB {
	if db == nil {
		db = connect()
	}

	return db
}
