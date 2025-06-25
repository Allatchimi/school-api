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

	MinimumResult        float64 `gorm:"default:null"`
	MaximumResult        float64 `gorm:"default:null"`
	IncludeMinimumResult bool    `gorm:"default:null"`
	IncludeMaximumResult bool    `gorm:"default:null"`
	Correspondence       float64 `gorm:"default:null"`
	Grade                string  `gorm:"default:null"`
	GradeDescription     string  `gorm:"default:null"`
}

func (item *ReportGrade) ToReportGradeResponse() *data.ReportGradeResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportGradeResponse{}
	resp.MinimumResult = item.MinimumResult
	resp.MaximumResult = item.MaximumResult
	resp.IncludeMinimumResult = item.IncludeMinimumResult
	resp.IncludeMaximumResult = item.IncludeMaximumResult
	resp.Correspondence = item.Correspondence
	resp.Grade = item.Grade
	resp.GradeDescription = item.GradeDescription

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
	resp.MinimumResult = item.MinimumResult
	resp.MaximumResult = item.MaximumResult
	resp.IncludeMinimumResult = item.IncludeMinimumResult
	resp.IncludeMaximumResult = item.IncludeMaximumResult
	resp.Correspondence = item.Correspondence
	resp.Grade = item.Grade
	resp.GradeDescription = item.GradeDescription

	resp.School = item.School.ToResponse()
	return resp
}

func ToReportGradeResponseList(itemList []ReportGrade) []data.ReportGradeResponse {
	resp := make([]data.ReportGradeResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToReportGradeResponse()
	}
	return resp
}
