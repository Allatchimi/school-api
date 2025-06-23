package data

import "time"

type UpdateProfileEmailInitRequest struct {
	Email       string `json:"email" required:"true" format:"email" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
}

type UpdateProfileEmailCheckCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number"`
}

type UpdateProfileEmailNewEmailRequest struct {
	Email       string `json:"email" required:"true" format:"email" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
}

type UpdateProfilePhoneNumberInitRequest struct {
	Email       string `json:"email" required:"true" format:"email" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
}

type UpdateProfilePhoneNumberCheckCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1"`
	Code  int    `json:"code" required:"true" doc:"Received Code on your phone number"`
}

type UpdateProfilePhoneNumberNewPhoneNumberRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minimum:"10000000" doc:"Phone number"`
}

type UpdateProfilePasswordInitRequest struct {
}

type UpdateProfilePasswordCheckCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1"`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number"`
}
type UpdateProfilePasswordNewPasswordRequest struct {
	Token       string `json:"token" required:"true" minLength:"3" doc:"Received token on step 2"`
	NewPassword string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password"`
}

type UpdateProfileInfoRequest struct {
	Username  string `json:"username" required:"false" maxLength:"30" doc:"User name"`
	FirstName string `json:"firstName" required:"false" maxLength:"30" doc:"First name"`
	LastName  string `json:"lastName" required:"false" maxLength:"30" doc:"Last name"`

	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Address       string     `json:"address" required:"false" maxLength:"30" doc:"Address"`
	Language      string     `json:"language" required:"false" minLength:"2" maxLength:"2" doc:"Language code with 2 letter"`
	Image         string     `json:"image" required:"false" doc:"Thumbnail"`
}

type UpdateProfileMfaRequest struct {
	Method string `json:"method" required:"true" maxLength:"30" doc:"Method to update MFA"`
	Value  bool   `json:"value" required:"true" doc:"Method status"`
}
