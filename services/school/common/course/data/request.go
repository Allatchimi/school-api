package data

import "api/common/types"

type CourseID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Course id"`
}

type CourseDocumentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Course document id"`
}

type CourseCommentID struct {
	ID int64 `json:"commentID" path:"commentID" required:"true" doc:"Course document id"`
}

type CourseRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Subject class id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`

	Title       string `json:"title" required:"true" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	Content     string `json:"content" required:"false" doc:"Content"`

	Documents []struct {
		Title       string `json:"title" required:"true" doc:"Title"`
		Description string `json:"description" required:"false" doc:"Description"`
		Url         string `json:"url" required:"false" doc:"Url"`
	} `json:"documents" required:"false" doc:"Documents"`

	Videos []struct {
		Title       string `json:"title" required:"true" doc:"Title"`
		Description string `json:"description" required:"false" doc:"Description"`
		Url         string `json:"url" required:"false" doc:"Url"`
	} `json:"videos" required:"false" doc:"Videos"`
}

type CourseCommentRequest struct {
	Message string `json:"message" required:"true" doc:"Message"`
	Rate    int    `json:"rate" required:"false" doc:"Rate"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
}

type GetAllCourseCommentRequest struct {
	SchoolID int64 `json:"schoolID" required:"false" doc:"School id"`
	CourseID int64 `json:"courseID" required:"false" doc:"Course id"`
}
