package model

import (
	"api/common/types"
	"api/services/school/common/exam/data"
	schoolModel "api/services/school/common/school/model"
	subjectModel "api/services/school/highschool/subject/model"
	TUmodel "api/services/school/university/tu/model"
)

type Exam struct {
	types.BaseGormModel
	Percentage  int    `gorm:"default:null"`
	Description string `gorm:"default:null"`

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	TypeID int64     `gorm:"default:null"`
	Type   *ExamType `gorm:"default:null;foreignKey:TypeID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	TeachingUnitID int64                           `gorm:"default:null"`
	TeachingUnit   *TUmodel.UniversityTeachingUnit `gorm:"default:null;foreignKey:TeachingUnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SubjectID int64                           `gorm:"default:null"`
	Subject   *subjectModel.HighschoolSubject `gorm:"default:null;foreignKey:SubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Exam) ToResponse() *data.ExamResponse {
	if item == nil {
		return nil
	}
	resp := &data.ExamResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.SchoolID = item.SchoolID
	resp.TeachingUnit = item.TeachingUnit.ToResponse()
	resp.Subject = item.Subject.ToResponse()
	resp.Type = item.Type.ToResponse()
	resp.Percentage = item.Percentage
	resp.Description = item.Description
	return resp
}

func ToExamResponseList(itemList []Exam) []data.ExamResponse {
	resp := make([]data.ExamResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
