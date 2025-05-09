package model

import (
	"api/common/types"
	"api/services/user/role/data"
)

type Role struct {
	types.BaseGormModel
	Name        string `gorm:"unique;not null"`
	Feature     string `gorm:"not null"`
	Description string `gorm:"default:null"`
}

func (item *Role) ToResponse() *data.RoleResponse {
	if item == nil {
		return nil
	}
	resp := &data.RoleResponse{}
	resp.Name = item.Name
	resp.Feature = item.Feature
	resp.Description = item.Description

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Role) ToPublicResponse() *data.RolePublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.RolePublicResponse{}
	resp.Name = item.Name
	resp.Feature = item.Feature
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
