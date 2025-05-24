package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataUser "api/services/user/user/data"
)

type QuizResponse struct {
	types.BaseGormModelResponse
	User   *dataUser.UserResponse     `json:"user" doc:"User"`
	School *dataSchool.SchoolResponse `json:"school" doc:"School"`
}

type QuizPublicResponse struct {
	User   *dataUser.UserPublicResponse     `json:"user" doc:"User"`
	School *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
}

type QuizResponseList struct {
	types.PaginatedResponse
	Data []QuizResponse `json:"data" required:"false" doc:"List of quiz" example:"[]"`
}
