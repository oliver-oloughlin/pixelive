package db

func Reset() error {
	db := get_db()
	_, err := db.Exec("UPDATE pixels SET color = 'transparent'")
	return err
}
