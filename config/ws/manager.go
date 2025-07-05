package configWS

import (
	"api/common/helpers"
	"api/config"
	"context"
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WSManager struct {
	connections map[string]map[*Client]bool // userID → set de connexions
	mu          sync.RWMutex
}

type Client struct {
	Conn   *websocket.Conn
	UserID string
}

type WSNotificationResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Href      string `json:"href"`
	Seen      bool   `json:"seen"`
	CreatedAt string `json:"createdAt"`
}

var WebsocketManager *WSManager

// Setup the websocket manager
func SetupWebsocket() *WSManager {
	if WebsocketManager == nil {
		WebsocketManager = &WSManager{
			connections: make(map[string]map[*Client]bool),
		}
	}
	return WebsocketManager
}

// Register a new client
func (m *WSManager) Register(userID string, conn *websocket.Conn) *Client {
	if m == nil {
		return nil
	}
	if conn == nil {
		return nil
	}
	if len(userID) < 1 {
		return nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	client := &Client{
		Conn:   conn,
		UserID: userID,
	}

	if m.connections[userID] == nil {
		m.connections[userID] = make(map[*Client]bool)
	}

	m.connections[userID][client] = true

	helpers.Logger.Info("New WS connection!", zap.String("UserID", userID))
	return client
}

// Unregister a client
func (m *WSManager) Unregister(client *Client) {
	if m == nil {
		return
	}
	if client == nil {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if clients, ok := m.connections[client.UserID]; ok {
		delete(clients, client)
		if len(clients) == 0 {
			delete(m.connections, client.UserID)
		}
		helpers.Logger.Info("WS connection closed!", zap.String("UserID", client.UserID))
	}
}

// Publish a notification to Redis
func PublishNotification(userID string, notif *WSNotificationResponse) error {
	msg := NotificationMessage{
		UserID:       userID,
		Notification: notif,
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return config.RedisClient.Publish(context.Background(), "notifications", payload).Err()
}
