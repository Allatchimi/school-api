package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
	"time"
)

type ParentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent id"`
}

type ParentStudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent level class id"`
}

type ParentRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	AutoGenerateEmail bool   `json:"autoGenerateEmail" required:"false" doc:"Auto generate email"`
	Email             string `json:"email" required:"false" format:"email" doc:"Email"`
	PhoneNumber       uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`
	Status            string `json:"status" required:"true" enum:"enabled,disabled" doc:"Status"`

	Info *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type ParentStudentRequest struct {
	SchoolID  int64 `json:"schoolID" required:"true" doc:"School id"`
	ParentID  int64 `json:"parentID" required:"true" doc:"Parent id"`
	StudentID int64 `json:"studentID" required:"true" doc:"Student id"`
}

type ParentAssignRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	StudentListID string `json:"studentListID" true:"false" doc:"Student list id"`

	Message string `json:"message" required:"false" doc:"Message"`

	Gender        string     `json:"gender" required:"false" doc:"Gender"`
	FirstName     string     `json:"firstName" required:"true" doc:"First name"`
	LastName      string     `json:"lastName" required:"true" doc:"Last name"`
	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`

	Document1 string `json:"document1" required:"false" doc:"Document 1"`
	Document2 string `json:"document2" required:"false" doc:"Document 2"`
	Document3 string `json:"Document3" required:"false" doc:"Document 3"`
	Document4 string `json:"document4" required:"false" doc:"Document 4"`
	Document5 string `json:"document5" required:"false" doc:"Document 5"`
}

type ParentAssignStatusRequest struct {
	ParentAssignID int64  `json:"parentAssignID" required:"true" doc:"Parent assign id"`
	Status         string `json:"status" required:"true" enum:"initiated,pending,approved,rejected" doc:"Status"`
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

type GetAllParentAssignRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	ParentID int64 `json:"parentID" query:"parentID" required:"false" doc:"Parent id"`
}
