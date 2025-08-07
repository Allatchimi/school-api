package data

// Login
type LoginDevice struct {
	Platform   string `json:"platform" required:"true" doc:"Platform name"`
	DeviceName string `json:"deviceName" required:"true" doc:"Device name"`
	App        string `json:"app" required:"true" doc:"Application used to login"`
}
type LoginWithEmailRequest struct {
	Email         string `json:"email" required:"true" format:"email" doc:"Email"`
	Password      string `json:"password" required:"true" doc:"Base64 encoded password"`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}
type LoginWithProviderRequest struct {
	Provider string `json:"provider" required:"true" doc:"Provider"`
	Token    string `json:"token" required:"true" doc:"Token" minLength:"3"`
}
type LoginRequest struct {
	Email         string `json:"email" required:"true" format:"email" doc:"Email"`
	Password      string `json:"password" required:"true" doc:"Base64 encoded password"`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}

// Register
type RegisterWithEmailRequest struct {
	Email    string `json:"email" required:"true" format:"email" doc:"Email"`
	Password string `json:"password" required:"true" doc:"Base64 encoded password"`
}
type RegisterRequest struct {
	Email    string `json:"email" required:"true" format:"email" doc:"Email"`
	Password string `json:"password" required:"true" doc:"Base64 encoded password"`
}

// Activate account
type ActivateAccountRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email"`
}

// Forgot password
type ForgotPasswordWithEmailInitRequest struct {
	Email string `json:"email" required:"true" format:"email" doc:"Email"`
}
type ForgotPasswordInitRequest struct {
	Email string `json:"email" required:"true" format:"email" doc:"Email"`
}
type ForgotPasswordCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on previous step"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email"`
}
type ForgotPasswordNewPasswordRequest struct {
	Token       string `json:"token" required:"true" minLength:"3" doc:"Received token on previous step"`
	NewPassword string `json:"password" required:"true" doc:"Base64 encoded password"`
}
