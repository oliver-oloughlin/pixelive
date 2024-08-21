package db

func GetPixels() []Pixel {
	db := get_db()
	rows, err := db.Query("SELECT * FROM pixels")
	var pixels []Pixel

	if err != nil {
		return pixels
	}

	defer rows.Close()

	for rows.Next() {
		var pixel Pixel

		if err := rows.Scan(&pixel.id, &pixel.color); err != nil {
			return pixels
		}

		pixels = append(pixels, pixel)
	}

	return pixels
}

func UpdatePixel(pixel Pixel) {
	db := get_db()
	db.Exec("UPDATE pixels SET color = ? WHERE id = ?", pixel.color, pixel.id)
}
