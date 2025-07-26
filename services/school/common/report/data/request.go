package data

import "api/common/types"

type ReportEntryID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Report id"`
}

type ReportGradeID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Report grade id"`
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

	Name           string  `json:"name" required:"true" doc:"Name"`
	Description    string  `json:"description" required:"false" doc:"Description"`
	Minimum        float64 `json:"minimum" required:"true" doc:"Minimum"`
	Maximum        float64 `json:"maximum" required:"true" doc:"Maximum"`
	IncludeMinimum bool    `json:"includeMinimum" required:"true" doc:"Include minimum"`
	IncludeMaximum bool    `json:"includeMaximum" required:"true" doc:"Include maximum"`
}

type ReportConfigRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	Notation                      float64 `json:"notation" required:"true" doc:"Notation"`
	MinimumRequiredValueToPromote float64 `json:"minimumRequiredValueToPromote" required:"true" doc:"Minimum required value to promote"`
}

type GetAllReportEntryRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	SequenceID int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	SemesterID int64 `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
	StudentID  int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}

type GetAllReportGradeRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllReportConfigRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
