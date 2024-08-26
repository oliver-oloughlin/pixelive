package api

import (
	"encoding/json"
	"net/http"
	"pixelive/src/server/db"
)

func PixelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pixels := db.GetPixels()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pixels)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
