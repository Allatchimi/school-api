package data

import "api/common/types"

type RequestID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Request id"`
}

type RequestRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"true" doc:"Class Subject id"`
	SequenceID     int64 `json:"sequenceID" required:"true" doc:"Sequence id"`
	UnitID         int64 `json:"unitID" required:"true" doc:"Unit id"`
	StudentID      int64 `json:"studentID" required:"true" doc:"Student id"`

	Audience string `json:"audience" required:"true" enum:"teacher,director" doc:"Audience"`
	Title    string `json:"title" required:"true" doc:"Title"`
	Message  string `json:"message" required:"true" doc:"Message"`

	Document1 string `json:"document1" required:"false" doc:"Document1"`
	Document2 string `json:"document2" required:"false" doc:"Document2"`
	Document3 string `json:"document3" required:"false" doc:"Document3"`
	Document4 string `json:"document4" required:"false" doc:"Document4"`
	Document5 string `json:"document5" required:"false" doc:"Document5"`
}

type RequestUpdateRequest struct {
	Status         string `json:"status" required:"true" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"true" doc:"Status feedback"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	SequenceID int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	StudentID  int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}
