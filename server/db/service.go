package db

func GetPixels() []Pixel {
	db := get_db()
	rows, err := db.Query("SELECT * FROM pixels")
	pixels := []Pixel{}

	if err != nil {
		return pixels
	}

	defer rows.Close()

	for rows.Next() {
		var pixel Pixel

		if err := rows.Scan(&pixel.ID, &pixel.Color); err != nil {
			return pixels
		}

		pixels = append(pixels, pixel)
	}

	return pixels
}

func UpdatePixel(pixel Pixel) error {
	db := get_db()
	_, err := db.Exec("UPDATE pixels SET color = $1 WHERE id = $2", pixel.Color, pixel.ID)
	return err
}
