package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	domainData "api/services/school/university/domain/data"
	"time"
)

type LevelResponse struct {
	types.BaseGormModelResponse
	LevelPublicResponse
}

type LevelPublicResponse struct {
	School      *schoolData.SchoolPublicResponse `json:"school" doc:"School"`
	Domain      *domainData.DomainPublicResponse `json:"domain" doc:"Domain"`
	Name        string                           `json:"name" required:"false" doc:"Level name"`
	Description string                           `json:"description" required:"false" doc:"Level description"`
}

type LevelDomainResponse struct {
	types.BaseGormModelResponse
	LevelDomainPublicResponse
}

type LevelDomainPublicResponse struct {
	Domain *domainData.DomainPublicResponse `json:"domain" required:"false" doc:"Domain"`
	Level  *LevelPublicResponse             `json:"level" required:"false" doc:"Level"`

	Fees         int64  `json:"fees" required:"false" doc:"Level domain fees"`
	Program      string `json:"program" required:"false" doc:"Program"`
	Requirements string `json:"requirements" required:"false" doc:"Requirements"`

	IsValid     bool       `json:"isValid" required:"false" doc:"Is valid"`
	InvalidDate *time.Time `json:"invalidDate" required:"false" doc:"Invalid date"`
}

type LevelResponseList struct {
	types.PaginatedResponse
	Data []LevelResponse `json:"data" required:"false" doc:"List of level" example:"[]"`
}

type LevelDomainResponseList struct {
	types.PaginatedResponse
	Data []LevelDomainResponse `json:"data" required:"false" doc:"List of level domain" example:"[]"`
}
