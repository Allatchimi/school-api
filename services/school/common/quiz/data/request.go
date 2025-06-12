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

	Questions []struct {
		Title       string `json:"title" required:"true" doc:"Title" example:""`
		Description string `json:"description" required:"false" doc:"Description" example:""`

		Options []struct {
			Title       string `json:"title" required:"true" doc:"Title" example:""`
			Description string `json:"description" required:"false" doc:"Description" example:""`
		} `json:"options" required:"false" doc:"Options" example:"[]"`
	} `json:"questions" required:"false" doc:"Questions" example:"[]"`
}

type QuizSolutionRequest struct {
	Solutions []struct {
		QuestionID int64 `json:"questionID" required:"true" doc:"Question id" example:"1"`
		SolutionID int64 `json:"solutionID" required:"true" doc:"Solution id" example:"1"`
	} `json:"solutions" required:"true" doc:"Solutions" example:"[]"`
}

type QuizAnswerRequest struct {
	StudentID int64 `json:"studentID" required:"true" doc:"Student id" example:"1"`
	Answers   []struct {
		QuestionID int64 `json:"questionID" required:"true" doc:"Question id" example:"1"`
		OptionID   int64 `json:"optionID" required:"true" doc:"Option id" example:"1"`
	} `json:"Answers" required:"true" doc:"Answers" example:"[]"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
}

type GetAllQuizQuestionOptionRequest struct {
	QuizQuestionID
}

type GetAllQuizAnswerRequest struct {
	QuizID
	QuizQuestionID int64 `json:"quizQuestionID" query:"quizQuestionID" required:"false" doc:"Quiz question id" example:"1"`
	StudentID      int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id" example:"1"`
}
