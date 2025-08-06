package data

import "api/common/types"

type ResultID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Result id"`
}

type ResultTableID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Result table id"`
}

type ResultRequest struct {
	SchoolID  int64 `json:"schoolID" required:"true" doc:"School id"`
	ExamID    int64 `json:"examID" required:"true" doc:"Exam id"`
	StudentID int64 `json:"studentID" required:"true" doc:"Student id"`

	Score float64 `json:"score" required:"true" doc:"Score"`
}

type ResultTableRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`
	ExamID   int64 `json:"examID" required:"true" doc:"Exam id"`

	Status string `json:"status" required:"true" enum:"draft,published" doc:"Status"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentParentRequest
	SequenceID      int64    `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	QuarterID       int64    `json:"quarterID" query:"quarterID" required:"false" doc:"Quarter id"`
	SemesterID      int64    `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
	ClassID         int64    `json:"classID" query:"classID" required:"false" doc:"Class id"`
	LevelDomainID   int64    `json:"levelDomainID" query:"levelDomainID" required:"false" doc:"Level domain id"`
	ExamID          int64    `json:"examID" query:"examID" required:"false" doc:"Exam id"`
	ExamTypeID      int64    `json:"examTypeID" query:"examTypeID" required:"false" doc:"Exam type id"`
	TableStatusList []string `json:"tableStatusList" query:"tableStatusList" required:"false" doc:"Table status list"`
}

type GetAllResultTableRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentParentRequest
	SequenceID     int64    `json:"sequenceID" query:"sequenceID" required:"false" doc:"Sequence id"`
	SemesterID     int64    `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
	ExamID         int64    `json:"examID" query:"examID" required:"false" doc:"Exam id"`
	ExamTypeID     int64    `json:"examTypeID" query:"examTypeID" required:"false" doc:"Exam type id"`
	ExamStatusList []string `json:"examStatusList" query:"examStatusList" required:"false" doc:"Exam status list"`
}
