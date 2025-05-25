package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
)

type QuizQuestion struct {
	types.BaseGormModel
	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`

	QuizID       int64 `gorm:"default:null"`
	CorrectionID int64 `gorm:"default:null"`

	Options []QuizQuestionOption `gorm:"foreignKey:QuizQuestionID;references:ID"`
}

func (item *QuizQuestion) ToResponse() *data.QuizQuestionResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizQuestionResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.QuizID = item.QuizID
	resp.CorrectionID = item.CorrectionID
	resp.Options = ToQuizQuestionOptionResponseList(item.Options)

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToQuizQuestionResponseList(itemList []QuizQuestion) []data.QuizQuestionResponse {
	resp := make([]data.QuizQuestionResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
