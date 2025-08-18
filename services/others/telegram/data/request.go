package data

type TelegramBotSchoolID struct {
	SchoolID int64 `json:"schoolID" path:"schoolID" required:"true" doc:"School id"`
}

type TelegramWebhookRequest map[string]any

// type TelegramWebhookRequest struct {
// 	UpdateID int64 `json:"update_id" required:"false" doc:"Update id"`
// 	Message  struct {
// 		MessageID int64 `json:"message_id" required:"false" doc:"Message id"`
// 		From      struct {
// 			ID           int64  `json:"id" required:"false" doc:"User id"`
// 			IsBot        bool   `json:"is_bot" required:"false" doc:"Is bot"`
// 			FirstName    string `json:"first_name" required:"false" doc:"First name"`
// 			Username     string `json:"username" required:"false" doc:"Username"`
// 			LanguageCode string `json:"language_code" required:"false" doc:"Language code"`
// 		} `json:"from" required:"false" doc:"From"`
// 		Chat struct {
// 			ID        int64  `json:"id" required:"false" doc:"Chat id"`
// 			FirstName string `json:"first_name" required:"false" doc:"First name"`
// 			Username  string `json:"username" required:"false" doc:"Username"`
// 			Type      string `json:"type" required:"false" doc:"Type"`
// 		} `json:"chat" required:"false" doc:"Chat"`
// 		Date int64  `json:"date" required:"false" doc:"Date"`
// 		Text string `json:"text" required:"false" doc:"Text"`
// 	} `json:"message" required:"false" doc:"Message"`
// }
