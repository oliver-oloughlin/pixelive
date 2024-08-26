package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pixelive/src/server/api"

	"github.com/joho/godotenv"
)

func serve() {
	args := os.Args
	isDev := args[1] == "dev"
	mux := http.NewServeMux()

	hub := api.NewHub()
	go hub.Run()

	if isDev {
		mux.Handle("/api", api.EnableCORS(mux))
	}

	mux.HandleFunc("/api/pixels", api.PixelsHandler)
	mux.HandleFunc("/api/ws", func(w http.ResponseWriter, r *http.Request) {
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

	log.Fatal(http.ListenAndServe(":8000", mux))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serve()
}
