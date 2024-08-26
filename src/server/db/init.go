package db

import (
	"log"
	"strconv"
	"strings"
)

func Init() {
	db := get_db()
	query := "INSERT INTO pixels (id, color) VALUES "
	vals := []interface{}{}

	for i := 1; i <= 100*100; i++ {
		query += "($" + strconv.Itoa(i*2-1) + ",$" + strconv.Itoa(i*2) + "),"
		vals = append(vals, i-1, "transparent")
	}

	query = strings.TrimSuffix(query, ",")

	_, err := db.Exec(query, vals...)
	if err != nil {
		log.Fatalln("Failed exec:", err)
	}
}
