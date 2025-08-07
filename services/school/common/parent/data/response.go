package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	dataUser "api/services/user/user/data"
)

type ParentResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	User   *dataUser.UserResponse           `json:"user,omitempty" required:"false" doc:"User"`
}

type ParentPublicResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user,omitempty" required:"false" doc:"User"`
}

type ParentStudentResponse struct {
	types.BaseGormModelResponse
	School  *dataSchool.SchoolPublicResponse   `json:"school,omitempty" required:"false" doc:"School"`
	Parent  *ParentPublicResponse              `json:"parent,omitempty" required:"false" doc:"Parent"`
	Student *dataStudent.StudentPublicResponse `json:"student,omitempty" required:"false" doc:"Student"`
}

type ParentResponseList struct {
	types.PaginatedResponse
	Data []ParentResponse `json:"data" required:"false" doc:"List of parent"`
}

type ParentStudentResponseList struct {
	types.PaginatedResponse
	Data []ParentStudentResponse `json:"data" required:"false" doc:"List of parent"`
}
