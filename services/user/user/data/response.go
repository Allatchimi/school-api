package data

import (
	"time"

	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataRole "api/services/user/role/data"
)

type UserResponse struct {
	types.BaseGormModelResponse
	Email       string `json:"email" required:"false" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`
	Status      string `json:"status" required:"false" doc:"Status"`

	LoginMethod    string     `json:"loginMethod" required:"false" doc:"How the user should login ? with email or external provider?"`
	Provider       string     `json:"provider" required:"false" doc:"Provider name"`
	ProviderUserID string     `json:"providerUserID" required:"false" doc:"User id from the provider"`
	IsActivated    bool       `json:"isActivated" required:"false" doc:"Is user account activated ?"`
	ActivatedAt    *time.Time `json:"activatedAt" required:"false" doc:"Activation date time"`

	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Role   *dataRole.RoleResponse           `json:"role" required:"false" doc:"Role" `
	Info   *UserInfoResponse                `json:"info" required:"false" doc:"Additional user info(e.g. address, first name, last name, ...)" `
	Config *UserConfigResponse              `json:"config" required:"false" doc:"Multiple factor authenticator enabled by the user"`
}

type UserPublicResponse struct {
	types.BaseGormModelResponse
	Email  string `json:"email" required:"false" doc:"Email"`
	Status string `json:"status" required:"false" doc:"Status"`

	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Role   *dataRole.RoleResponse           `json:"role" required:"false" doc:"Role"`
	Info   *UserInfoPublicResponse          `json:"info" required:"false" doc:"Additional user info(e.g. address, first name, last name, ...)" `
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

type UserConfigResponse struct {
	WhatsappPhoneNumber int64 `json:"whatsappPhoneNumber" required:"false" doc:"Whatsapp phone number"`
	TelegramChatID      int64 `json:"telegramChatID" required:"false" doc:"Telegram chat id"`

	AllowNotifications bool `json:"allowNotifications" required:"false" doc:"Allow notifications"`
	MfaEmail           bool `json:"email" required:"false" doc:"Is 2FA enabled with email ?"`
	MfaAuthenticator   bool `json:"authenticator" required:"false" doc:"Is 2FA enabled with authenticator ?"`

	WebPushSubscriptionEndpoint  string `json:"webPushSubscriptionEndpoint" required:"false" doc:"Web push subscription endpoint"`
	WebPushSubscriptionKeyP256dh string `json:"webPushSubscriptionKeyP256dh" required:"false" doc:"Web push subscription key P256dh"`
	WebPushSubscriptionKeyAuth   string `json:"webPushSubscriptionKeyAuth" required:"false" doc:"Web push subscription key auth"`
}

type UserResponseList struct {
	types.PaginatedResponse
	Data []UserResponse `json:"data" required:"false" doc:"List of users" example:"[]"`
}
