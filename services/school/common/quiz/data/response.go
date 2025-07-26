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
	School       *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearResponse           `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectResponse  `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Unit         *dataUnit.UnitResponse           `json:"unit" required:"false" doc:"Unit"`
	Questions    []QuizQuestionListResponse       `json:"questions" required:"false" doc:"Questions"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
	Status      string `json:"status" required:"false" doc:"Status" example:""`
}

type QuizQuestionListResponse struct {
	Question *QuizQuestionResponse        `json:"question" required:"false" doc:"Question"`
	Options  []QuizQuestionOptionResponse `json:"options" required:"false" doc:"Options"`
}

type QuizQuestionResponse struct {
	types.BaseGormModelResponse
	QuizID   int64                       `json:"quizID" required:"false" doc:"Quiz id"`
	Solution *QuizQuestionOptionResponse `json:"solution" required:"false" doc:"Solution"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type QuizQuestionOptionResponse struct {
	types.BaseGormModelResponse
	QuizQuestionID int64 `json:"quizQuestionID" required:"false" doc:"Quiz question id"`

	Title       string `json:"title" required:"false" doc:"Title"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type QuizAnswerResponse struct {
	types.BaseGormModelResponse
	QuizID  int64                              `json:"quizID" required:"false" doc:"Quiz id"`
	Student *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`

	Answers []QuizAnswersResponse `json:"answers" required:"false" doc:"Answers"`
}

type QuizAnswersResponse struct {
	Question *QuizQuestionResponse       `json:"question" required:"false" doc:"Question" example:""`
	Answer   *QuizQuestionOptionResponse `json:"answer" required:"false" doc:"Answer" example:""`
}

type QuizResultResponse struct {
	Student *dataStudent.StudentPublicResponse `json:"student" required:"false" doc:"Student"`
	Result  float64                            `json:"result" required:"false" doc:"Result"`
}

type QuizResponseList struct {
	types.PaginatedResponse
	Data []QuizResponse `json:"data" required:"false" doc:"List of quiz"`
}

type QuizAnswerResponseList struct {
	types.PaginatedResponse
	Data []QuizAnswerResponse `json:"data" required:"false" doc:"List of quiz answer"`
}

type QuizResultResponseList struct {
	types.PaginatedResponse
	Data []QuizResultResponse `json:"data" required:"false" doc:"List of quiz result"`
}
