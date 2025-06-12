package data

import "api/common/types"

type CourseID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Course id" example:"1"`
}

type CourseDocumentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Course document id" example:"1"`
}

type CourseRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"true" doc:"Subject class id"`
	UnitID         int64 `json:"unitID" required:"true" doc:"Unit id"`

	Title       string `json:"title" required:"true" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	Content     string `json:"content" required:"false" doc:"Content"`

	Documents []struct {
		Title       string `json:"title" required:"true" doc:"Title" example:""`
		Description string `json:"description" required:"false" doc:"Description" example:""`
		Url         string `json:"url" required:"false" doc:"Url" example:"[]"`
	} `json:"documents" required:"false" doc:"Documents" example:"[]"`

	Videos []struct {
		Title       string `json:"title" required:"true" doc:"Title" example:""`
		Description string `json:"description" required:"false" doc:"Description" example:""`
		Url         string `json:"url" required:"false" doc:"Url" example:"[]"`
	} `json:"videos" required:"false" doc:"Videos" example:"[]"`
}

type CourseDocumentRequest struct {
	CourseID int64 `json:"courseID" required:"true" doc:"Course id"`

	Title       string `json:"title" required:"true" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	URL         string `json:"url" required:"true" doc:"URL"`
}

type CourseCommentRequest struct {
	CourseID int64 `json:"courseID" required:"true" doc:"Course id"`

	Message string `json:"message" required:"true" doc:"Message"`
	Rate    int    `json:"rate" required:"false" doc:"Rate"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
}

type GetAllCourseDocumentRequest struct {
	CourseID int64 `json:"courseID" required:"false" doc:"Course id"`
}

type GetAllCourseCommentRequest struct {
	CourseID int64 `json:"courseID" required:"false" doc:"Course id"`
	UserID   int64 `json:"userID" required:"false" doc:"User id"`
}
