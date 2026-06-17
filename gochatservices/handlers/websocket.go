package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"SistemadeChat/gochatservices/services"

	"github.com/gorilla/websocket"
)

var chatService = services.NewChatService()

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type VerifyResponse struct {
	Valid    bool   `json:"valid"`
	Username string `json:"username"`
}

func verifyToken(token string) (string, bool) {
	apiURL := "http://localhost:8000/verify-token?token=" + url.QueryEscape(token)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}

	var result VerifyResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", false
	}

	return result.Username, result.Valid
}
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	// 1. obtener token
	token := r.URL.Query().Get("token")

	if token == "" {
		http.Error(w, "Token requerido", http.StatusUnauthorized)
		return
	}

	// 2. validar token contra FastAPI
	username, valid := verifyToken(token)
	if !valid {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	// 3. upgrade a websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error websocket:", err)
		return
	}

	// 4. agregar cliente
	chatService.AddClient(conn)

	fmt.Println("Usuario conectado:", username)
	fmt.Println("Clientes conectados:", len(chatService.Clients))

	// 5. cleanup
	defer func() {
		chatService.RemoveClient(conn)
		fmt.Println("Usuario desconectado:", username)
	}()

	// 6. loop de mensajes
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println(username+":", string(message))

		// 7. mensaje estructurado (IMPORTANTE)
		msg := map[string]string{
			"user": username,
			"text": string(message),
		}

		// 8. broadcast correcto
		chatService.BroadcastJSON(msg)
	}
}
