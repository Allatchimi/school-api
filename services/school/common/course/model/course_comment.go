package model

import (
	"api/common/types"
	"api/services/school/common/course/data"
	"api/services/user/user/model"
)

type CourseComment struct {
	types.BaseGormModel
	CourseID int64   `gorm:"default:null"`
	Course   *Course `gorm:"default:null;foreignKey:CourseID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64       `gorm:"default:null"`
	User   *model.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Message   string `gorm:"default:null;type:text"`
	Rate      int    `gorm:"default:null"`
	IsDeleted bool   `gorm:"default:null"`
}

func (item *CourseComment) ToResponse() *data.CourseCommentResponse {
	if item == nil {
		return nil
	}
	resp := &data.CourseCommentResponse{}
	resp.Message = item.Message
	resp.Rate = item.Rate
	resp.IsDeleted = item.IsDeleted

	resp.Course = item.Course.ToPublicResponse()
	resp.User = item.User.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToCourseCommentResponseList(itemList []CourseComment) []data.CourseCommentResponse {
	resp := make([]data.CourseCommentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
