package data

import "time"

type UserID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"User id"`
}

type UserRequest struct {
	RoleID int64 `json:"roleID" required:"true" doc:"Role id"`

	Email       string           `json:"email" required:"false" format:"email" doc:"Email"`
	PhoneNumber uint64           `json:"phoneNumber" required:"false" minimum:"10000000" doc:"Phone number"`
	IsActivated bool             `json:"isActivated" required:"true" doc:"Is activated"`
	Info        *UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type UserInfoRequest struct {
	Gender    string `json:"Gender" required:"true" enum:"male,female" doc:"Gender"`
	Username  string `json:"username" required:"false" maxLength:"30" doc:"User name"`
	FirstName string `json:"firstName" required:"true" maxLength:"30" doc:"First name"`
	LastName  string `json:"lastName" required:"true" maxLength:"30" doc:"Last name"`

	Birthday      *time.Time `json:"birthday" required:"true" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"true" doc:"Birth location"`
	Address       string     `json:"address" required:"false" maxLength:"30" doc:"Address"`
	Language      string     `json:"language" required:"false" min:"2" maxLength:"2" doc:"Language code with 2 letter"`
	Image         string     `json:"image" required:"false" doc:"Thumbnail"`
}

type GetAllRequest struct {
	RoleName string `json:"roleName" query:"roleName" required:"false" doc:"Role name"`
}
