package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type School struct {
	types.BaseGormModel
	Name         string `gorm:"unique;not null"`
	Type         string `gorm:"not null"`
	Logo         string `gorm:"default:null"`
	Currency     string `gorm:"default:null"`
	PaymentCount int64  `gorm:"default:1"`

	ConfigID int64         `gorm:"default:null"`
	Config   *SchoolConfig `gorm:"default:null;foreignKey:ConfigID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	InfoID int64       `gorm:"default:null"`
	Info   *SchoolInfo `gorm:"default:null;foreignKey:InfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *School) ToResponse() *data.SchoolResponse {
	if item == nil {
		return nil
	}
	resp := &data.SchoolResponse{}
	resp.Name = item.Name
	resp.Type = item.Type
	resp.Logo = item.Logo
	resp.Currency = item.Currency
	resp.PaymentCount = item.PaymentCount

	resp.Info = item.Info.ToResponse()
	resp.Config = item.Config.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *School) ToPublicResponse() *data.SchoolPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.SchoolPublicResponse{}
	resp.Name = item.Name
	resp.Type = item.Type
	resp.Logo = item.Logo
	resp.Currency = item.Currency
	resp.PaymentCount = item.PaymentCount

	resp.Info = item.Info.ToResponse()
	return resp
}

func ToSchoolResponseList(itemList []School) []data.SchoolResponse {
	resp := make([]data.SchoolResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
