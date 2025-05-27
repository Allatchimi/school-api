package data

import (
	"api/common/types"
)

type DocumentResponse struct {
	types.BaseGormModelResponse
	SchoolID       int64  `json:"schoolID" required:"false" doc:"School id"`
	YearID         int64  `json:"yearID" required:"false" doc:"Year id"`
	ClassSubjectID int64  `json:"classSubjectID" required:"false" doc:"Subject class id"`
	UnitID         int64  `json:"unitID" required:"false" doc:"Teaching unit id"`
	URL            string `json:"url" required:"false" doc:"URL"`
	Name           string `json:"name" required:"false" doc:"Name"`
	Description    string `json:"description" required:"false" doc:"Description"`
}

type DocumentResponseList struct {
	types.PaginatedResponse
	Data []DocumentResponse `json:"data" required:"false" doc:"List of documents" example:"[]"`
}
