package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataUser "api/services/user/user/data"
)

type DirectorResponse struct {
	types.BaseGormModelResponse
	User   *dataUser.UserResponse     `json:"user" doc:"User"`
	School *dataSchool.SchoolResponse `json:"school" doc:"School"`
}

type DirectorPublicResponse struct {
	User   *dataUser.UserPublicResponse     `json:"user" doc:"User"`
	School *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
}

type DirectorResponseList struct {
	types.PaginatedResponse
	Data []DirectorResponse `json:"data" required:"false" doc:"List of academic directors" example:"[]"`
}
