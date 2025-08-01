package model

import (
	"api/common/types"
	"api/services/school/common/manager/data"
	modelSchool "api/services/school/common/school/model"
	modelUser "api/services/user/user/model"
)

type Manager struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UID string `gorm:"default:null"`
}

func (item *Manager) ToResponse() *data.ManagerResponse {
	if item == nil {
		return &data.ManagerResponse{}
	}
	resp := &data.ManagerResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.User = item.User.ToResponse()

	resp.UID = item.UID
	return resp
}

func (item *Manager) ToPublicResponse() *data.ManagerPublicResponse {
	if item == nil {
		return &data.ManagerPublicResponse{}
	}
	resp := &data.ManagerPublicResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.User = item.User.ToPublicResponse()

	resp.UID = item.UID
	return resp
}

func ToManagerResponseList(itemList []Manager) []data.ManagerResponse {
	resp := make([]data.ManagerResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
