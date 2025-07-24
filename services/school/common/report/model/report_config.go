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

	Notation               float64 `gorm:"default:null"`
	NotationMinimumSuccess float64 `gorm:"default:null"`
}

func (item *ReportConfig) ToResponse() *data.ReportConfigResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportConfigResponse{}
	resp.Notation = item.Notation
	resp.NotationMinimumSuccess = item.NotationMinimumSuccess

	resp.School = item.School.ToPublicResponse()

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
