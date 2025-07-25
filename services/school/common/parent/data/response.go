package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	dataUser "api/services/user/user/data"
	"time"
)

type ParentResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserResponse           `json:"user" required:"false" doc:"User"`
}

type ParentPublicResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`
}

type ParentStudentResponse struct {
	types.BaseGormModelResponse
	School  *dataSchool.SchoolPublicResponse   `json:"school" required:"false" doc:"School"`
	Parent  *ParentPublicResponse              `json:"parent" required:"false" doc:"Parent"`
	Student *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`
}

type ParentAssignResponse struct {
	types.BaseGormModelResponse
	StudentListID string `json:"studentListID" required:"false" doc:"Student list id"`

	Status         string `json:"status" required:"false" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"false" doc:"Status feedback"`

	Message string `json:"message" required:"false" doc:"Message"`

	Gender        string     `json:"gender" required:"false" doc:"Gender"`
	FirstName     string     `json:"firstName" required:"false" doc:"First name"`
	LastName      string     `json:"lastName" required:"false" doc:"Last name"`
	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`

	Document1 string `json:"document1" required:"false" doc:"Document 1"`
	Document2 string `json:"document2" required:"false" doc:"Document 2"`
	Document3 string `json:"Document3" required:"false" doc:"Document 3"`
	Document4 string `json:"document4" required:"false" doc:"Document 4"`
	Document5 string `json:"document5" required:"false" doc:"Document 5"`

	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`
}

type ParentResponseList struct {
	types.PaginatedResponse
	Data []ParentResponse `json:"data" required:"false" doc:"List of parent"`
}

type ParentStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentStudentResponse `json:"data" required:"false" doc:"List of parent"`
}

type ParentAssignResponseList struct {
	types.PaginatedResponse
	Data []ParentAssignResponse `json:"data" required:"false" doc:"List of parent assign"`
}
