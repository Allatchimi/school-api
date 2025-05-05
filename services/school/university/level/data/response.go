package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	domainData "api/services/school/university/domain/data"
)

type LevelResponse struct {
	types.BaseGormModelResponse
	School      *schoolData.SchoolResponse `json:"school" doc:"School"`
	Domain      *domainData.DomainResponse `json:"domain" doc:"Domain"`
	Name        string                     `json:"name" required:"false" doc:"Level name"`
	Description string                     `json:"description" required:"false" doc:"Level description"`
}

type LevelResponseList struct {
	types.PaginatedResponse
	Data []LevelResponse `json:"data" required:"false" doc:"List of levels" example:"[]"`
}
