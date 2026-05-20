package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Mensaje recibido:", string(message))

		conn.WriteMessage(websocket.TextMessage, message)
	}
}
