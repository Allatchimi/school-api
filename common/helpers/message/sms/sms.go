package smsHelper

import (
	"api/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SendSMS Sends a SMS to a specified receiver
func SendSMS(senderID string, message string, receiver string) error {
	if len(senderID) < 1 {
		errMsg := "Sender ID is empty!"
		return fmt.Errorf("%s", errMsg)
	}
	if len(message) < 1 {
		errMsg := "Message is empty!"
		return fmt.Errorf("%s", errMsg)
	}
	if len(receiver) < 1 {
		errMsg := "Receiver is empty!"
		return fmt.Errorf("%s", errMsg)
	}

	url := "https://api.africastalking.com/version1/messaging"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("apiKey", config.Env.SmsAfrikaTalkingApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	body := map[string]interface{}{
		"username": "username",
		"message":  message,
		"senderId": senderID,
		"phoneNumbers": []string{
			receiver,
		},
	}
	jsonData, _ := json.Marshal(body)
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))
	// curl -X POST \
	// https://api.africastalking.com/version1/messaging/bulk \
	// -H 'Accept: application/json' \
	// -H 'Content-Type: application/json' \
	// -H 'apiKey: MyAppApiKey' \
	// -d '{
	// "username": "username",
	// "message": "This is a sample message.",
	// "senderId": "ABC",
	// "phoneNumbers": [
	//     "+254711XXXYYY",
	//     "+254711YYYZZZ"
	// ]
	// }'
	return nil
}

// SendBulkSMS Send a SMS to multiple recipients
func SendBulkSMS(senderID string, message string, recipients []string) error {
	if len(senderID) < 1 {
		errMsg := "Sender ID is empty!"
		return fmt.Errorf("%s", errMsg)
	}
	if len(message) < 1 {
		errMsg := "Message is empty!"
		return fmt.Errorf("%s", errMsg)
	}
	if len(recipients) < 1 {
		errMsg := "Recipients is empty!"
		return fmt.Errorf("%s", errMsg)
	}
	return nil
}
