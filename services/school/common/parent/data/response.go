package data

import (
	"api/common/types"
	dataStudent "api/services/school/common/student/data"
	dataUser "api/services/user/user/data"
)

type ParentResponse struct {
	types.BaseGormModelResponse
	User *dataUser.UserPublicResponse `json:"userID" required:"false" doc:"User id"`
}

type ParentPublicResponse struct {
	User *dataUser.UserPublicResponse `json:"userID" required:"false" doc:"User id"`
}

type ParentStudentResponse struct {
	types.BaseGormModelResponse
	Parent  *ParentPublicResponse              `json:"parent" required:"false" doc:"Parent"`
	Student *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`
}

type ParentStudentPublicResponse struct {
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

	Parent *ParentPublicResponse `json:"parent" required:"false" doc:"Parent"`
}

type ParentAssignPublicResponse struct {
	FirstName      string `json:"firstName" required:"false" doc:"First name"`
	LastName       string `json:"lastName" required:"false" doc:"Last name"`
	IDCard         string `json:"IDCard" required:"false" doc:"ID Card"`
	Document1      string `json:"document1" required:"false" doc:"Document 1"`
	Document2      string `json:"document2" required:"false" doc:"Document 2"`
	Status         string `json:"status" required:"false" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"false" doc:"Status feedback"`

	Parent *ParentPublicResponse `json:"parent" required:"false" doc:"Parent"`
}

type ParentAssignStudentResponse struct {
	types.BaseGormModelResponse
	ParentAssign *ParentAssignPublicResponse        `json:"parentAssign" required:"false" doc:"Parent assign request"`
	Student      *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`
}

type ParentAssignStudentPublicResponse struct {
	ParentAssign *ParentAssignPublicResponse        `json:"parentAssign" required:"false" doc:"Parent assign request"`
	Student      *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`
}

type ParentResponseList struct {
	types.PaginatedResponse
	Data []ParentResponse `json:"data" required:"false" doc:"List of parents" example:"[]"`
}

type ParentStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentStudentResponse `json:"data" required:"false" doc:"List of parents levels/classes" example:"[]"`
}

type ParentAssignResponseList struct {
	types.PaginatedResponse
	Data []ParentAssignResponse `json:"data" required:"false" doc:"List of parent assign requests" example:"[]"`
}

type ParentAssignStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentAssignStudentResponse `json:"data" required:"false" doc:"List of parent assign student" example:"[]"`
}
