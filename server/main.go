package main

import (
	"log"
	"net/http"
	"pixelive/api"

	"github.com/joho/godotenv"
)

func serve() {
	hub := api.NewHub()
	go hub.Run()

	http.HandleFunc("/pixels", api.PixelsHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		api.WSHandler(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serve()
}
