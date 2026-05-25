package routes

import (
	"net/http"

	"Sistema-de-Chat/go-chat-service/handlers"
)

func SetupRoutes() {

	http.HandleFunc("/ws", handlers.HandleWebSocket)
}
