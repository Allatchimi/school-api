package data

import "api/common/types"

type QuizID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quiz id"`
}

type QuizQuestionID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quiz question id"`
}

type QuizRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Class subject id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`

	Title       string                     `json:"title" required:"true" doc:"Title"`
	Description string                     `json:"description" required:"false" doc:"Description"`
	Status      string                     `json:"status" required:"false" enum:"draft,published,closed,result" doc:"Status"`
	Questions   []QuizQuestionGroupRequest `json:"questions" required:"false" doc:"Questions"`
}

type QuizQuestionGroupRequest struct {
	Question QuizQuestionRequest         `json:"question" required:"false" doc:"Question"`
	Options  []QuizQuestionOptionRequest `json:"options" required:"false" doc:"Options"`
}

type QuizQuestionRequest struct {
	Title       string `json:"title" required:"true" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type QuizQuestionOptionRequest struct {
	Title       string `json:"title" required:"true" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type QuizSolutionRequest struct {
	Solutions []QuestionSolutionRequest `json:"solutions" required:"true" doc:"Solutions"`
}

type QuestionSolutionRequest struct {
	QuestionID int64 `json:"questionID" required:"true" doc:"Question id"`
	OptionID   int64 `json:"optionID" required:"true" doc:"Option id"`
}

type QuizAnswerRequest struct {
	StudentID int64                   `json:"studentID" required:"true" doc:"Student id"`
	Answers   []QuestionAnswerRequest `json:"answers" required:"true" doc:"Answers"`
}

type QuestionAnswerRequest struct {
	QuestionID int64 `json:"questionID" required:"true" doc:"Question id"`
	OptionID   int64 `json:"optionID" required:"true" doc:"Option id"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
}

type GetAllQuizAnswerRequest struct {
	QuizID
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	StudentID int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}

type GetAllQuizResultRequest struct {
	QuizID
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	StudentID int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}
