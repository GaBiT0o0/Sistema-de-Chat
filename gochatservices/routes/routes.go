package routes

import (
	"SistemadeChat/gochatservices/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/ws", handlers.HandleWebSocket)
}
