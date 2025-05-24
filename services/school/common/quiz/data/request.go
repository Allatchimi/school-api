package data

type QuizID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quiz id" example:"1"`
}

type QuizRequest struct {
	UserID   int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
}
