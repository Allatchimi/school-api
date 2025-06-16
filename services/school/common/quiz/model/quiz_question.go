package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
)

type QuizQuestion struct {
	types.BaseGormModel
	QuizID int64 `gorm:"default:null"`
	Quiz   *Quiz `gorm:"default:null;foreignKey:QuizID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SolutionID int64               `gorm:"default:null"`
	Solution   *QuizQuestionOption `gorm:"default:null;foreignKey:SolutionID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`

	Options []QuizQuestionOption `gorm:"foreignKey:QuizQuestionID;references:ID"`
}
type QuizQuestionWithoutFk struct {
	types.BaseGormModel
	QuizID     int64 `gorm:"index"`
	SolutionID int64 `gorm:"index"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (QuizQuestionWithoutFk) TableName() string {
	return "quiz_questions"
}

func (item *QuizQuestion) ToResponse() *data.QuizQuestionResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizQuestionResponse{}
	resp.Title = item.Title
	resp.Description = item.Description

	resp.Quiz = item.Quiz.ToResponse()
	resp.Solution = item.Solution.ToResponse()

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

func (item *QuizQuestion) ToResponseWithOptions() *data.QuizQuestionResponseV2 {
	if item == nil {
		return nil
	}
	resp := &data.QuizQuestionResponseV2{}
	resp.Question = item.ToResponse()
	resp.Options = ToQuizQuestionOptionResponseList(item.Options)
	return resp
}

func ToQuizQuestionResponseListWithOptions(itemList []QuizQuestion) []data.QuizQuestionResponseV2 {
	resp := make([]data.QuizQuestionResponseV2, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponseWithOptions()
	}
	return resp
}
