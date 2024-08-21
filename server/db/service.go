package db

func GetPixels() []Pixel {
	rows, err := database.Query("SELECT * FROM pixels")
	var pixels []Pixel

	if err != nil {
		return pixels
	}

	defer rows.Close()

	for rows.Next() {
		var pixel Pixel

		if err := rows.Scan(&pixel.X, &pixel.Y, &pixel.Color); err != nil {
			return pixels
		}

		pixels = append(pixels, pixel)
	}

	return pixels
}
