package data

import "api/common/types"

type ReportEntryID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Report id"`
}

type ReportGradeID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Report grade id"`
}

type ReportCorrespondenceID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Report correspondence id"`
}

type ReportConfigID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Report config id"`
}

type ReportEntryRequest struct {
	SchoolID      int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID        int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassID       int64 `json:"classID" required:"false" doc:"Class id"`
	LevelDomainID int64 `json:"levelDomainID" required:"false" doc:"Level domain id"`

	PeriodType string `json:"periodType" required:"true" enum:"final,semester,quarter,sequence" doc:"Period type"`
	QuarterID  int64  `json:"quarterID" required:"false" doc:"Quarter id"`
	SequenceID int64  `json:"sequenceID" required:"false" doc:"Sequence id"`
	SemesterID int64  `json:"semesterID" required:"false" doc:"Semester id"`
}

type ReportGradeRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	Type           string  `json:"type" required:"true" enum:"average,report" doc:"Type"`
	Name           string  `json:"name" required:"true" doc:"Name"`
	Description    string  `json:"description" required:"false" doc:"Description"`
	Minimum        float64 `json:"minimum" required:"true" doc:"Minimum"`
	Maximum        float64 `json:"maximum" required:"true" doc:"Maximum"`
	IncludeMinimum bool    `json:"includeMinimum" required:"false" doc:"Include minimum"`
	IncludeMaximum bool    `json:"includeMaximum" required:"false" doc:"Include maximum"`
}

type ReportCorrespondenceRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	Minimum        float64 `json:"minimum" required:"true" doc:"Minimum"`
	Maximum        float64 `json:"maximum" required:"true" doc:"Maximum"`
	IncludeMinimum bool    `json:"includeMinimum" required:"false" doc:"Include minimum"`
	IncludeMaximum bool    `json:"includeMaximum" required:"false" doc:"Include maximum"`
	NewScore       float64 `json:"newScore" required:"true" doc:"New score"`
}

type ReportConfigRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	NotationAverage               float64 `json:"notationAverage" required:"true" doc:"Notation average"`
	NotationReport                float64 `json:"notationReport" required:"true" doc:"Notation report"`
	MinimumRequiredScoreToPromote float64 `json:"minimumRequiredScoreToPromote" required:"true" doc:"Minimum required score to promote"`
}

type GetAllReportEntryRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	SequenceID int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	SemesterID int64 `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
	StudentID  int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}

type GetAllReportGradeRequest struct {
	SchoolID int64  `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	Type     string `json:"type" query:"type" required:"false" doc:"Type"`
}

type GetAllReportCorrespondenceRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllReportConfigRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllReportTableRequest struct {
	types.FilterSchoolYearClassLevelDomainRequest
	PeriodType string `json:"periodType" query:"periodType" required:"false" doc:"Period type"`
	PeriodName string `json:"periodName" query:"periodName" required:"false" doc:"Period name"`
}
