package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataUnit "api/services/school/university/unit/data"
	"api/services/user/user/data"
)

type CourseResponse struct {
	types.BaseGormModelResponse
	School       *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearResponse           `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectResponse  `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Unit         *dataUnit.UnitResponse           `json:"unit" required:"false" doc:"Unit"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	Content     string `json:"content" required:"false" doc:"Content"`

	Documents []CourseDocumentResponse `json:"documents" required:"false" doc:"Documents"`
	Videos    []CourseVideoResponse    `json:"videos" required:"false" doc:"Videos"`
}

type CourseDocumentResponse struct {
	types.BaseGormModelResponse
	Course *CourseResponse `json:"course" required:"false" doc:"Course"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	URL         string `json:"url" required:"false" doc:"URL"`
}

type CourseVideoResponse struct {
	types.BaseGormModelResponse
	Course *CourseResponse `json:"course" required:"false" doc:"Course"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	URL         string `json:"url" required:"false" doc:"URL"`
}

type CourseCommentResponse struct {
	types.BaseGormModelResponse
	Course *CourseResponse          `json:"course" required:"false" doc:"Course"`
	User   *data.UserPublicResponse `json:"user" required:"false" doc:"User"`

	Message   string `json:"message" required:"false" doc:"Message"`
	Rate      int    `json:"rate" required:"false" doc:"Rate"`
	IsDeleted bool   `json:"isDeleted" required:"false" doc:"Is deleted"`
}

type CourseResponseList struct {
	types.PaginatedResponse
	Data []CourseResponse `json:"data" required:"false" doc:"List of course"`
}

type CourseDocumentResponseList struct {
	types.PaginatedResponse
	Data []CourseDocumentResponse `json:"data" required:"false" doc:"List of document"`
}

type CourseVideoResponseList struct {
	types.PaginatedResponse
	Data []CourseDocumentResponse `json:"data" required:"false" doc:"List of video"`
}

type CourseCpmmentResponseList struct {
	types.PaginatedResponse
	Data []CourseCommentResponse `json:"data" required:"false" doc:"List of comment"`
}
