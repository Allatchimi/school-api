package data

import dataUser "api/services/user/user/data"

type DirectorID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Director id" example:"1"`
}

type DirectorRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`

	UID               string                    `json:"uid" required:"false" doc:"User UID" example:"1"`
	AutoGenerateEmail bool                      `json:"autoGenerateEmail" required:"true" doc:"Auto generate email" example:"true"`
	Email             string                    `json:"email" required:"false" minLength:"3" max:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber       uint64                    `json:"phoneNumber" required:"false" doc:"Phone number" example:"237690909090"`
	Info              *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
