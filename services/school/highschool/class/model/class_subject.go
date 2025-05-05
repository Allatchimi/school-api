package model

import (
	"api/common/types"
	"api/services/school/highschool/class/data"
)

type HighschoolClassSubject struct {
	types.BaseGormModel
	ClassID   int64 `gorm:"not null"`
	SubjectID int64 `gorm:"not null"`

	Coefficient  int    `gorm:"default:1"`
	Program      string `gorm:"default null"`
	Requirements string `gorm:"default null"`
}

func (item *HighschoolClassSubject) ToClassSubjectResponse() *data.ClassSubjectResponse {
	if item == nil {
		return nil
	}
	resp := &data.ClassSubjectResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.SubjectID = item.SubjectID
	resp.ClassID = item.ClassID

	resp.Coefficient = item.Coefficient
	resp.Program = item.Program
	resp.Requirements = item.Requirements
	return resp
}

func ToClassSubjectResponseList(itemList []HighschoolClassSubject) []data.ClassSubjectResponse {
	resp := make([]data.ClassSubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToClassSubjectResponse()
	}
	return resp
}
