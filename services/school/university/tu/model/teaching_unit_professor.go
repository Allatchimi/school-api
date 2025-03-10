package model

import (
	"api/common/types"
	"api/services/school/university/tu/data"
	modelUser "api/services/user/user/model"
)

type TeachingUnitProfessor struct {
	types.BaseGormModel
	TeachingUnitID int64         `gorm:"default:null"`
	TeachingUnit   *TeachingUnit `gorm:"default:null;foreignKey:TeachingUnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *TeachingUnitProfessor) ToResponse() *data.TeachingUnitProfessorResponse {
	resp := &data.TeachingUnitProfessorResponse{}
	resp.TeachingUnitID = item.TeachingUnitID
	resp.UserID = item.UserID
	return resp
}

func ToTeachingUnitProfessorResponseList(itemList []TeachingUnitProfessor) []data.TeachingUnitProfessorResponse {
	resp := make([]data.TeachingUnitProfessorResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
