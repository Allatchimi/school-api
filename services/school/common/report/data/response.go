package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataSequence "api/services/school/highschool/sequence/data"
	dataLevel "api/services/school/university/level/data"
	dataUnit "api/services/school/university/unit/data"
)

type ReportEntryResponse struct {
	types.BaseGormModelResponse
	School       *dataSchool.SchoolResponse      `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearResponse          `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Sequence     *dataSequence.SequenceResponse  `json:"sequence" required:"false" doc:"Sequence"`
	Unit         *dataUnit.UnitResponse          `json:"unit" required:"false" doc:"Unit"`
	Student      *dataStudent.StudentResponse    `json:"student" required:"false" doc:"Student"`

	Coefficient  int     `json:"coefficient" required:"false" doc:"Coefficient"`
	Credit       int     `json:"credit" required:"false" doc:"Credit"`
	Value        float64 `json:"value" required:"false" doc:"Value"`
	Notation     float64 `json:"notation" required:"false" doc:"Notation"`
	IsRetry      bool    `json:"isRetry" required:"false" doc:"Is retry"`
	RetryCount   int64   `json:"retryCount" required:"false" doc:"Retry count"`
	RetryDetails string  `json:"retryDetails" required:"false" doc:"Retry details"`
}

type ReportGradeResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`

	Name           string  `json:"name" required:"false" doc:"Name"`
	Description    string  `json:"description" required:"false" doc:"Description"`
	Minimum        float64 `json:"minimum" required:"false" doc:"Minimum"`
	Maximum        float64 `json:"maximum" required:"false" doc:"Maximum"`
	IncludeMinimum bool    `json:"includeMinimum" required:"false" doc:"Include minimum"`
	IncludeMaximum bool    `json:"includeMaximum" required:"false" doc:"Include maximum"`
}

type ReportConfigResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`

	Notation                      float64 `json:"notation" required:"false" doc:"Notation"`
	MinimumRequiredValueToPromote float64 `json:"minimumRequiredValueToPromote" required:"false" doc:"Minimum required value to promote"`
}

type ReportBoardResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolResponse     `json:"school" required:"false" doc:"School"`
	Year        *dataYear.YearResponse         `json:"year" required:"false" doc:"Year"`
	Class       *dataClass.ClassResponse       `json:"class" required:"false" doc:"Class"`
	LevelDomain *dataLevel.LevelDomainResponse `json:"levelDomain" required:"false" doc:"Level domain"`

	PeriodType                    string  `json:"periodType" required:"false" doc:"Period type"`
	PeriodName                    string  `json:"periodName" required:"false" doc:"Period name"`
	Status                        string  `json:"status" required:"false" doc:"Status"`
	Notation                      float64 `json:"notation" required:"false" doc:"Notation"`
	MinimumRequiredValueToPromote float64 `json:"minimumRequiredValueToPromote" required:"false" doc:"Minimum required value to promote"`
	GradeName                     string  `json:"gradeName" required:"false" doc:"Grade name"`
	GradeDescription              string  `json:"gradeDescription" required:"false" doc:"Grade description"`
}

type ReportEntryResponseList struct {
	types.PaginatedResponse
	Data []ReportEntryResponse `json:"data" required:"false" doc:"List of report entry"`
}

type ReportGradeResponseList struct {
	types.PaginatedResponse
	Data []ReportGradeResponse `json:"data" required:"false" doc:"List of report grade"`
}

type ReportConfigResponseList struct {
	types.PaginatedResponse
	Data []ReportConfigResponse `json:"data" required:"false" doc:"List of report config"`
}

type ReportBoardResponseList struct {
	types.PaginatedResponse
	Data []ReportBoardResponse `json:"data" required:"false" doc:"List of report board"`
}
