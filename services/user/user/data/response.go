package data

import (
	"time"

	"api/common/types"
	"api/services/user/role/data"
)

type UserResponse struct {
	types.BaseGormModelResponse
	Email       string `json:"email" required:"false" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`

	LoginMethod    string     `json:"loginMethod" required:"false" doc:"How the user should login ? with email, phone number or external provider?"`
	Provider       string     `json:"provider" required:"false" doc:"Provider name"`
	ProviderUserID string     `json:"providerUserID" required:"false" doc:"User id from the provider"`
	IsActivated    bool       `json:"isActivated" required:"false" doc:"Is user account activated ?"`
	ActivatedAt    *time.Time `json:"activatedAt" required:"false" doc:"Activation date time"`

	Role *data.RoleResponse `json:"role" required:"false" doc:"Role" `
	Info *UserInfoResponse  `json:"info" required:"false" doc:"Additional user info(e.g. address, first name, last name, ...)" `
	Mfa  *UserMfaResponse   `json:"mfa" required:"false" doc:"Multiple factor authenticator enabled by the user"`
}

type UserPublicResponse struct {
	Email string                  `json:"email" required:"false" doc:"Email"`
	Info  *UserInfoPublicResponse `json:"info" required:"false" doc:"Additional user info(e.g. address, first name, last name, ...)" `
}

type UserInfoResponse struct {
	Gender    string `json:"gender" required:"false" doc:"Gender"`
	Username  string `json:"username" required:"false" doc:"User name"`
	FirstName string `json:"firstName" required:"false" doc:"First name"`
	LastName  string `json:"lastName" required:"false" doc:"Last name or family name"`

	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Address       string     `json:"address" required:"false" doc:"Address"`
	Language      string     `json:"language" required:"false" doc:"Language"`
	Image         string     `json:"image" required:"false" doc:"Thumbnail"`
}

type UserInfoPublicResponse struct {
	Gender    string `json:"gender" required:"false" doc:"Gender"`
	Username  string `json:"username" required:"false" doc:"User name"`
	FirstName string `json:"firstName" required:"false" doc:"First name"`
	LastName  string `json:"lastName" required:"false" doc:"Last name or family name"`

	Image string `json:"image" required:"false" doc:"Thumbnail"`
}

type UserMfaResponse struct {
	Email         bool `json:"email" required:"false" doc:"Is 2FA enabled with email ?"`
	PhoneNumber   bool `json:"phoneNumber" required:"false" doc:"Is 2FA enabled with phone number ?"`
	Authenticator bool `json:"authenticator" required:"false" doc:"Is 2FA enabled with authenticator ?"`
}

type UserResponseList struct {
	types.PaginatedResponse
	Data []UserResponse `json:"data" required:"false" doc:"List of users" example:"[]"`
}
