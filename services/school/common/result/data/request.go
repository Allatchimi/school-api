package data

import "api/common/types"

type ResultID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Result id"`
}

type ResultRequest struct {
	SchoolID  int64 `json:"schoolID" required:"true" doc:"School id"`
	StudentID int64 `json:"studentID" required:"true" doc:"Student id"`
	ExamID    int64 `json:"examID" required:"true" doc:"Exam id"`

	Value float64 `json:"value" required:"true" doc:"Value"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentRequest
	SequenceID int64 `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	SemesterID int64 `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
	ExamID     int64 `json:"examID" query:"examID" required:"false" doc:"Exam id"`
	TypeID     int64 `json:"typeID" query:"typeID" required:"false" doc:"Type id"`
}
