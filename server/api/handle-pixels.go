package api

import "net/http"

func PixelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
