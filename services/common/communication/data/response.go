package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataRole "api/services/user/role/data"
)

type CommunicationResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Role   *dataRole.RoleResponse           `json:"role" required:"false" doc:"Role"`

	Subject  string `json:"subject" required:"true" doc:"Subject"`
	Message  string `json:"message" required:"true" doc:"Message"`
	Audience string `json:"audience" required:"true" doc:"Audience"`
}

type CommunicationResponseList struct {
	types.PaginatedResponse
	Data []CommunicationResponse `json:"data" required:"false" doc:"List of communications" example:"[]"`
}
