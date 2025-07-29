package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataDomain "api/services/school/university/domain/data"
	dataLevel "api/services/school/university/level/data"
	dataSemester "api/services/school/university/semester/data"
	"time"
)

type UnitResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Domain      *dataDomain.DomainResponse       `json:"domain" required:"false" doc:"Domain"`
	LevelDomain *dataLevel.LevelDomainResponse   `json:"levelDomain" required:"false" doc:"Level domain"`
	Semester    *dataSemester.SemesterResponse   `json:"semester" required:"false" doc:"Semester"`

	Name         string     `json:"name" required:"false" doc:"Name"`
	Description  string     `json:"description" required:"false" doc:"Description"`
	Credit       int        `json:"credit" required:"false" doc:"Credit"`
	Program      string     `json:"program" required:"false" doc:"Program"`
	Requirements string     `json:"requirements" required:"false" doc:"Requirements"`
	IsValid      bool       `json:"isValid" required:"false" doc:"Is valid"`
	InvalidDate  *time.Time `json:"invalidDate" required:"false" doc:"Invalid date"`
}

type UnitResponseList struct {
	types.PaginatedResponse
	Data []UnitResponse `json:"data" required:"false" doc:"List of unit"`
}
