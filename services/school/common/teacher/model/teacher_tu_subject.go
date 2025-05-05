package model

import (
	"api/common/types"
	"api/services/school/common/teacher/data"
	modelYear "api/services/school/common/year/model"
	modelSubject "api/services/school/highschool/subject/model"
	modelTeachingUnit "api/services/school/university/tu/model"
)

type TUSubject struct {
	types.BaseGormModel

	TeacherID int64    `gorm:"default:null"`
	Teacher   *Teacher `gorm:"default:null;foreignKey:TeacherID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	TeachingUnitID int64                                     `gorm:"default:null"`
	TeachingUnit   *modelTeachingUnit.UniversityTeachingUnit `gorm:"default:null;foreignKey:TeachingUnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SubjectID int64                           `gorm:"default:null"`
	Subject   *modelSubject.HighschoolSubject `gorm:"default:null;foreignKey:SubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *TUSubject) ToTUSubjectResponse() *data.TUSubjectResponse {
	if item == nil {
		return nil
	}
	resp := &data.TUSubjectResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Teacher = item.Teacher.ToTeacherResponse()
	resp.Year = item.Year.ToResponse()

	resp.TeachingUnit = item.TeachingUnit.ToResponse()
	resp.Subject = item.Subject.ToResponse()
	return resp
}

func ToTUSubjectResponseList(itemList []TUSubject) []data.TUSubjectResponse {
	resp := make([]data.TUSubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToTUSubjectResponse()
	}
	return resp
}
