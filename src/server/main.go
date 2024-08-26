package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pixelive/api"

	"github.com/joho/godotenv"
)

func serve() {
	args := os.Args
	isDev := args[0] == "dev"

	hub := api.NewHub()
	go hub.Run()

	mux := http.NewServeMux()

	http.HandleFunc("/api/pixels", api.PixelsHandler)
	http.HandleFunc("/api/ws", func(w http.ResponseWriter, r *http.Request) {
		api.WSHandler(hub, w, r)
	})

	if !isDev {
		staticDir := "./dist"
		fileServer := http.FileServer(http.Dir(staticDir))

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join(staticDir, r.URL.Path)
			if _, err := os.Stat(path); os.IsNotExist(err) {
					http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
					return
			}
			fileServer.ServeHTTP(w, r)
	})
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serve()
}
