package whatsappHelper

import (
	httpHelper "api/common/helpers/http"
	"api/services/user/user/model"
	"fmt"
)

type WhatsAppClient struct {
	Token   string
	PhoneID string
}

var instances map[string]*WhatsAppClient

func NewClient(token, phoneID string) *WhatsAppClient {
	if instances == nil {
		instances = make(map[string]*WhatsAppClient)
	}
	if _, ok := instances[phoneID]; !ok {
		instances[phoneID] = &WhatsAppClient{
			Token:   token,
			PhoneID: phoneID,
		}
	}
	return instances[phoneID]
}

func (w *WhatsAppClient) SendMessage(users []model.User, message string) (data any, err error) {
	if w == nil {
		errMsg := "Whatsapp client is not instanced!"
		return nil, fmt.Errorf("%s", errMsg)
	}
	for _, user := range users {
		payload := map[string]any{
			"messaging_product": "whatsapp",
			"to":                user.Config.WhatsappPhoneNumber,
			"type":              "text",
			"text": map[string]string{
				"body": message,
			},
		}
		url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", w.PhoneID)
		headers := []httpHelper.HttpHeader{
			{Label: "Authorization", Value: "Bearer " + w.Token},
			{Label: "Content-Type", Value: "application/json"},
		}
		data = nil
		err := httpHelper.HttpPost(url, headers, payload, data)
		if err != nil {
			errMsg := "Failed to send message!"
			return nil, fmt.Errorf("%s: %d %w", errMsg, user.Config.WhatsappPhoneNumber, err)
		}
	}
	return nil, nil
}
