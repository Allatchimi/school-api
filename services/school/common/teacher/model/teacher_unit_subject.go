package model

import (
	"api/common/types"
	"api/services/school/common/teacher/data"
	modelYear "api/services/school/common/year/model"
	modelSubject "api/services/school/highschool/subject/model"
	modelUnit "api/services/school/university/unit/model"
)

type TeacherUnitSubject struct {
	types.BaseGormModel

	TeacherID int64    `gorm:"default:null"`
	Teacher   *Teacher `gorm:"default:null;foreignKey:TeacherID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SubjectID int64                           `gorm:"default:null"`
	Subject   *modelSubject.HighschoolSubject `gorm:"default:null;foreignKey:SubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *TeacherUnitSubject) ToTeacherUnitSubjectResponse() *data.TeacherUnitSubjectResponse {
	if item == nil {
		return nil
	}
	resp := &data.TeacherUnitSubjectResponse{}
	resp.Teacher = item.Teacher.ToTeacherResponse()
	resp.Year = item.Year.ToResponse()
	resp.Unit = item.Unit.ToResponse()
	resp.Subject = item.Subject.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToTeacherUnitSubjectResponseList(itemList []TeacherUnitSubject) []data.TeacherUnitSubjectResponse {
	resp := make([]data.TeacherUnitSubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToTeacherUnitSubjectResponse()
	}
	return resp
}
