package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataDomain "api/services/school/university/domain/data"
	dataLevel "api/services/school/university/level/data"
	dataSemester "api/services/school/university/semester/data"
	"time"
)

type TeachingUnitResponse struct {
	types.BaseGormModelResponse
	School   *dataSchool.SchoolResponse     `json:"school" required:"false" doc:"School"`
	Domain   *dataDomain.DomainResponse     `json:"domain" required:"false" doc:"Domain"`
	Level    *dataLevel.LevelResponse       `json:"level" required:"false" doc:"Level"`
	Semester *dataSemester.SemesterResponse `json:"semester" required:"false" doc:"Semester"`

	Name         string `json:"name" required:"false" doc:"Name"`
	Description  string `json:"description" required:"false" doc:"Description"`
	Credit       int    `json:"credit" required:"false" doc:"Credit"`
	Program      string `json:"program" required:"false" doc:"Program"`
	Requirements string `json:"requirements" required:"false" doc:"Requirements"`

	IsValid     bool       `json:"isValid" required:"false" doc:"Is valid"`
	InvalidDate *time.Time `json:"invalidDate" required:"false" doc:"Invalid date"`
}

type TeachingUnitResponseList struct {
	types.PaginatedResponse
	Data []TeachingUnitResponse `json:"data" required:"false" doc:"List of teaching unit" example:"[]"`
}
