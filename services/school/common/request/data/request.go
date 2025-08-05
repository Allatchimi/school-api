package data

import "api/common/types"

type RequestID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Request id"`
}

type RequestRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Class Subject id"`
	SequenceID     int64 `json:"sequenceID" required:"false" doc:"Sequence id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`

	Audience  string `json:"audience" required:"true" enum:"teacher,director" doc:"Audience"`
	Title     string `json:"title" required:"true" doc:"Title"`
	Message   string `json:"message" required:"true" doc:"Message"`
	Document1 string `json:"document1" required:"false" doc:"Document1"`
	Document2 string `json:"document2" required:"false" doc:"Document2"`
	Document3 string `json:"document3" required:"false" doc:"Document3"`
	Document4 string `json:"document4" required:"false" doc:"Document4"`
	Document5 string `json:"document5" required:"false" doc:"Document5"`
}

type RequestStatusRequest struct {
	Status         string `json:"status" required:"true" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"true" doc:"Status feedback"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentParentRequest
	SequenceID    int64  `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	ClassID       int64  `json:"classID" query:"classID" required:"false" doc:"Class id"`
	LevelDomainID int64  `json:"levelDomainID" query:"levelDomainID" required:"false" doc:"Level domain id"`
	Audience      string `json:"audience" query:"audience" required:"false" doc:"Audience"`
}
