package data

type ExamID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Exam id" example:"1"`
}

type ExamTypeID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Exam id" example:"1"`
}

type CreateExamRequest struct {
	SchoolID       int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	TeachingUnitID int64  `json:"teachingUnitID" required:"true" doc:"Teaching unit id" example:"1"`
	SubjectID      int64  `json:"subjectID" required:"true" doc:"Subject id" example:"1"`
	TypeID         int64  `json:"typeID" required:"true" doc:"Type id" example:"1"`
	Percentage     int    `json:"percentage" required:"true" doc:"Percentage" example:""`
	Description    string `json:"description" required:"false" doc:"Description" example:""`
}

type UpdateExamRequest struct {
	TypeID      int64  `json:"typeID" required:"true" doc:"Type id" example:"1"`
	Percentage  int    `json:"percentage" required:"true" doc:"Percentage" example:""`
	Description string `json:"description" required:"false" doc:"Description" example:""`
}

type ExamTypeRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}
