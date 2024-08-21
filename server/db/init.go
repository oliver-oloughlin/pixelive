package db

import (
	"log"
	"strings"
)

func Init() {
	db := get_db()
	query := "INSERT INTO pixels (id, color) VALUES "
	vals := []interface{}{}

	for i := 0; i < 100*100; i++ {
		query += "(?, ?),"
		vals = append(vals, i, "transparent")
	}

	query = strings.TrimSuffix(query, ",")

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln("Failed prepare:", err)
	}

	_, err = stmt.Exec(vals...)
	if err != nil {
		log.Fatalln("Failed exec:", err)
	}
}
