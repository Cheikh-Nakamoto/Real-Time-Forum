package handlers

import (
	"bim/internal/config"
	"bim/internal/metier/database"
	"bim/internal/metier/dto"
	"bim/internal/metier/entities"
	"bim/internal/metier/services/implementations"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var clients = make(map[uint]*websocket.Conn)
var clientsMutex sync.Mutex

func SocketMessage(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	header, payload, signature := config.SplitToken(token)
	if !config.CheckToken(header, payload, signature, os.Getenv("JWT_SECRET")) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		http.Error(w, "Invalid token payload", http.StatusUnauthorized)
		return
	}

	var payloadData map[string]interface{}
	err = json.Unmarshal(payloadBytes, &payloadData)
	if err != nil {
		http.Error(w, "Invalid token payload", http.StatusUnauthorized)
		return
	}

	userID, ok := payloadData["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid token payload", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusInternalServerError)
		return
	}

	db, err := database.InitDatabase()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	var (
		messageService implementations.MessageService
	)
	messageService.SetDB(db.GetDatabase())

	clientsMutex.Lock()
	clients[uint(userID)] = conn
	clientsMutex.Unlock()

	for {
		var (
			msg    dto.MessageDto
			toSave entities.Message
		)
		err := conn.ReadJSON(&msg)
		if err != nil {
			clientsMutex.Lock()
			delete(clients, uint(userID))
			clientsMutex.Unlock()
			conn.Close()
			break
		}

		toSave.SetSenderID(msg.SenderID)
		toSave.SetReceiverID(msg.ReceiverID)
		toSave.SetContent(msg.Content)

		err = messageService.SaveMessage(toSave)
		if err != nil {
			http.Error(w, "Error saving message", http.StatusInternalServerError)
			return
		}

		clientsMutex.Lock()
		if receiverConn, ok := clients[msg.ReceiverID]; ok {
			receiverConn.WriteJSON(msg)
		}
		clientsMutex.Unlock()
	}
}
