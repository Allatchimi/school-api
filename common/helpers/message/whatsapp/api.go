package whatsappHelper

import (
	"api/common/constants"
	httpHelper "api/common/helpers/http"
	"fmt"
)

const (
	version = "v22.0"
)

// Send a WhatsApp Utility template message via Cloud API
func postTemplateMessage(job *WhatsAppJob) (any, error) {
	// Basic input validation
	if job == nil || job.PhoneID == "" || job.AccessToken == "" ||
		job.ReceiverPhoneNumber == "" || job.Template == "" {
		return nil, fmt.Errorf("invalid inputs")
	}
	if job.Language == "" {
		job.Language = constants.LANGUAGE_ENGLISH
	}
	if job.TTLSeconds < 60 || job.TTLSeconds > 3600 {
		job.TTLSeconds = 3600
	}

	// Build parameters array for the BODY component
	params := make([]map[string]any, 0, len(job.BodyParams))
	for _, v := range job.BodyParams {
		params = append(params, map[string]any{
			"type": "text",
			"text": v,
		})
	}

	// Build template object
	template := map[string]any{
		"name": job.Template,
		"language": map[string]any{
			"code": job.Language,
		},
		"components": []map[string]any{
			{
				"type":       "body",
				"parameters": params,
			},
		},
	}

	payload := map[string]any{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                job.ReceiverPhoneNumber, // E.164 format
		"type":              "template",
		"template":          template,
	}

	// Add TTL for Utility (Cloud API supports 60–3600 seconds). Omit if 0.
	if job.TTLSeconds >= 60 && job.TTLSeconds <= 3600 {
		payload["message_send_ttl_seconds"] = job.TTLSeconds
	}

	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", version, job.PhoneID)
	headers := []httpHelper.HttpHeader{
		{Label: "Authorization", Value: "Bearer " + job.AccessToken},
		{Label: "Content-Type", Value: "application/json"},
	}

	var resp any
	if err := httpHelper.HttpPost(url, headers, payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
