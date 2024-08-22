package api

import (
	"encoding/json"
	"log"
	"net/http"
	"pixelive/db"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = make(map[string]*websocket.Conn)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	id := uuid.New().String()
	connections[id] = c

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			connections[id] = nil
			break
		}

		var pixel db.Pixel
		json.Unmarshal(message, &pixel)
		err = db.UpdatePixel(pixel)

		if err != nil {
			log.Println("update:", err)
			connections[id] = nil
			break
		}

		for connId, connection := range connections {
			err := connection.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				connections[connId] = nil
				break
			}
		}
	}
}
