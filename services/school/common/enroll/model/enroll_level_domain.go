package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type EnrollLevelDomain struct {
	types.BaseGormModel
	Name     string `gorm:"unique;not null"`
	Type     string `gorm:"not null"`
	Logo     string `gorm:"default:null"`
	Currency string `gorm:"default:null"`

	EnrollLevelDomainConfigID int64                    `gorm:"default:null"`
	Config                    *EnrollLevelDomainConfig `gorm:"default:null;foreignKey:EnrollLevelDomainConfigID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	EnrollLevelDomainInfoID int64                  `gorm:"default:null"`
	Info                    *EnrollLevelDomainInfo `gorm:"default:null;foreignKey:EnrollLevelDomainInfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *EnrollLevelDomain) ToResponse() *data.EnrollLevelDomainResponse {
	if item == nil {
		return nil
	}
	resp := &data.EnrollLevelDomainResponse{}
	resp.Name = item.Name
	resp.Type = item.Type
	resp.Logo = item.Logo
	resp.Currency = item.Currency

	resp.Info = item.Info.ToResponse()
	resp.Config = item.Config.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *EnrollLevelDomain) ToPublicResponse() *data.EnrollLevelDomainPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.EnrollLevelDomainPublicResponse{}
	resp.Name = item.Name
	resp.Type = item.Type
	resp.Logo = item.Logo
	resp.Currency = item.Currency

	resp.Info = item.Info.ToResponse()
	return resp
}

func ToEnrollLevelDomainResponseList(itemList []EnrollLevelDomain) []data.EnrollLevelDomainResponse {
	resp := make([]data.EnrollLevelDomainResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
