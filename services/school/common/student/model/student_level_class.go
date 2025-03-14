package model

import (
	"api/common/types"
	"api/services/school/common/student/data"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
)

type StudentLevelClass struct {
	types.BaseGormModel

	StudentID int64    `gorm:"default:null"`
	Student   *Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelID int64                       `gorm:"default:null"`
	Level   *modelLevel.UniversityLevel `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                       `gorm:"default:null"`
	Class   *modelClass.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *StudentLevelClass) ToStudentLevelClassResponse() *data.StudentLevelClassResponse {
	resp := &data.StudentLevelClassResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Student = item.Student.ToStudentResponse()
	resp.Year = item.Year.ToResponse()
	resp.Level = item.Level.ToResponse()
	resp.Class = item.Class.ToResponse()
	return resp
}

func ToStudentLevelClassResponseList(itemList []StudentLevelClass) []data.StudentLevelClassResponse {
	resp := make([]data.StudentLevelClassResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToStudentLevelClassResponse()
	}
	return resp
}
