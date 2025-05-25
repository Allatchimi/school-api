package data

import "time"

type QuizID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quiz id" example:"1"`
}

type QuizRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id" example:"1"`
	ClassSubjectID int64 `json:"ClassSubjectID" required:"false" doc:"Class subject id" example:"1"`

	Title       string     `json:"Title" required:"true" doc:"Title" example:"1"`
	Description string     `json:"Description" required:"false" doc:"Description" example:"1"`
	StartDate   *time.Time `json:"startDate" required:"false" doc:"Start date" example:""`
	EndDate     *time.Time `json:"endDate" required:"false" doc:"End date" example:""`

	Questions []QuizQuestionRequest `json:"questions" required:"false" doc:"Questions" example:"[]"`
}

type QuizQuestionRequest struct {
	Title       string                      `json:"Title" required:"true" doc:"Title" example:"1"`
	Description string                      `json:"Description" required:"false" doc:"Description" example:"1"`
	Options     []QuizQuestionOptionRequest `json:"options" required:"false" doc:"Options" example:"[]"`
}

type QuizQuestionOptionRequest struct {
	Title       string `json:"Title" required:"true" doc:"Title" example:"1"`
	Description string `json:"Description" required:"false" doc:"Description" example:"1"`
}

type QuizCorrectionRequest struct {
	UserID   int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
}

type QuizAttemptRequest struct {
	UserID   int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
}
