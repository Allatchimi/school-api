package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	dataUser "api/services/user/user/data"
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
	Parent  *ParentPublicResponse              `json:"parent" required:"false" doc:"Parent"`
	Student *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`
}

type ParentAssignResponse struct {
	types.BaseGormModelResponse
	FirstName      string `json:"firstName" required:"false" doc:"First name"`
	LastName       string `json:"lastName" required:"false" doc:"Last name"`
	IDCard         string `json:"IDCard" required:"false" doc:"ID Card"`
	Document1      string `json:"document1" required:"false" doc:"Document 1"`
	Document2      string `json:"document2" required:"false" doc:"Document 2"`
	Status         string `json:"status" required:"false" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"false" doc:"Status feedback"`

	Parent               *ParentResponse               `json:"parent" required:"false" doc:"Parent"`
	ParentAssignStudents []ParentAssignStudentResponse `json:"parentAssignStudents" required:"false" doc:"Parent assign students"`
}

type ParentAssignStudentResponse struct {
	types.BaseGormModelResponse
	ParentAssign *ParentAssignResponse        `json:"parentAssign" required:"false" doc:"Parent assign request"`
	Student      *dataStudent.StudentResponse `json:"student" required:"false" doc:"Student"`
}

type ParentResponseList struct {
	types.PaginatedResponse
	Data []ParentResponse `json:"data" required:"false" doc:"List of parents"`
}

type ParentStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentStudentResponse `json:"data" required:"false" doc:"List of parents levels/classes"`
}

type ParentAssignResponseList struct {
	types.PaginatedResponse
	Data []ParentAssignResponse `json:"data" required:"false" doc:"List of parent assign requests"`
}

type ParentAssignStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentAssignStudentResponse `json:"data" required:"false" doc:"List of parent assign student"`
}
