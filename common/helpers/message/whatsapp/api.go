package whatsappHelper

import (
	httpHelper "api/common/helpers/http"
	"fmt"
)

func postHttpMessage(job *WhatsAppJob) (data any, err error) {
	url := fmt.Sprintf("https://graph.facebook.com/v17.0/%s/messages", job.PhoneID)
	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                job.ReceiverPhoneNumber,
		"type":              "text",
		"text": map[string]string{
			"body": job.Message,
		},
	}
	headers := []httpHelper.HttpHeader{
		{Label: "Authorization", Value: "Bearer " + job.AccessToken},
		{Label: "Content-Type", Value: "application/json"},
	}
	err = httpHelper.HttpPost(url, headers, payload, data)
	return
}
