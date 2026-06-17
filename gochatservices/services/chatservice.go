package services

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type ChatService struct {
	Clients map[*websocket.Conn]bool
	Mutex   sync.Mutex
}

func NewChatService() *ChatService {
	return &ChatService{
		Clients: make(map[*websocket.Conn]bool),
	}
}

// agregar cliente
func (c *ChatService) AddClient(conn *websocket.Conn) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Clients[conn] = true
}

// eliminar cliente
func (c *ChatService) RemoveClient(conn *websocket.Conn) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	delete(c.Clients, conn)
	conn.Close()
}

// enviar mensaje a todos (JSON seguro)
func (c *ChatService) BroadcastJSON(msg interface{}) {
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for client := range c.Clients {
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			client.Close()
			delete(c.Clients, client)
		}
	}
}
