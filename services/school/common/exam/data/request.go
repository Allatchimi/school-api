package data

import (
	"api/common/types"
	"time"
)

type ExamID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Exam id" example:"1"`
}

type ExamTypeID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Exam id" example:"1"`
}

type ExamRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`
	TypeID         int64 `json:"typeID" required:"true" doc:"Type id" example:"1"`
	UnitID         int64 `json:"unitID" required:"true" doc:"Unit id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" required:"true" doc:"Class subject id" example:"1"`
	SequenceID     int64 `json:"sequenceID" required:"true" doc:"Sequence id" example:"1"`

	Percentage      int        `json:"percentage" required:"true" min:"&" max:"100" doc:"Percentage" example:"100"`
	Description     string     `json:"description" required:"false" doc:"Description" example:""`
	LocationType    string     `json:"locationType" required:"true" enum:"online,onsite" doc:"Location type" example:"online"`
	LocationDetails string     `json:"locationDetails" required:"false" doc:"Location details" example:"AMPHI 520"`
	Requirements    string     `json:"requirements" required:"false" doc:"Requirements" example:"ID card"`
	AllowedItems    string     `json:"allowedItems" required:"false" doc:"Allowed items" example:"Pen, pencil, ruler, eraser"`
	StartDate       *time.Time `json:"startDate" required:"true" doc:"Start date" example:""`
	EndDate         *time.Time `json:"endDate" required:"true" doc:"End date" example:""`
}

type ExamTypeRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	TypeID     int64 `json:"typeID" query:"typeID" required:"false" doc:"Type id" example:"1"`
	SequenceID int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id" example:"1"`
}

type GetAllExamTypeRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
