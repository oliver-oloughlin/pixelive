package api

import (
	"encoding/json"
	"log"
	"pixelive/db"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("client registered %s", client.id)

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Printf("client unregistered %s", client.id)
				close(client.send)
				delete(h.clients, client)
			}

		case msg := <-h.broadcast:
			go func() {
				var pixel db.Pixel
				json.Unmarshal(msg, &pixel)
				db.UpdatePixel(pixel)
			}()

			for client := range h.clients {
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
