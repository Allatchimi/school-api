package data

import (
	"api/common/types"
	"time"
)

type ExamID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Exam id"`
}

type ExamTypeID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Exam id"`
}

type ExamRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Class subject id"`
	SequenceID     int64 `json:"sequenceID" required:"false" doc:"Sequence id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`
	TypeID         int64 `json:"typeID" required:"true" doc:"Type id"`

	Status          string     `json:"status" required:"true" enum:"draft,online,results" doc:"Status"`
	Notation        float64    `json:"notation" required:"true" minimum:"1" doc:"Notation"`
	Percentage      int        `json:"percentage" required:"true" minimum:"1" maximum:"100" doc:"Percentage"`
	Description     string     `json:"description" required:"false" doc:"Description"`
	LocationType    string     `json:"locationType" required:"true" enum:"online,onsite" doc:"Location type"`
	LocationDetails string     `json:"locationDetails" required:"false" doc:"Location details"`
	Requirements    string     `json:"requirements" required:"false" doc:"Requirements"`
	AllowedItems    string     `json:"allowedItems" required:"false" doc:"Allowed items"`
	StartDate       *time.Time `json:"startDate" required:"true" doc:"Start date"`
	EndDate         *time.Time `json:"endDate" required:"true" doc:"End date"`
	IsRetry         bool       `json:"isRetry" required:"false" doc:"Is retry"`
}

type ExamTypeRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	Name        string `json:"name" required:"true" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentParentRequest
	SequenceID    int64    `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	QuarterID     int64    `json:"quarterID" query:"quarterID" required:"false" doc:"Quarter id"`
	SemesterID    int64    `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
	ClassID       int64    `json:"classID" query:"classID" required:"false" doc:"Class id"`
	LevelDomainID int64    `json:"levelDomainID" query:"levelDomainID" required:"false" doc:"Level domain id"`
	ExamID        int64    `json:"examID" query:"examID" required:"false" doc:"Exam id"`
	ExamTypeID    int64    `json:"examTypeID" query:"examTypeID" required:"false" doc:"Exam type id"`
	StatusList    []string `json:"statusList" query:"statusList" required:"false" doc:"Status list"`
}

type GetAllExamTypeRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
