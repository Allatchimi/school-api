package model

import (
	"api/common/types"
	"api/services/school/common/course/data"
)

type CourseDocument struct {
	types.BaseGormModel
	CourseID int64   `gorm:"default:null"`
	Course   *Course `gorm:"default:null;foreignKey:CourseID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Url         string `gorm:"default:null"`
}

func (item *CourseDocument) ToResponse() *data.CourseDocumentResponse {
	if item == nil {
		return &data.CourseDocumentResponse{}
	}
	resp := &data.CourseDocumentResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.Url = item.Url

	resp.Course = item.Course.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToCourseDocumentResponseList(itemList []CourseDocument) []data.CourseDocumentResponse {
	resp := make([]data.CourseDocumentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
