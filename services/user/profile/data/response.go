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

// Update MFA for email
type UpdateProfileMfaEmailInitResponse struct {
	Token string `json:"token" required:"false" doc:"Token"`
}

// Get web push subscription public key
type GetProfileWebPushSubscriptionResponse struct {
	PublicKey string `json:"publicKey" required:"false" doc:"Public key in base64"`
}
