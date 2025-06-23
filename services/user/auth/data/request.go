package data

// Login
type LoginDevice struct {
	Platform   string `json:"platform" required:"true" minLength:"2" maxLength:"30" doc:"Platform name"`
	DeviceName string `json:"deviceName" required:"true" minLength:"2" maxLength:"50" doc:"Device name"`
	App        string `json:"app" required:"true" minLength:"2" maxLength:"50" doc:"Application used to login"`
}
type LoginWithEmailRequest struct {
	Email         string `json:"email" required:"true" format:"email" doc:"Email"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}
type LoginWithPhoneNumberRequest struct {
	PhoneNumber   uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}
type LoginWithProviderRequest struct {
	Provider string `json:"provider" required:"true" doc:"Provider" minLength:"2" maxLength:"30"`
	Token    string `json:"token" required:"true" doc:"Token" minLength:"3"`
}
type LoginRequest struct {
	Email         string `json:"email" required:"true" format:"email" doc:"Email"`
	PhoneNumber   uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}

// Register
type RegisterWithEmailRequest struct {
	Email    string `json:"email" required:"true" format:"email" doc:"Email"`
	Password string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
}
type RegisterWithPhoneNumberRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
	Password    string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
}
type RegisterRequest struct {
	Email       string `json:"email" required:"true" format:"email" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
	Password    string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
}

// Activate account
type ActivateAccountRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number"`
}

// Forgot password
type ForgotPasswordWithEmailInitRequest struct {
	Email string `json:"email" required:"true" format:"email" doc:"Email"`
}
type ForgotPasswordWithPhoneNumberInitRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
}
type ForgotPasswordInitRequest struct {
	Email       string `json:"email" required:"true" format:"email" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
}
type ForgotPasswordCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number"`
}
type ForgotPasswordNewPasswordRequest struct {
	Token       string `json:"token" required:"true" minLength:"3" doc:"Received token on step 2"`
	NewPassword string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
}
