package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataUnit "api/services/school/university/unit/data"
	"time"
)

type QuizResponse struct {
	types.BaseGormModelResponse
	Title       string     `json:"title" required:"false" doc:"Title"`
	Description string     `json:"description" required:"false" doc:"Description"`
	StartDate   *time.Time `json:"startDate" required:"false" doc:"Start date"`
	EndDate     *time.Time `json:"endDate" required:"false" doc:"End date"`

	School       *dataSchool.SchoolPublicResponse      `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearPublicResponse          `json:"year" required:"false" doc:"Year"`
	Unit         *dataUnit.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`
	ClassSubject *dataClass.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`

	Questions []QuizQuestionResponse `json:"questions" required:"false" doc:"Questions"`
}

type QuizQuestionResponse struct {
	types.BaseGormModelResponse
	Title        string `json:"title" required:"false" doc:"Title"`
	Description  string `json:"description" required:"false" doc:"Description"`
	QuizID       int64  `json:"quizID" required:"false" doc:"Quiz id"`
	CorrectionID int64  `json:"correctionID" required:"false" doc:"Correct choice id"`

	Options []QuizQuestionOptionResponse `json:"options" required:"false" doc:"Options"`
}

type QuizQuestionOptionResponse struct {
	types.BaseGormModelResponse
	Title          string `json:"title" required:"false" doc:"Title"`
	Description    string `json:"description" required:"false" doc:"Description"`
	QuizQuestionID int64  `json:"quizQuestionID" required:"false" doc:"Quiz question id"`
}

type QuizAttemptResponse struct {
	types.BaseGormModelResponse
	QuizID    int64 `json:"quizID" required:"false" doc:"Quiz id"`
	StudentID int64 `json:"studentID" required:"false" doc:"Student id"`

	Answers []QuizAttemptQuestionResponse `json:"answers" required:"false" doc:"Answers"`
}

type QuizAttemptQuestionResponse struct {
	types.BaseGormModelResponse
	QuizAttemptID int64 `json:"quizAttemptID" required:"false" doc:"Quiz atrempt id"`

	QuizQuestion       *QuizQuestionResponse       `json:"quizQuestion" required:"false" doc:"Quiz question id"`
	QuizQuestionOption *QuizQuestionOptionResponse `json:"quizQuestionOption" required:"false" doc:"Quiz question option id"`
}

type QuizResponseList struct {
	types.PaginatedResponse
	Data []QuizResponse `json:"data" required:"false" doc:"List of quiz" example:"[]"`
}
