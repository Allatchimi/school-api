package data

import (
	"api/common/types"
	dataStudent "api/services/school/common/student/data"
	dataUser "api/services/user/user/data"
)

type ParentResponse struct {
	types.BaseGormModelResponse
	User *dataUser.UserResponse `json:"userID" required:"false" doc:"User id"`

	UID string `json:"uid" required:"false" doc:"Parent UID"`
}

type ParentStudentResponse struct {
	types.BaseGormModelResponse
	Parent  *ParentResponse              `json:"parent" required:"false" doc:"Parent"`
	Student *dataStudent.StudentResponse `json:"student" required:"false" doc:"Student"`
}

type ParentResponseList struct {
	types.PaginatedResponse
	Data []ParentResponse `json:"data" required:"false" doc:"List of parents" example:"[]"`
}

type ParentStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentStudentResponse `json:"data" required:"false" doc:"List of parents levels/classes" example:"[]"`
}
