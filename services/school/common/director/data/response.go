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
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`
}

type DirectorResponseList struct {
	types.PaginatedResponse
	Data []DirectorResponse `json:"data" required:"false" doc:"List of directors" example:"[]"`
}
