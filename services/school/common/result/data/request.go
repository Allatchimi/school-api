package data

type ResultID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Result id" example:"1"`
}

type ResultRequest struct {
	StudentID int64 `json:"studentID" required:"true" doc:"Student id" example:"1"`
	ExamID    int64 `json:"resultID" required:"true" doc:"Result id" example:"1"`

	Value int `json:"value" required:"true" doc:"Value" example:"1"`
}

type GetAllRequest struct {
	SchoolID       int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	YearID         int64 `json:"yearID" query:"yearID" required:"false" doc:"Year id" example:"1"`
	TypeID         int64 `json:"typeID" query:"typeID" required:"false" doc:"Type id" example:"1"`
	UnitID         int64 `json:"unitID" query:"unitID" required:"false" doc:"Unit id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" query:"classSubjectID" required:"false" doc:"Class Subject id" example:"1"`
	SequenceID     int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id" example:"1"`
	StudentID      int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id" example:"1"`
	ExamID         int64 `json:"examID" query:"examID" required:"false" doc:"Exam id" example:"1"`
}
