package model

import (
	"api/common/types"
	"api/services/school/common/teacher/data"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
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

	ClassSubjectID int64                                     `gorm:"default:null"`
	ClassSubject   *modelClassSubject.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *TeacherUnitSubject) ToTeacherUnitSubjectResponse() *data.TeacherUnitSubjectResponse {
	if item == nil {
		return nil
	}
	resp := &data.TeacherUnitSubjectResponse{}
	resp.Teacher = item.Teacher.ToTeacherPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *TeacherUnitSubject) ToTeacherUnitSubjectPublicResponse() *data.TeacherUnitSubjectPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.TeacherUnitSubjectPublicResponse{}
	resp.Teacher = item.Teacher.ToTeacherPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	return resp
}

func ToTeacherUnitSubjectResponseList(itemList []TeacherUnitSubject) []data.TeacherUnitSubjectResponse {
	resp := make([]data.TeacherUnitSubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToTeacherUnitSubjectResponse()
	}
	return resp
}
