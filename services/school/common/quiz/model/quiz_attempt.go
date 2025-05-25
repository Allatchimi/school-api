package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
)

type QuizAttempt struct {
	types.BaseGormModel
	StudentID int64 `gorm:"default:null"`
	QuizID    int64 `gorm:"default:null"`

	Answers []QuizAttemptQuestion `gorm:"foreignKey:QuizAttemptID;references:ID"`
}

func (item *QuizAttempt) ToResponse() *data.QuizAttemptResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizAttemptResponse{}
	resp.StudentID = item.StudentID
	resp.QuizID = item.QuizID

	resp.Answers = ToQuizAttemptQuestionResponseList(item.Answers)

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToQuizAttemptResponseList(itemList []QuizAttempt) []data.QuizAttemptResponse {
	resp := make([]data.QuizAttemptResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
