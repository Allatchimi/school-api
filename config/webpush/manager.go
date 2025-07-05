package configWebpush

import (
	"api/common/helpers"
	"api/config"
	"api/services/user/user"
	"api/services/user/user/model"
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
	user *model.User,
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
	users []*model.User,
	payload *WebPushPayload,
	ttl *int,
	userRepository *user.Repository,
) (errs []error) {
	errs = make([]error, 0, len(users))
	for _, user := range users {
		errs = append(errs, SendPushNotificationToUser(user, payload, ttl, userRepository))
	}
	return
}
