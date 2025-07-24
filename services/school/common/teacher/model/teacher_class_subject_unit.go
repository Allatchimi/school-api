package model

import (
	"api/common/types"
	"api/services/school/common/teacher/data"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
)

type TeacherClassSubjectUnit struct {
	types.BaseGormModel

	TeacherID int64    `gorm:"default:null"`
	Teacher   *Teacher `gorm:"default:null;foreignKey:TeacherID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                              `gorm:"default:null"`
	ClassSubject   *modelClass.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *TeacherClassSubjectUnit) ToResponse() *data.TeacherClassSubjectUnitResponse {
	if item == nil {
		return nil
	}
	resp := &data.TeacherClassSubjectUnitResponse{}
	resp.Teacher = item.Teacher.ToPublicResponse()
	resp.Year = item.Year.ToResponse()
	resp.ClassSubject = item.ClassSubject.ToResponse()
	resp.Unit = item.Unit.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToTeacherClassSubjectUnitResponseList(itemList []TeacherClassSubjectUnit) []data.TeacherClassSubjectUnitResponse {
	resp := make([]data.TeacherClassSubjectUnitResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
