package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	"time"
)

type YearResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`

	Name      string     `json:"name" required:"false" doc:"Name"`
	StartDate *time.Time `json:"startDate" required:"false" doc:"Academic year start date"`
	EndDate   *time.Time `json:"endDate" required:"false" doc:"Academic year end date"`
}

type YearResponseList struct {
	types.PaginatedResponse
	Data []YearResponse `json:"data" required:"false" doc:"List of academic year"`
}
