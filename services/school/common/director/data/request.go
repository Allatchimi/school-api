package data

import dataUser "api/services/user/user/data"

type DirectorID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Director id"`
}

type DirectorRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	UID               string `json:"uid" required:"false" doc:"User UID"`
	AutoGenerateEmail bool   `json:"autoGenerateEmail" required:"false" doc:"Auto generate email"`
	Email             string `json:"email" required:"false" format:"email" doc:"Email"`
	PhoneNumber       uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`
	Status            string `json:"status" required:"true" enum:"enabled,disabled" doc:"Status"`

	Info *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
