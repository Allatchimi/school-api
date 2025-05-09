package model

import (
	"api/common/types"
	"api/services/user/permission/data"
	"api/services/user/role/model"
)

type Permission struct {
	types.BaseGormModel
	TableName string `gorm:"default:null"`
	Create    bool   `gorm:"default:null"`
	Read      bool   `gorm:"default:null"`
	Update    bool   `gorm:"default:null"`
	Delete    bool   `gorm:"default:null"`

	RoleID int64       `gorm:"default:null"`
	Role   *model.Role `gorm:"default:null;foreignKey:RoleID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Permission) ToResponse() *data.PermissionResponse {
	if item == nil {
		return nil
	}
	resp := &data.PermissionResponse{}
	resp.TableName = item.TableName
	resp.Create = item.Create
	resp.Read = item.Read
	resp.Update = item.Update
	resp.Delete = item.Delete

	resp.Role = item.Role.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Permission) ToPublicResponse() *data.PermissionPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.PermissionPublicResponse{}
	resp.TableName = item.TableName
	resp.Create = item.Create
	resp.Read = item.Read
	resp.Update = item.Update
	resp.Delete = item.Delete

	resp.Role = item.Role.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []Permission) []data.PermissionResponse {
	resp := make([]data.PermissionResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
