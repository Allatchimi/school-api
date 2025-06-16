package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataUnit "api/services/school/university/unit/data"
)

type QuizResponse struct {
	types.BaseGormModelResponse
	School       *dataSchool.SchoolPublicResponse      `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearPublicResponse          `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Unit         *dataUnit.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`
	Questions    []QuizQuestionResponseV2              `json:"questions" required:"false" doc:"Questions"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	Status      string `json:"status" required:"false" doc:"Status" example:""`
}

type QuizQuestionResponseV2 struct {
	Question *QuizQuestionResponse        `json:"question" required:"false" doc:"Question"`
	Options  []QuizQuestionOptionResponse `json:"options" required:"false" doc:"Options"`
}

type QuizQuestionResponse struct {
	types.BaseGormModelResponse
	Quiz     *QuizResponse               `json:"quiz" required:"false" doc:"Quiz"`
	Solution *QuizQuestionOptionResponse `json:"solution" required:"false" doc:"Solution"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type QuizQuestionOptionResponse struct {
	types.BaseGormModelResponse
	QuizQuestion *QuizQuestionResponse `json:"quizQuestion" required:"false" doc:"Quiz question"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type QuizAnswerResponse struct {
	types.BaseGormModelResponse
	Quiz    *QuizResponse                      `json:"quiz" required:"false" doc:"Quiz"`
	Student *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`

	Answers []QuizAnswersResponse `json:"answers" required:"false" doc:"Answers"`
}

type QuizAnswersResponse struct {
	Question *QuizQuestionResponse       `json:"question" required:"false" doc:"Question" example:""`
	Answer   *QuizQuestionOptionResponse `json:"answer" required:"false" doc:"Answer" example:""`
}

type QuizResponseList struct {
	types.PaginatedResponse
	Data []QuizResponse `json:"data" required:"false" doc:"List of quiz" example:"[]"`
}

type QuizQuestionOptionResponseList struct {
	types.PaginatedResponse
	Data []QuizQuestionOptionResponse `json:"data" required:"false" doc:"List of quiz question options" example:"[]"`
}

type QuizAnswerResponseList struct {
	types.PaginatedResponse
	Data []QuizAnswerResponse `json:"data" required:"false" doc:"List of quiz answers" example:"[]"`
}
