package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
)

type ParentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent id"`
}

type ParentStudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent level class id"`
}

type ParentRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	AutoGenerateEmail bool                      `json:"autoGenerateEmail" required:"true" doc:"Auto generate email"`
	Email             string                    `json:"email" required:"false" format:"email" doc:"Email"`
	PhoneNumber       uint64                    `json:"phoneNumber" required:"false" minimum:"10000000" doc:"Phone number"`
	Info              *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type ParentStudentRequest struct {
	ParentID  int64 `json:"parentID" required:"true" doc:"Parent id"`
	StudentID int64 `json:"studentID" required:"true" doc:"Student id"`
}

type ParentAssignRequest struct {
	FirstName string `json:"firstName" required:"true" doc:"First name"`
	LastName  string `json:"lastName" required:"true" doc:"Last name"`
	IDCard    string `json:"IDCard" required:"true" doc:"ID Card"`
	Document1 string `json:"document1" required:"true" doc:"Document 1"`
	Document2 string `json:"document2" required:"true" doc:"Document 2"`
}

type ParentAssignStatusRequest struct {
	ParentAssignID int64  `json:"parentAssignID" required:"false" doc:"Parent assign id"`
	Status         string `json:"status" required:"false" enum:"initiated,pending,approved,declined" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"false" doc:"Status feedback"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllParentStudentRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	ParentID  int64 `json:"parentID" query:"parentID" required:"false" doc:"Parent id"`
	StudentID int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}
