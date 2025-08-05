package model

import (
	"api/common/types"
	"api/services/school/common/report/data"
	modelSchool "api/services/school/common/school/model"
)

type ReportConfig struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ReportGradeToFailID int64        `gorm:"default:null"`
	ReportGradeToFail   *ReportGrade `gorm:"default:null;foreignKey:ReportGradeToFailID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	NotationAverage               float64 `gorm:"default:null"`
	NotationReport                float64 `gorm:"default:null"`
	MinimumRequiredScoreToPromote float64 `gorm:"default:null"`
	OnlyFailedExams               bool    `gorm:"default:null"`
}

func (item *ReportConfig) ToResponse() *data.ReportConfigResponse {
	if item == nil {
		return &data.ReportConfigResponse{}
	}
	resp := &data.ReportConfigResponse{}
	resp.NotationAverage = item.NotationAverage
	resp.NotationReport = item.NotationReport
	resp.MinimumRequiredScoreToPromote = item.MinimumRequiredScoreToPromote
	resp.OnlyFailedExams = item.OnlyFailedExams

	resp.School = item.School.ToPublicResponse()
	resp.ReportGradeToFail = item.ReportGradeToFail.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToReportConfigResponseList(itemList []ReportConfig) []data.ReportConfigResponse {
	resp := make([]data.ReportConfigResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
