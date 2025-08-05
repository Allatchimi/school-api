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
	School       *dataSchool.SchoolResponse         `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearResponse             `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectResponse    `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Sequence     *dataSequence.SequenceResponse     `json:"sequence" required:"false" doc:"Sequence"`
	Unit         *dataUnit.UnitResponse             `json:"unit" required:"false" doc:"Unit"`
	Student      *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`

	CoefficientCredit int     `json:"coefficientCredit" required:"false" doc:"Coefficient/credit"`
	Score             float64 `json:"score" required:"false" doc:"Score"`
	Notation          float64 `json:"notation" required:"false" doc:"Notation"`
	GradeName         string  `json:"gradeName" required:"false" doc:"Grade name"`
	GradeDescription  string  `json:"gradeDescription" required:"false" doc:"Grade description"`
	IsRetry           bool    `json:"isRetry" required:"false" doc:"Is retry"`
	RetryCount        int64   `json:"retryCount" required:"false" doc:"Retry count"`
	RetryDetails      string  `json:"retryDetails" required:"false" doc:"Retry details"`
}

type ReportGradeResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`

	Type           string  `json:"type" required:"false" doc:"Type"`
	Name           string  `json:"name" required:"true" doc:"Name"`
	Description    string  `json:"description" required:"false" doc:"Description"`
	Minimum        float64 `json:"minimum" required:"false" doc:"Minimum"`
	Maximum        float64 `json:"maximum" required:"false" doc:"Maximum"`
	IncludeMinimum bool    `json:"includeMinimum" required:"false" doc:"Include minimum"`
	IncludeMaximum bool    `json:"includeMaximum" required:"false" doc:"Include maximum"`
}

type ReportCorrespondenceResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`

	Minimum        float64 `json:"minimum" required:"false" doc:"Minimum"`
	Maximum        float64 `json:"maximum" required:"false" doc:"Maximum"`
	IncludeMinimum bool    `json:"includeMinimum" required:"false" doc:"Include minimum"`
	IncludeMaximum bool    `json:"includeMaximum" required:"false" doc:"Include maximum"`
	NewScore       float64 `json:"newScore" required:"false" doc:"New score"`
}

type ReportConfigResponse struct {
	types.BaseGormModelResponse
	School            *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	ReportGradeToFail *ReportGradeResponse             `json:"reportGradeToFail" required:"false" doc:"Report grade to fail"`

	NotationAverage               float64 `json:"notationAverage" required:"false" doc:"Notation average"`
	NotationReport                float64 `json:"notationReport" required:"false" doc:"Notation report"`
	MinimumRequiredScoreToPromote float64 `json:"minimumRequiredScoreToPromote" required:"false" doc:"Minimum required score to promote"`
	OnlyFailedExams               bool    `json:"onlyFailedExams" required:"false" doc:"Only failed exams"`
}

type ReportAverageResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolResponse         `json:"school" required:"false" doc:"School"`
	Year        *dataYear.YearResponse             `json:"year" required:"false" doc:"Year"`
	Class       *dataClass.ClassResponse           `json:"class" required:"false" doc:"Class"`
	LevelDomain *dataLevel.LevelDomainResponse     `json:"levelDomain" required:"false" doc:"Level domain"`
	Student     *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`

	PeriodType                 string  `json:"periodType" required:"false" doc:"Period type"`
	PeriodName                 string  `json:"periodName" required:"false" doc:"Period name"`
	Score                      float64 `json:"score" required:"false" doc:"Score"`
	Notation                   float64 `json:"notation" required:"false" doc:"Notation"`
	GradeName                  string  `json:"gradeName" required:"false" doc:"Grade name"`
	GradeDescription           string  `json:"gradeDescription" required:"false" doc:"Grade description"`
	Rank                       int64   `json:"rank" required:"false" doc:"Rank"`
	IsSuccessful               bool    `json:"isSuccessful" required:"false" doc:"Is successful"`
	TotalCreditCoefficient     int     `json:"totalCreditCoefficient" required:"false" doc:"Total credit coefficient"`
	ValidatedCreditCoefficient int     `json:"validatedCreditCoefficient" required:"false" doc:"Validated credit coefficient"`
}

type ReportTableResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolResponse     `json:"school" required:"false" doc:"School"`
	Year        *dataYear.YearResponse         `json:"year" required:"false" doc:"Year"`
	Class       *dataClass.ClassResponse       `json:"class" required:"false" doc:"Class"`
	LevelDomain *dataLevel.LevelDomainResponse `json:"levelDomain" required:"false" doc:"Level domain"`

	PeriodType string `json:"periodType" required:"false" doc:"Period type"`
	PeriodName string `json:"periodName" required:"false" doc:"Period name"`
	Status     string `json:"status" required:"false" doc:"Status"`
}

type ReportEntryResponseList struct {
	types.PaginatedResponse
	Data []ReportEntryResponse `json:"data" required:"false" doc:"List of report entry"`
}

type ReportGradeResponseList struct {
	types.PaginatedResponse
	Data []ReportGradeResponse `json:"data" required:"false" doc:"List of report grade"`
}

type ReportCorrespondenceResponseList struct {
	types.PaginatedResponse
	Data []ReportCorrespondenceResponse `json:"data" required:"false" doc:"List of report correspondence"`
}

type ReportConfigResponseList struct {
	types.PaginatedResponse
	Data []ReportConfigResponse `json:"data" required:"false" doc:"List of report config"`
}

type ReportAverageResponseList struct {
	types.PaginatedResponse
	Data []ReportAverageResponse `json:"data" required:"false" doc:"List of report average"`
}

type ReportTableResponseList struct {
	types.PaginatedResponse
	Data []ReportTableResponse `json:"data" required:"false" doc:"List of report table"`
}
