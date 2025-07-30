package model

import (
	"api/common/types"
	"api/services/user/role/data"
)

type Role struct {
	types.BaseGormModel
	Feature     string `gorm:"default:null"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"default:null"`
}

func (item *Role) ToResponse() *data.RoleResponse {
	if item == nil {
		return &data.RoleResponse{}
	}
	resp := &data.RoleResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Feature = item.Feature
	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Role) []data.RoleResponse {
	resp := make([]data.RoleResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
