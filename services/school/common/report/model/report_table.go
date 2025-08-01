package model

import (
	"api/common/types"
	"api/services/school/common/report/data"
	modelSchool "api/services/school/common/school/model"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
)

type ReportTable struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                       `gorm:"default:null"`
	Class   *modelClass.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelDomainID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	PeriodType                    string  `gorm:"default:null"`
	PeriodName                    string  `gorm:"default:null"`
	Status                        string  `gorm:"default:null"`
	Notation                      float64 `gorm:"default:null"`
	MinimumRequiredScoreToPromote float64 `gorm:"default:null"`
	GradeName                     string  `gorm:"default:null"`
	GradeDescription              string  `gorm:"default:null"`
}

func (item *ReportTable) ToResponse() *data.ReportTableResponse {
	if item == nil {
		return &data.ReportTableResponse{}
	}
	resp := &data.ReportTableResponse{}
	resp.PeriodType = item.PeriodType
	resp.PeriodName = item.PeriodName
	resp.Status = item.Status
	resp.Notation = item.Notation
	resp.MinimumRequiredScoreToPromote = item.MinimumRequiredScoreToPromote
	resp.GradeName = item.GradeName
	resp.GradeDescription = item.GradeDescription

	resp.School = item.School.ToResponse()
	resp.Year = item.Year.ToResponse()
	resp.Class = item.Class.ToResponse()
	resp.LevelDomain = item.LevelDomain.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToReportTableResponseList(itemList []ReportTable) []data.ReportTableResponse {
	resp := make([]data.ReportTableResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
