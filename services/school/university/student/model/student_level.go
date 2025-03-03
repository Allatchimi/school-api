package model

import (
	"api/common/types"
	modelYear "api/services/school/common/year/model"
	modelLevel "api/services/school/university/level/model"
	"api/services/school/university/student/data"
)

type StudentLevel struct {
	types.BaseGormModel

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelID int64                       `gorm:"default:null"`
	Level   *modelLevel.UniversityLevel `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *StudentLevel) ToStudentLevelResponse() *data.StudentLevelResponse {
	resp := &data.StudentLevelResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Year = item.Year.ToResponse()
	resp.Level = item.Level.ToResponse()
	return resp
}

func ToStudentLevelResponseList(itemList []StudentLevel) []data.StudentLevelResponse {
	resp := make([]data.StudentLevelResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToStudentLevelResponse()
	}
	return resp
}
