package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
)

type QuizAttemptQuestion struct {
	types.BaseGormModel
	QuizAttemptID int64 `gorm:"default:null"`

	QuizQuestionID int64         `gorm:"default:null"`
	QuizQuestion   *QuizQuestion `gorm:"default:null;foreignKey:QuizQuestionID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	QuizQuestionOptionID int64               `gorm:"default:null"`
	QuizQuestionOption   *QuizQuestionOption `gorm:"default:null;foreignKey:QuizQuestionOptionID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *QuizAttemptQuestion) ToResponse() *data.QuizAttemptQuestionResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizAttemptQuestionResponse{}
	resp.QuizAttemptID = item.QuizAttemptID
	resp.QuizQuestion = item.QuizQuestion.ToResponse()
	resp.QuizQuestionOption = item.QuizQuestionOption.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToQuizAttemptQuestionResponseList(itemList []QuizAttemptQuestion) []data.QuizAttemptQuestionResponse {
	resp := make([]data.QuizAttemptQuestionResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
