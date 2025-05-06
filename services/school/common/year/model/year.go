package model

import (
	"api/common/types"
	schoolModel "api/services/school/common/school/model"
	"api/services/school/common/year/data"
	"time"
)

type Year struct {
	types.BaseGormModel
	Name      string     `gorm:"default:null"`
	StartDate *time.Time `gorm:"default:null"`
	EndDate   *time.Time `gorm:"default:null"`

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Year) ToResponse() *data.YearResponse {
	if item == nil {
		return nil
	}
	resp := &data.YearResponse{}
	resp.Name = item.Name
	resp.StartDate = item.StartDate
	resp.EndDate = item.EndDate

	resp.School = item.School.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Year) []data.YearResponse {
	resp := make([]data.YearResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
