package model

import (
	"api/common/types"
	"api/services/school/highschool/subject/data"
)

type Subject struct {
	types.BaseGormModel
	SchoolID int64 `gorm:"not null"`
	ClassID  int64 `gorm:"not null"`

	Name              string             `gorm:"not null"`
	Description       string             `gorm:"default:null"`
	Coefficient       int                `gorm:"default:1"`
	Program           string             `gorm:"default null"`
	Requirements      string             `gorm:"default null"`
	SubjectProfessors []SubjectProfessor `gorm:"default:null;foreignKey:SubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Subject) ToResponse() *data.SubjectResponse {
	resp := &data.SubjectResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.SchoolID = item.SchoolID
	resp.ClassID = item.ClassID
	resp.Name = item.Name
	resp.Description = item.Description
	resp.Coefficient = item.Coefficient
	resp.Program = item.Program
	resp.Requirements = item.Requirements
	resp.SubjectProfessors = ToSubjectProfessorResponseList(item.SubjectProfessors)
	return resp
}

func ToSubjectResponseList(itemList []Subject) []data.SubjectResponse {
	resp := make([]data.SubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
