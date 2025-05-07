package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolConfig struct {
	types.BaseGormModel
	EmailDomain  string `gorm:"default:null"`
	ColorPrimary string `gorm:"default:null"`
}

func (item *SchoolConfig) ToResponse() *data.SchoolConfigResponse {
	if item == nil {
		return nil
	}
	resp := &data.SchoolConfigResponse{}
	resp.EmailDomain = item.EmailDomain
	resp.ColorPrimary = item.ColorPrimary
	return resp
}

func FromConfigRequest(item *data.SchoolConfigRequest) *SchoolConfig {
	resp := &SchoolConfig{
		EmailDomain:  item.EmailDomain,
		ColorPrimary: item.ColorPrimary,
	}
	return resp
}
