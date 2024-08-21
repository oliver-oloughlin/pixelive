package db

import (
	"database/sql"
	"fmt"
	"os"
)

func connect() *sql.DB {
	db, err := sql.Open("libsql", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
	return db
}

var database = connect()
