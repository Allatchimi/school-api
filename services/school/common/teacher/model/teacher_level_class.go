package model

import (
	"api/common/types"
	"api/services/school/common/teacher/data"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelDomain "api/services/school/university/domain/model"
	modelLevel "api/services/school/university/level/model"
)

type TeacherLevelClass struct {
	types.BaseGormModel

	TeacherID int64    `gorm:"default:null"`
	Teacher   *Teacher `gorm:"default:null;foreignKey:TeacherID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	DomainID int64                         `gorm:"default:null"`
	Domain   *modelDomain.UniversityDomain `gorm:"default:null;foreignKey:DomainID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	LevelID  int64                         `gorm:"default:null"`
	Level    *modelLevel.UniversityLevel   `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                       `gorm:"default:null"`
	Class   *modelClass.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *TeacherLevelClass) ToTeacherLevelClassResponse() *data.TeacherLevelClassResponse {
	resp := &data.TeacherLevelClassResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Teacher = item.Teacher.ToTeacherResponse()
	resp.Year = item.Year.ToResponse()

	resp.Domain = item.Domain.ToResponse()
	resp.Level = item.Level.ToResponse()

	resp.Class = item.Class.ToResponse()
	return resp
}

func ToTeacherLevelClassResponseList(itemList []TeacherLevelClass) []data.TeacherLevelClassResponse {
	resp := make([]data.TeacherLevelClassResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToTeacherLevelClassResponse()
	}
	return resp
}
