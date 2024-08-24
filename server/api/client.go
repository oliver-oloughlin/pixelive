package api

import (
	"encoding/json"
	"log"
	"net/http"
	"pixelive/db"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id   string
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func WSHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := &Client{
		id:   uuid.New().String(),
		hub:  hub,
		conn: conn,
		send: make(chan []byte),
	}

	client.hub.register <- client

	go client.writePump()
	go client.readPump()

	// for {
	// 	mt, message, err := c.ReadMessage()

	// 	if err != nil {
	// 		log.Println("read:", err)
	// 		delete(connections, id)
	// 		break
	// 	}

	// 	var pixel db.Pixel
	// 	json.Unmarshal(message, &pixel)
	// 	err = db.UpdatePixel(pixel)

	// 	if err != nil {
	// 		log.Println("update:", err)
	// 		delete(connections, id)
	// 		break
	// 	}

	// 	for connId, connection := range connections {
	// 		err := connection.WriteMessage(mt, message)
	// 		if err != nil {
	// 			log.Println("write:", err)
	// 			delete(connections, connId)
	// 			continue
	// 		}
	// 	}
	// }
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		c.hub.unregister <- c
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("read error: %v", err)
			break
		}

		var pixel db.Pixel
		json.Unmarshal(msg, &pixel)

		if err := db.UpdatePixel(pixel); err != nil {
			continue
		}

		c.hub.broadcast <- msg
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add((writeWait)))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(msg)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(msg)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
