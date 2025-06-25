package model

import (
	"api/common/types"
	"api/services/school/highschool/class/data"
	modelSubject "api/services/school/highschool/subject/model"
	"time"
)

type HighschoolClassSubject struct {
	types.BaseGormModel
	ClassID int64            `gorm:"default:null"`
	Class   *HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SubjectID int64                           `gorm:"default:null"`
	Subject   *modelSubject.HighschoolSubject `gorm:"default:null;foreignKey:SubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Coefficient  int        `gorm:"default:1"`
	Program      string     `gorm:"default null"`
	Requirements string     `gorm:"default null"`
	IsValid      bool       `gorm:"default:true"`
	InvalidDate  *time.Time `gorm:"default:null"`
}

func (item *HighschoolClassSubject) ToClassSubjectResponse() *data.ClassSubjectResponse {
	if item == nil {
		return nil
	}
	resp := &data.ClassSubjectResponse{}
	resp.Coefficient = item.Coefficient
	resp.Program = item.Program
	resp.Requirements = item.Requirements
	resp.IsValid = item.IsValid
	resp.InvalidDate = item.InvalidDate

	resp.Subject = item.Subject.ToPublicResponse()
	resp.Class = item.Class.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *HighschoolClassSubject) ToClassSubjectPublicResponse() *data.ClassSubjectPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ClassSubjectPublicResponse{}
	resp.Coefficient = item.Coefficient
	resp.Program = item.Program
	resp.Requirements = item.Requirements
	resp.IsValid = item.IsValid
	resp.InvalidDate = item.InvalidDate

	resp.Subject = item.Subject.ToPublicResponse()
	resp.Class = item.Class.ToPublicResponse()
	return resp
}

func ToClassSubjectResponseList(itemList []HighschoolClassSubject) []data.ClassSubjectResponse {
	resp := make([]data.ClassSubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToClassSubjectResponse()
	}
	return resp
}
