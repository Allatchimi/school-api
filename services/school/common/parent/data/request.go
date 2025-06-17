package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
)

type ParentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent id" example:"1"`
}

type ParentStudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent level class id" example:"1"`
}

type ParentRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`

	AutoGenerateEmail bool                      `json:"autoGenerateEmail" required:"true" doc:"Auto generate email" example:"true"`
	Email             string                    `json:"email" required:"false" minLength:"3" max:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber       uint64                    `json:"phoneNumber" required:"false" doc:"Phone number" example:"237690909090"`
	Info              *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information" example:""`
}

type ParentStudentRequest struct {
	ParentID  int64 `json:"parentID" required:"true" doc:"Parent id" example:"1"`
	StudentID int64 `json:"studentID" required:"true" doc:"Student id" example:"1"`
}

type ParentAssignRequest struct {
	FirstName string `json:"firstName" required:"true" doc:"First name" example:"John"`
	LastName  string `json:"lastName" required:"true" doc:"Last name" example:"Doe"`
	IDCard    string `json:"IDCard" required:"true" doc:"ID Card" example:"1234567890"`
	Document1 string `json:"document1" required:"true" doc:"Document 1" example:"1234567890"`
	Document2 string `json:"document2" required:"true" doc:"Document 2" example:"1234567890"`
}

type ParentAssignStatusRequest struct {
	ParentAssignID int64  `json:"parentAssignID" required:"false" doc:"Parent assign id" example:"1"`
	Status         string `json:"status" required:"false" enum:"initiated,pending,approved,declined" doc:"Status" example:"initiated"`
	StatusFeedback string `json:"statusFeedback" required:"false" doc:"Status feedback" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllParentStudentRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	ParentID  int64 `json:"parentID" query:"parentID" required:"false" doc:"Parent id" example:"1"`
	StudentID int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id" example:"1"`
}
