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

type ParentResponseList struct {
	types.PaginatedResponse
	Data []ParentResponse `json:"data" required:"false" doc:"List of parents" example:"[]"`
}

type ParentStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentStudentResponse `json:"data" required:"false" doc:"List of parents levels/classes" example:"[]"`
}
