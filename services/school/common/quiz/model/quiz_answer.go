package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
	modelStudent "api/services/school/common/student/model"
	"sort"
)

type QuizAnswer struct {
	types.BaseGormModel
	StudentID int64                 `gorm:"default:null"`
	Student   *modelStudent.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	QuizQuestionID int64         `gorm:"default:null"`
	QuizQuestion   *QuizQuestion `gorm:"default:null;foreignKey:QuizQuestionID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	QuizQuestionOptionID int64               `gorm:"default:null"`
	QuizQuestionOption   *QuizQuestionOption `gorm:"default:null;foreignKey:QuizQuestionOptionID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}
type QuizAnswerWithoutFk struct {
	types.BaseGormModel
	StudentID            int64 `gorm:"index"`
	QuizQuestionID       int64 `gorm:"index"`
	QuizQuestionOptionID int64 `gorm:"index"`
}

func (QuizAnswerWithoutFk) TableName() string {
	return "quiz_answers"
}

func ToQuizAnswerResponseList(itemList []QuizAnswer) []data.QuizAnswerResponse {
	studentResponsesMap := make(map[int64]*data.QuizAnswerResponse)

	for _, qa := range itemList {
		// Skip nil student
		if qa.Student == nil {
			continue
		}

		studentID := qa.StudentID

		// Check if student already exists
		qar, exists := studentResponsesMap[studentID]
		if !exists {
			// Create new answer
			qar = &data.QuizAnswerResponse{
				QuizID:  qa.QuizQuestion.QuizID,
				Student: qa.Student.ToPublicResponse(),
				Answers: []data.QuizAnswersResponse{},
			}
			studentResponsesMap[studentID] = qar
		}

		// Add question and answer to the student
		if qa.QuizQuestion != nil && qa.QuizQuestionOption != nil {
			qar.Answers = append(qar.Answers, data.QuizAnswersResponse{
				Question: qa.QuizQuestion.ToResponse(),
				Answer:   qa.QuizQuestionOption.ToResponse(),
			})
		}
	}

	// Convert the map to a slice
	result := make([]data.QuizAnswerResponse, 0, len(studentResponsesMap))
	for _, qar := range studentResponsesMap {
		result = append(result, *qar)
	}

	return result
}

func ToQuizAnswerResultResponseList(itemList []QuizAnswer) []data.QuizResultResponse {
	resultsMap := make(map[int64]*data.QuizResultResponse)

	for _, qa := range itemList {
		// Skip nil student
		if qa.Student == nil {
			continue
		}
		student := qa.Student

		// Check if student already exists
		sr, exists := resultsMap[student.ID]
		if !exists {
			// Create new student notation
			sr = &data.QuizResultResponse{
				Student: student.ToPublicResponse(),
				Result:  0,
			}
			resultsMap[student.ID] = sr
		}

		// Increment notation if answer is correct
		if qa.QuizQuestion != nil && qa.QuizQuestionOption != nil {
			if qa.QuizQuestionOption.ID == qa.QuizQuestion.SolutionID {
				sr.Result += 1
			}
		}
	}

	// Convert the map to a slice
	result := make([]data.QuizResultResponse, 0, len(resultsMap))
	for _, sr := range resultsMap {
		result = append(result, *sr)
	}

	// Sort desc by result
	sort.Slice(result, func(i, j int) bool {
		return result[i].Result > result[j].Result
	})

	return result
}
