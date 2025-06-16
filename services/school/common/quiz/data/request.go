package data

import "api/common/types"

type QuizID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quiz id" example:"1"`
}

type QuizQuestionID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quiz question id" example:"1"`
}

type QuizRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Class subject id" example:"1"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id" example:"1"`

	Title       string `json:"title" required:"true" doc:"Title" example:""`
	Description string `json:"description" required:"false" doc:"Description" example:""`
	Status      string `json:"status" required:"false" enum:"default,active,closed"  doc:"Status" example:""`

	Questions []QuizQuestionRequest `json:"questions" required:"false" doc:"Questions" example:"[]"`
}

type QuizQuestionRequest struct {
	Title       string `json:"title" required:"true" doc:"Title" example:""`
	Description string `json:"description" required:"false" doc:"Description" example:""`

	Options []QuizQuestionOptionRequest `json:"options" required:"false" doc:"Options" example:"[]"`
}

type QuizQuestionOptionRequest struct {
	Title       string `json:"title" required:"true" doc:"Title" example:""`
	Description string `json:"description" required:"false" doc:"Description" example:""`
}

type QuizSolutionRequest struct {
	Solutions []QuestionSolutionRequest `json:"solutions" required:"true" doc:"Solutions" example:"[]"`
}

type QuestionSolutionRequest struct {
	QuestionID int64 `json:"questionID" required:"true" doc:"Question id" example:"1"`
	OptionID   int64 `json:"optionID" required:"true" doc:"Option id" example:"1"`
}

type QuizAnswerRequest struct {
	StudentID int64                   `json:"studentID" required:"true" doc:"Student id" example:"1"`
	Answers   []QuestionAnswerRequest `json:"answers" required:"true" doc:"Answers" example:"[]"`
}

type QuestionAnswerRequest struct {
	QuestionID int64 `json:"questionID" required:"true" doc:"Question id" example:"1"`
	OptionID   int64 `json:"optionID" required:"true" doc:"Option id" example:"1"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
}

type GetAllQuizAnswerRequest struct {
	QuizID
	QuizQuestionID int64 `json:"quizQuestionID" query:"quizQuestionID" required:"false" doc:"Quiz question id" example:"1"`
	StudentID      int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id" example:"1"`
}
