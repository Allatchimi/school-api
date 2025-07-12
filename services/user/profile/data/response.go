package data

// Update password
type UpdateProfilePasswordInitResponse struct {
	Token string `json:"token" required:"false" doc:"Token"`
}
type UpdateProfilePasswordCheckCodeResponse struct {
	Token string `json:"token" required:"false" doc:"Token"`
}

// Update phone number
type UpdateProfilePhoneNumberInitResponse struct {
	Token string `json:"token" required:"false" doc:"Token"`
}
type UpdateProfilePhoneNumberCheckCodeResponse struct {
	Token string `json:"token" required:"false" doc:"Token"`
}

// Update profile message
type UpdateProfileConfigMessageResponse struct {
	WhatsappPhoneNumber int64 `json:"whatsappPhoneNumber" required:"false" doc:"Whatsapp phone number"`
	TelegramChatID      int64 `json:"telegramChatID" required:"false" doc:"Telegram chat id"`
}

// Update MFA for email
type UpdateProfileMfaEmailInitResponse struct {
	Token string `json:"token" required:"false" doc:"Token"`
}

// Get web push subscription public key
type GetProfileWebPushSubscriptionResponse struct {
	PublicKey string `json:"publicKey" required:"false" doc:"Public key in base64"`
}
