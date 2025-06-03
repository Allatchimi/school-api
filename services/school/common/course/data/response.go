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
	CoursePublicResponse
	Documents []CourseDocumentResponse `json:"documents" required:"false" doc:"Documents"`
	Videos    []CourseVideoResponse    `json:"videos" required:"false" doc:"Videos"`
}

type CoursePublicResponse struct {
	types.BaseGormModelResponse
	School       *dataSchool.SchoolPublicResponse      `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearPublicResponse          `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Unit         *dataUnit.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	Content     string `json:"content" required:"false" doc:"Content"`
}

type CourseDocumentResponse struct {
	types.BaseGormModelResponse
	Course *CoursePublicResponse `json:"course" required:"false" doc:"Course"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	URL         string `json:"url" required:"false" doc:"URL"`
}

type CourseVideoResponse struct {
	types.BaseGormModelResponse
	Course *CoursePublicResponse `json:"course" required:"false" doc:"Course"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	URL         string `json:"url" required:"false" doc:"URL"`
}

type CourseCommentResponse struct {
	types.BaseGormModelResponse
	Course *CoursePublicResponse    `json:"course" required:"false" doc:"Course"`
	User   *data.UserPublicResponse `json:"user" required:"false" doc:"User"`

	Message   string `json:"message" required:"false" doc:"Message"`
	Rate      int    `json:"rate" required:"false" doc:"Rate"`
	IsDeleted bool   `json:"isDeleted" required:"false" doc:"Is deleted"`
}

type CourseResponseList struct {
	types.PaginatedResponse
	Data []CoursePublicResponse `json:"data" required:"false" doc:"List of courses" example:"[]"`
}

type CourseDocumentResponseList struct {
	types.PaginatedResponse
	Data []CourseDocumentResponse `json:"data" required:"false" doc:"List of documents" example:"[]"`
}

type CourseVideoResponseList struct {
	types.PaginatedResponse
	Data []CourseDocumentResponse `json:"data" required:"false" doc:"List of videos" example:"[]"`
}

type CourseCpmmentResponseList struct {
	types.PaginatedResponse
	Data []CourseCommentResponse `json:"data" required:"false" doc:"List of comments" example:"[]"`
}
