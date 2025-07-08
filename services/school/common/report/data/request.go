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
}

type ReportGradeRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	MinimumResult        float64 `json:"minimumResult" required:"true" doc:"Minimum result"`
	MaximumResult        float64 `json:"maximumResult" required:"true" doc:"Maximum result"`
	IncludeMinimumResult bool    `json:"includeMinimumResult" required:"true" doc:"Include minimum result"`
	Correspondence       float64 `json:"correspondence" required:"true" doc:"Correspondence"`
	Grade                string  `json:"grade" required:"false" doc:"Grade"`
	GradeDescription     string  `json:"gradeDescription" required:"false" doc:"Grade description"`
}

type ReportConfigRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	Notation               float64 `json:"notation" required:"true" doc:"Notation"`
	NotationMinimumSuccess float64 `json:"notationMinimumSuccess" required:"true" doc:"Notation minimum success"`
}

type GetAllReportEntryRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	SequenceID int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	SemesterID int64 `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
	StudentID  int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}

type GetAllReportGradeRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllReportConfigRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
