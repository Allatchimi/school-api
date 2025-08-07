package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
)

type ContactResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`

	Subject string `json:"subject" required:"true" doc:"Subject"`
	Email   string `json:"email" required:"true" doc:"Email"`
	Message string `json:"message" required:"true" doc:"Message"`
}

type ContactResponseList struct {
	types.PaginatedResponse
	Data []ContactResponse `json:"data" required:"false" doc:"List of contact"`
}
