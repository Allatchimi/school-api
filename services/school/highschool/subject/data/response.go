package data

import (
	"api/common/types"
)

type SubjectResponse struct {
	types.BaseGormModelResponse
	SchoolID     int64  `json:"schoolID" required:"false" doc:"School id"`
	ClassID      int64  `json:"classID" required:"false" doc:"Class id"`
	Name         string `json:"name" required:"false" doc:"Name"`
	Description  string `json:"description" required:"false" doc:"Description"`
	Coefficient  int    `json:"Coefficient" required:"false" doc:"Coefficient"`
	Program      string `json:"Program" required:"false" doc:"Program"`
	Requirements string `json:"Requirements" required:"false" doc:"Requirements"`
}

type SubjectTeacherResponse struct {
	types.BaseGormModelResponse
	SubjectID int64 `json:"subjectID" required:"false" doc:"Subject id"`
	UserID    int64 `json:"userID" required:"false" doc:"User id"`
}

type SubjectResponseList struct {
	types.PaginatedResponse
	Data []SubjectResponse `json:"data" required:"false" doc:"List of Subject" example:"[]"`
}
