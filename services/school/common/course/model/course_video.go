package model

import (
	"api/common/types"
	"api/services/school/common/course/data"
)

type CourseVideo struct {
	types.BaseGormModel
	CourseID int64   `gorm:"default:null"`
	Course   *Course `gorm:"default:null;foreignKey:CourseID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Url         string `gorm:"default:null"`
}

func (item *CourseVideo) ToResponse() *data.CourseVideoResponse {
	if item == nil {
		return &data.CourseVideoResponse{}
	}
	resp := &data.CourseVideoResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.Url = item.Url

	resp.Course = item.Course.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToCourseVideoResponseList(itemList []CourseVideo) []data.CourseVideoResponse {
	resp := make([]data.CourseVideoResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
