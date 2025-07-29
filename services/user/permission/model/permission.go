package model

import (
	"api/common/types"
	"api/services/user/permission/data"
	modelRole "api/services/user/role/model"
)

type Permission struct {
	types.BaseGormModel
	RoleID int64           `gorm:"default:null"`
	Role   *modelRole.Role `gorm:"default:null;foreignKey:RoleID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	TableName string `gorm:"default:null"`
	Create    bool   `gorm:"default:null"`
	Read      bool   `gorm:"default:null"`
	Update    bool   `gorm:"default:null"`
	Delete    bool   `gorm:"default:null"`
}

func (item *Permission) ToResponse() *data.PermissionResponse {
	if item == nil {
		return nil
	}
	resp := &data.PermissionResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Role = item.Role.ToResponse()

	resp.TableName = item.TableName
	resp.Create = item.Create
	resp.Read = item.Read
	resp.Update = item.Update
	resp.Delete = item.Delete
	return resp
}

func ToResponseList(itemList []Permission) []data.PermissionResponse {
	resp := make([]data.PermissionResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
