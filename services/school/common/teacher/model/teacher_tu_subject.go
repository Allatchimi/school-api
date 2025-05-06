package model

import (
	"api/common/types"
	"api/services/school/common/teacher/data"
	modelYear "api/services/school/common/year/model"
	modelSubject "api/services/school/highschool/subject/model"
	modelTeachingUnit "api/services/school/university/tu/model"
)

type TeacherTeachingUnitSubject struct {
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

func (item *TeacherTeachingUnitSubject) ToTeacherTUSubjectResponse() *data.TeacherTUSubjectResponse {
	if item == nil {
		return nil
	}
	resp := &data.TeacherTUSubjectResponse{}
	resp.Teacher = item.Teacher.ToTeacherResponse()
	resp.Year = item.Year.ToResponse()
	resp.TeachingUnit = item.TeachingUnit.ToResponse()
	resp.Subject = item.Subject.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToTeacherTUSubjectResponseList(itemList []TeacherTeachingUnitSubject) []data.TeacherTUSubjectResponse {
	resp := make([]data.TeacherTUSubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToTeacherTUSubjectResponse()
	}
	return resp
}
