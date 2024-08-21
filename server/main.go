package main

import (
	"log"
	"net/http"
	"pixelive/api"
)

func main() {
	http.HandleFunc("/pixels", api.PixelsHandler)
	http.HandleFunc("/ws", api.WSHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
