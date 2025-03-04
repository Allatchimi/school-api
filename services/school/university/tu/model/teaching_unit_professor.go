package model

import (
	"api/common/types"
	"api/services/school/university/tu/data"
)

type TeachingUnitProfessor struct {
	types.BaseGormModel
	TeachingUnitID int64 `gorm:"default:null"`
	UserID         int64 `gorm:"default:null"`
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
