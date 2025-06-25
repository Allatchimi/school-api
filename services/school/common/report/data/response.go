package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataSequence "api/services/school/highschool/sequence/data"
	dataUnit "api/services/school/university/unit/data"
)

type ReportResponse struct {
	types.BaseGormModelResponse
	ReportPublicResponse
}

type ReportPublicResponse struct {
	School       *dataSchool.SchoolResponse      `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearResponse          `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Sequence     *dataSequence.SequenceResponse  `json:"sequence" required:"false" doc:"Sequence"`
	Unit         *dataUnit.UnitResponse          `json:"unit" required:"false" doc:"Unit"`
	Student      *dataStudent.StudentResponse    `json:"student" required:"false" doc:"Student"`

	Coefficient int     `json:"coefficient" required:"false" doc:"Coefficient"`
	Credit      int     `json:"credit" required:"false" doc:"Credit"`
	Value       float64 `json:"value" required:"false" doc:"Value"`
	Notation    float64 `json:"notation" required:"false" doc:"Notation"`
}

type ReportGradeResponse struct {
	types.BaseGormModelResponse
	ReportGradePublicResponse
}

type ReportGradePublicResponse struct {
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`

	MinimumResult        float64 `json:"minimumResult" required:"false" doc:"Minimum result"`
	MaximumResult        float64 `json:"maximumResult" required:"false" doc:"Maximum result"`
	IncludeMinimumResult bool    `json:"includeMinimumResult" required:"false" doc:"Include minimum result"`
	IncludeMaximumResult bool    `json:"includeMaximumResult" required:"false" doc:"Include maximum result"`
	Correspondence       float64 `json:"correspondence" required:"false" doc:"Correspondence"`
	Grade                string  `json:"grade" required:"false" doc:"Grade"`
	GradeDescription     string  `json:"gradeDescription" required:"false" doc:"Grade description"`
}

type ReportConfigResponse struct {
	types.BaseGormModelResponse
	ReportConfigPublicResponse
}

type ReportConfigPublicResponse struct {
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`

	Notation               float64 `json:"notation" required:"false" doc:"Notation"`
	NotationMinimumSuccess float64 `json:"notationMinimumSuccess" required:"false" doc:"Notation minimum success"`
}

type ReportResponseList struct {
	types.PaginatedResponse
	Data []ReportResponse `json:"data" required:"false" doc:"List of reports" example:"[]"`
}

type ReportGradeResponseList struct {
	types.PaginatedResponse
	Data []ReportGradeResponse `json:"data" required:"false" doc:"List of report grades" example:"[]"`
}

type ReportConfigResponseList struct {
	types.PaginatedResponse
	Data []ReportConfigResponse `json:"data" required:"false" doc:"List of report configs" example:"[]"`
}
