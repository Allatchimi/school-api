package data

import "time"

type UserID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"User id"`
}

type UserRequest struct {
	RoleID   int64 `json:"roleID" required:"true" doc:"Role id"`
	SchoolID int64 `json:"schoolID" required:"false" doc:"School id"`

	Email       string           `json:"email" required:"true" format:"email" doc:"Email"`
	PhoneNumber uint64           `json:"phoneNumber" required:"false" doc:"Phone number"`
	IsActivated bool             `json:"isActivated" required:"true" doc:"Is activated"`
	Status      string           `json:"status" required:"true" enum:"enabled,disabled" doc:"Status"`
	Info        *UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type UserInfoRequest struct {
	Username      string     `json:"username" required:"false" doc:"User name"`
	FirstName     string     `json:"firstName" required:"true" doc:"First name"`
	LastName      string     `json:"lastName" required:"true" doc:"Last name"`
	Gender        string     `json:"Gender" required:"true" enum:"male,female" doc:"Gender"`
	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Address       string     `json:"address" required:"false" doc:"Address"`
	Language      string     `json:"language" required:"false" doc:"Language code with 2 letter"`
	Image         string     `json:"image" required:"false" doc:"Thumbnail"`
}

type GetAllRequest struct {
	SchoolID int64  `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	RoleName string `json:"roleName" query:"roleName" required:"false" doc:"Role name"`
	RoleID   int64  `json:"roleID" query:"roleID" required:"false" doc:"Role id"`
	Feature  string `json:"feature" query:"feature" required:"false" doc:"Feature name"`
}
