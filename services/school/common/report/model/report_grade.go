package model

import (
	"api/common/types"
	"api/services/school/common/report/data"
	modelSchool "api/services/school/common/school/model"
)

type ReportGrade struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Type           string  `gorm:"default:null"`
	Name           string  `gorm:"default:null"`
	Description    string  `gorm:"default:null"`
	Minimum        float64 `gorm:"default:null"`
	Maximum        float64 `gorm:"default:null"`
	IncludeMinimum bool    `gorm:"default:null"`
	IncludeMaximum bool    `gorm:"default:null"`
}

func (item *ReportGrade) ToResponse() *data.ReportGradeResponse {
	if item == nil {
		return &data.ReportGradeResponse{}
	}
	resp := &data.ReportGradeResponse{}
	resp.Type = item.Type
	resp.Name = item.Name
	resp.Description = item.Description
	resp.Minimum = item.Minimum
	resp.Maximum = item.Maximum
	resp.IncludeMinimum = item.IncludeMinimum
	resp.IncludeMaximum = item.IncludeMaximum

	resp.School = item.School.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToReportGradeResponseList(itemList []ReportGrade) []data.ReportGradeResponse {
	resp := make([]data.ReportGradeResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
