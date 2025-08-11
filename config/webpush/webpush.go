package webpushConfig

import (
	"api/common/helpers"
	"api/config"
	wsConfig "api/config/ws"
	"api/services/others/notification"
	modelNotification "api/services/others/notification/model"
	"api/services/user/user"
	modelUser "api/services/user/user/model"

	"encoding/json"
	"fmt"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"go.uber.org/zap"
)

type WebPushSubscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

type WebPushPayload struct {
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Icon      string     `json:"icon"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"createdAt"`
}

// SendPushNotificationToUser sends a push notification
// to a user with a given payload and TTL
func SendPushNotificationToUser(
	user *modelUser.User,
	payload *WebPushPayload,
	ttl *int,
	userRepository *user.Repository,
) error {
	// Validate user config
	if user == nil || user.Config == nil ||
		len(user.Config.WebPushSubscriptionEndpoint) < 1 ||
		len(user.Config.WebPushSubscriptionKeyP256dh) < 1 ||
		len(user.Config.WebPushSubscriptionKeyAuth) < 1 {
		helpers.Logger.Warn("Missing subscription information for user! Skipped!")
		return nil
	}

	// Convert payload to JSON string
	now := time.Now()
	payloadValue := *payload
	payloadValue.CreatedAt = &now
	payloadBytes, err := json.Marshal(payloadValue)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Set TTL with default
	ttlValue := 604800 // Default 7 days
	if ttl != nil || ttlValue < 0 {
		ttlValue = *ttl
	}

	// Send notification
	subscription := &webpush.Subscription{
		Endpoint: user.Config.WebPushSubscriptionEndpoint,
		Keys: webpush.Keys{
			P256dh: user.Config.WebPushSubscriptionKeyP256dh,
			Auth:   user.Config.WebPushSubscriptionKeyAuth,
		},
	}
	resp, err := webpush.SendNotification(
		payloadBytes,
		subscription,
		&webpush.Options{
			Subscriber:      config.Env.WebPushSubscriber,
			VAPIDPublicKey:  config.Env.WebPushVapidPublicKey,
			VAPIDPrivateKey: config.Env.WebPushVapidPrivateKey,
			TTL:             ttlValue,
		})
	if err != nil {
		return fmt.Errorf("failed to send push notification: %w", err)
	}
	defer resp.Body.Close()

	// Check if the subscription is expired or deleted from client side
	if resp.StatusCode == 410 || resp.StatusCode == 404 {
		helpers.Logger.Warn("Subscription expired or deleted from client side!", zap.Int64("userID", user.ID))
		helpers.Logger.Warn("Deleting subscription from database...", zap.Int64("userID", user.ID))
		// Delete subscription from database
		if userRepository != nil {
			_, errDelete := userRepository.UpdateUserConfigWebPushSubscriptionByID(user.ID, "", "", "")
			if errDelete != nil {
				helpers.Logger.Error("Failed to delete subscription from database!", zap.Int64("userID", user.ID))
			}
			helpers.Logger.Warn("Subscription deleted from database!", zap.Int64("userID", user.ID))
		}
	}

	return nil
}

// SendPushNotificationToUserBulk sends a push notification
// to multiple users with a given payload and TTL
func SendPushNotificationToUserBulk(
	users []modelUser.User,
	payload *WebPushPayload,
	ttl *int,
	userRepository *user.Repository,
	notificationRepository *notification.Repository,
) (errs []error) {
	if payload == nil || userRepository == nil || notificationRepository == nil {
		return
	}

	errs = make([]error, 0, len(users))
	for _, user := range users {
		if user.ID < 1 {
			continue
		}

		helpers.Logger.Info("Sending notification message to user: ", zap.Int64("userID", user.ID), zap.String("email", user.Email))

		// Save in database
		notification, err := notificationRepository.Create(&modelNotification.Notification{
			UserID:  user.ID,
			Title:   payload.Title,
			Message: payload.Body,
		})
		if err != nil {
			errs = append(errs, err)
			return
		}

		if notification != nil {
			// Publish WS notification
			wsConfig.PublishNotification(
				fmt.Sprintf("%d", user.ID),
				&wsConfig.WSNotificationResponse{
					ID:        fmt.Sprintf("%d", notification.ID),
					Audience:  fmt.Sprintf("%d", user.ID),
					Title:     notification.Title,
					Message:   notification.Message,
					Seen:      notification.Seen,
					CreatedAt: fmt.Sprintf("%s", notification.CreatedAt.String()),
					Href:      "",
				},
			)
		}

		// Send push notification
		if user.Config.AllowNotifications {
			errs = append(errs, SendPushNotificationToUser(&user, payload, ttl, userRepository))
		}
	}
	return
}
