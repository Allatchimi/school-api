package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataUser "api/services/user/user/data"
)

type ManagerResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	User   *dataUser.UserResponse           `json:"user,omitempty" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Manager UID"`
}

type ManagerPublicResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user,omitempty" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Manager UID"`
}

type ManagerResponseList struct {
	types.PaginatedResponse
	Data []ManagerResponse `json:"data" required:"false" doc:"List of manager" example:"[]"`
}
