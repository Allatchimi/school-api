package data

import "api/common/types"

type RequestID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Request id" example:"1"`
}

type RequestRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" required:"true" doc:"Class Subject id" example:"1"`
	SequenceID     int64 `json:"sequenceID" required:"true" doc:"Sequence id" example:"1"`
	UnitID         int64 `json:"unitID" required:"true" doc:"Unit id" example:"1"`
	StudentID      int64 `json:"studentID" required:"true" doc:"Student id" example:"1"`

	Audience string `json:"audience" required:"true" enum:"teacher,director" doc:"Audience" example:"teacher"`
	Title    string `json:"title" required:"true" doc:"Title" example:""`
	Message  string `json:"message" required:"true" doc:"Message" example:""`

	Document1 string `json:"document1" required:"false" doc:"Document1" example:""`
	Document2 string `json:"document2" required:"false" doc:"Document2" example:""`
	Document3 string `json:"document3" required:"false" doc:"Document3" example:""`
	Document4 string `json:"document4" required:"false" doc:"Document4" example:""`
	Document5 string `json:"document5" required:"false" doc:"Document5" example:""`
}

type RequestUpdateRequest struct {
	Status         string `json:"status" required:"true" doc:"Status" example:""`
	StatusFeedback string `json:"statusFeedback" required:"true" doc:"Status feedback" example:""`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	SequenceID int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id" example:"1"`
	StudentID  int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id" example:"1"`
}
