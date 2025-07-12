package wsConfig

import (
	"api/common/helpers"
	"api/config"
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

type NotificationMessage struct {
	UserID       string                  `json:"user_id"`
	Notification *WSNotificationResponse `json:"notification"`
}

// Send a notification to a specific user
func (m *WSManager) sendNotification(userID string, notification *WSNotificationResponse) {
	if m == nil {
		return
	}
	if len(userID) < 1 {
		return
	}
	if notification == nil {
		return
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	clients, ok := m.connections[userID]
	if !ok {
		helpers.Logger.Warn("No active connection for user!", zap.String("UserID", userID))
		return
	}

	for client := range clients {
		if err := client.Conn.WriteJSON(notification); err != nil {
			helpers.Logger.Error("Error sending notification!", zap.String("Error", err.Error()))
		}
	}
}

// Subscribe to notifications from Redis in order
// to sync notifications on multiple websocket servers
func (m *WSManager) SubscribeToNotifications() {
	pubsub := config.RedisClient.Subscribe(context.Background(), "notifications")

	go func() {
		for msg := range pubsub.Channel() {
			var notifMsg NotificationMessage
			if err := json.Unmarshal([]byte(msg.Payload), &notifMsg); err != nil {
				helpers.Logger.Error("Error parsing message from Redis!", zap.String("Error", err.Error()))
				continue
			}

			m.sendNotification(notifMsg.UserID, notifMsg.Notification)
		}
	}()
}
