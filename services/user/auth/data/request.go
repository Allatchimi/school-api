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
	SchoolID      int64  `json:"schoolID" required:"false" doc:"School ID"`
}
type LoginWithProviderRequest struct {
	Provider string `json:"provider" required:"true" doc:"Provider" minLength:"2" maxLength:"30"`
	Token    string `json:"token" required:"true" doc:"Token" minLength:"3"`
	SchoolID int64  `json:"schoolID" required:"false" doc:"School ID"`
}
type LoginRequest struct {
	Email         string `json:"email" required:"true" format:"email" doc:"Email"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
	SchoolID      int64  `json:"schoolID" required:"false" doc:"School ID"`
}

// Register
type RegisterWithEmailRequest struct {
	Email    string `json:"email" required:"true" format:"email" doc:"Email"`
	Password string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
	SchoolID int64  `json:"schoolID" required:"false" doc:"School ID"`
}
type RegisterRequest struct {
	Email    string `json:"email" required:"true" format:"email" doc:"Email"`
	Password string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
	SchoolID int64  `json:"schoolID" required:"false" doc:"School ID"`
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
	NewPassword string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
}
