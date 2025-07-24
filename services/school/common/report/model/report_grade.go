package model

import (
	"api/common/types"
	"api/services/school/common/report/data"
	modelSchool "api/services/school/common/school/model"
)

type ReportGrade struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name                 string  `gorm:"default:null"`
	Description          string  `gorm:"default:null"`
	MinimumResult        float64 `gorm:"default:null"`
	MaximumResult        float64 `gorm:"default:null"`
	IncludeMinimumResult bool    `gorm:"default:null"`
	IncludeMaximumResult bool    `gorm:"default:null"`
	Correspondence       float64 `gorm:"default:null"`
}

func (item *ReportGrade) ToResponse() *data.ReportGradeResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportGradeResponse{}
	resp.Name = item.Name
	resp.Description = item.Description
	resp.MinimumResult = item.MinimumResult
	resp.MaximumResult = item.MaximumResult
	resp.IncludeMinimumResult = item.IncludeMinimumResult
	resp.IncludeMaximumResult = item.IncludeMaximumResult
	resp.Correspondence = item.Correspondence

	resp.School = item.School.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *ReportGrade) ToReportGradePublicResponse() *data.ReportGradePublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportGradePublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description
	resp.MinimumResult = item.MinimumResult
	resp.MaximumResult = item.MaximumResult
	resp.IncludeMinimumResult = item.IncludeMinimumResult
	resp.IncludeMaximumResult = item.IncludeMaximumResult
	resp.Correspondence = item.Correspondence

	resp.School = item.School.ToResponse()
	return resp
}

func ToReportGradeResponseList(itemList []ReportGrade) []data.ReportGradeResponse {
	resp := make([]data.ReportGradeResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
