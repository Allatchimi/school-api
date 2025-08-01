package model

import (
	"api/common/types"
	"api/services/school/common/report/data"
	modelSchool "api/services/school/common/school/model"
)

type ReportCorrespondence struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Minimum        float64 `gorm:"default:null"`
	Maximum        float64 `gorm:"default:null"`
	IncludeMinimum bool    `gorm:"default:null"`
	IncludeMaximum bool    `gorm:"default:null"`
	NewScore       float64 `gorm:"default:null"`
}

func (item *ReportCorrespondence) ToResponse() *data.ReportCorrespondenceResponse {
	if item == nil {
		return &data.ReportCorrespondenceResponse{}
	}
	resp := &data.ReportCorrespondenceResponse{}
	resp.Minimum = item.Minimum
	resp.Maximum = item.Maximum
	resp.IncludeMinimum = item.IncludeMinimum
	resp.IncludeMaximum = item.IncludeMaximum
	resp.NewScore = item.NewScore

	resp.School = item.School.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToReportCorrespondenceResponseList(itemList []ReportCorrespondence) []data.ReportCorrespondenceResponse {
	resp := make([]data.ReportCorrespondenceResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
