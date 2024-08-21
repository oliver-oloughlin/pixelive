package main

import (
	"log"
	"net/http"
	"pixelive/api"

	"github.com/joho/godotenv"
)

func serve() {
	http.HandleFunc("/pixels", api.PixelsHandler)
	http.HandleFunc("/ws", api.WSHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serve()
}
