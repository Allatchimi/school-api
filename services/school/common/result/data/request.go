package data

type ResultID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Result id"`
}

type ResultRequest struct {
	StudentID int64 `json:"studentID" required:"true" doc:"Student id"`
	ExamID    int64 `json:"resultID" required:"true" doc:"Result id"`

	Value float64 `json:"value" required:"true" doc:"Value"`
}

type GetAllRequest struct {
	SchoolID       int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	YearID         int64 `json:"yearID" query:"yearID" required:"false" doc:"Year id"`
	TypeID         int64 `json:"typeID" query:"typeID" required:"false" doc:"Type id"`
	UnitID         int64 `json:"unitID" query:"unitID" required:"false" doc:"Unit id"`
	ClassSubjectID int64 `json:"classSubjectID" query:"classSubjectID" required:"false" doc:"Class Subject id"`
	SequenceID     int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	StudentID      int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
	ExamID         int64 `json:"examID" query:"examID" required:"false" doc:"Exam id"`
}
