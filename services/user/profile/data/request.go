package data

import "time"

// Update profile information
type UpdateProfileInfoRequest struct {
	Username      string     `json:"username" required:"false" minLength:"2" maxLength:"30" doc:"User name"`
	FirstName     string     `json:"firstName" required:"true" minLength:"2" maxLength:"30" doc:"First name"`
	LastName      string     `json:"lastName" required:"true" doc:"Last name"`
	Gender        string     `json:"Gender" required:"true" enum:"male,female" doc:"Gender"`
	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Address       string     `json:"address" required:"false" doc:"Address"`
	Language      string     `json:"language" required:"false" doc:"Language code with 2 letter"`
	Image         string     `json:"image" required:"false" doc:"Thumbnail"`
}

// Update password
type UpdateProfilePasswordCheckCodeRequest struct {
	Token string `json:"token" required:"true" doc:"Received token on previous step"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email"`
}
type UpdateProfilePasswordNewPasswordRequest struct {
	Token           string `json:"token" required:"true" doc:"Received token on previous step"`
	CurrentPassword string `json:"currentPassword" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
	NewPassword     string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
}

// Update phone number
type UpdateProfilePhoneNumberCheckCodeRequest struct {
	Token string `json:"token" required:"true" doc:"Received token on previous step"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email"`
}
type UpdateProfilePhoneNumberNewPhoneNumberRequest struct {
	Token       string `json:"token" required:"true" doc:"Received token on previous step"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number"`
}

// Update message
type UpdateProfileMessageRequest struct {
	WhatsappPhoneNumber int64 `json:"whatsappPhoneNumber" required:"false" doc:"Whatsapp phone number"`
	TelegramChatID      int64 `json:"telegramChatID" required:"false" doc:"Telegram chat id"`
}

// Update MFA for email
type UpdateProfileMfaEmailCheckCodeRequest struct {
	Token string `json:"token" required:"true" doc:"Received token on previous step"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email"`
}

// Update notification setting
type UpdateProfileSettingNotificationRequest struct {
	IsEnabled bool `json:"isEnabled" required:"true" doc:"Is enabled"`
}

// Subscribe to web push notification
type UpdateProfileWebPushSubscriptionRequest struct {
	Endpoint string `json:"endpoint" required:"true" doc:"Endpoint"`
	Keys     *struct {
		P256dh string `json:"p256dh" required:"true" doc:"P256dh"`
		Auth   string `json:"auth" required:"true" doc:"Auth"`
	} `json:"keys" required:"true" doc:"Keys"`
}
