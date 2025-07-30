package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
)

type QuizQuestionOption struct {
	types.BaseGormModel
	QuizQuestionID int64         `gorm:"default:null"`
	QuizQuestion   *QuizQuestion `gorm:"foreignKey:QuizQuestionID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}
type QuizQuestionOptionWithoutFk struct {
	types.BaseGormModel
	QuizQuestionID int64 `gorm:"index"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (QuizQuestionOptionWithoutFk) TableName() string {
	return "quiz_question_options"
}

func (item *QuizQuestionOption) ToResponse() *data.QuizQuestionOptionResponse {
	if item == nil {
		return &data.QuizQuestionOptionResponse{}
	}
	resp := &data.QuizQuestionOptionResponse{}
	resp.Title = item.Title
	resp.Description = item.Description

	resp.QuizQuestionID = item.QuizQuestionID

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

func ToQuizQuestionOptionIDList(itemList []QuizQuestionOption) []int64 {
	resp := make([]int64, len(itemList))
	for index, item := range itemList {
		resp[index] = item.ID
	}
	return resp
}
