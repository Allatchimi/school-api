package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataUser "api/services/user/user/data"
)

type DirectorResponse struct {
	types.BaseGormModelResponse
	DirectorPublicResponse
}

type DirectorPublicResponse struct {
	User   *dataUser.UserPublicResponse     `json:"user" doc:"User"`
	School *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
}

type DirectorResponseList struct {
	types.PaginatedResponse
	Data []DirectorResponse `json:"data" required:"false" doc:"List of directors" example:"[]"`
}
