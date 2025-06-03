package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
)

type QuizQuestionOption struct {
	types.BaseGormModel
	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`

	QuizQuestionID int64         `gorm:"default:null"`
	QuizQuestion   *QuizQuestion `gorm:"default:null;foreignKey:QuizQuestionID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *QuizQuestionOption) ToResponse() *data.QuizQuestionOptionResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizQuestionOptionResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.QuizQuestion = item.QuizQuestion.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToQuizQuestionOptionResponseList(itemList []QuizQuestionOption) []data.QuizQuestionOptionResponse {
	resp := make([]data.QuizQuestionOptionResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
