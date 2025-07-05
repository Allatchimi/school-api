package configWS

import (
	"api/common/helpers"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Listen to websocket connection
func (m *WSManager) Listen(c *gin.Context) {
	if m == nil {
		helpers.Logger.Error("No available WS manager!")
		return
	}

	// Check if the user is authenticated
	jwtToken, err := helpers.GetJwtContextFromQuery(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Get user id from query
	userID := fmt.Sprintf("%d", jwtToken.UserID)
	if len(userID) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID! Please enter valid information."})
		return
	}

	// Upgrade to websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		helpers.Logger.Error("Error upgrading websocket!", zap.String("Error", err.Error()))
		return
	}
	defer conn.Close()

	// Register user to manager
	client := m.Register(userID, conn)
	defer m.Unregister(client)

	// Handle Pong to keep connection alive
	conn.SetPongHandler(func(appData string) error {
		helpers.Logger.Info("Received pong!", zap.String("UserID", userID))
		return nil
	})

	// Send Ping every 30 seconds
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				helpers.Logger.Warn("Ping failed!", zap.String("UserID", userID), zap.String("Error", err.Error()))
				return
			}
		}
	}()

	// Loop to keep the connection open
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			helpers.Logger.Warn("Connection closed!", zap.String("UserID", userID), zap.String("Error", err.Error()))
			break
		}
	}
}
