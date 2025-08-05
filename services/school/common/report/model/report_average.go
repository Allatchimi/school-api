package model

import (
	"api/common/types"
	"api/services/school/common/report/data"
	modelSchool "api/services/school/common/school/model"
	modelStudent "api/services/school/common/student/model"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
)

type ReportAverage struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                       `gorm:"default:null"`
	Class   *modelClass.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelDomainID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentID int64                 `gorm:"default:null"`
	Student   *modelStudent.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	PeriodType                 string  `gorm:"default:null"`
	PeriodName                 string  `gorm:"default:null"`
	Score                      float64 `gorm:"default:null"`
	Notation                   float64 `gorm:"default:null"`
	GradeName                  string  `gorm:"default:null"`
	GradeDescription           string  `gorm:"default:null"`
	Rank                       int64   `gorm:"default:null"`
	IsSuccessful               bool    `gorm:"default:null"`
	TotalCreditCoefficient     int     `gorm:"default:null"`
	ValidatedCreditCoefficient int     `gorm:"default:null"`
}

func (item *ReportAverage) ToResponse() *data.ReportAverageResponse {
	if item == nil {
		return &data.ReportAverageResponse{}
	}
	resp := &data.ReportAverageResponse{}
	resp.PeriodType = item.PeriodType
	resp.PeriodName = item.PeriodName
	resp.Score = item.Score
	resp.Notation = item.Notation
	resp.GradeName = item.GradeName
	resp.GradeDescription = item.GradeDescription
	resp.Rank = item.Rank
	resp.IsSuccessful = item.IsSuccessful
	resp.TotalCreditCoefficient = item.TotalCreditCoefficient
	resp.ValidatedCreditCoefficient = item.ValidatedCreditCoefficient

	resp.School = item.School.ToResponse()
	resp.Year = item.Year.ToResponse()
	resp.Class = item.Class.ToResponse()
	resp.LevelDomain = item.LevelDomain.ToResponse()
	resp.Student = item.Student.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToReportAverageResponseList(itemList []ReportAverage) []data.ReportAverageResponse {
	resp := make([]data.ReportAverageResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
