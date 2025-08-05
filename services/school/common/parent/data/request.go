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

	AutoGenerateEmail bool                      `json:"autoGenerateEmail" required:"false" doc:"Auto generate email"`
	Email             string                    `json:"email" required:"false" format:"email" doc:"Email"`
	PhoneNumber       uint64                    `json:"phoneNumber" required:"false" doc:"Phone number"`
	Status            string                    `json:"status" required:"true" enum:"enabled,disabled" doc:"Status"`
	Info              *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type ParentStudentRequest struct {
	SchoolID  int64 `json:"schoolID" required:"true" doc:"School id"`
	ParentID  int64 `json:"parentID" required:"true" doc:"Parent id"`
	StudentID int64 `json:"studentID" required:"true" doc:"Student id"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllParentStudentRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	ParentID  int64 `json:"parentID" query:"parentID" required:"false" doc:"Parent id"`
	StudentID int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}
