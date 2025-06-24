package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
)

type StatisticResponse struct {
	Count                 *CountResponse                  `json:"count,omitempty" required:"false" doc:"Count of schools, directors, teachers, students, and parents"`
	UsersByYear           []UsersByYearResponse           `json:"usersByYear,omitempty" required:"false" doc:"Users by year"`
	SuccessBySchool       []SuccessBySchoolResponse       `json:"successBySchool,omitempty" required:"false" doc:"Success by school"`
	SuccessBySchoolGender []SuccessBySchoolGenderResponse `json:"successBySchoolGender,omitempty" required:"false" doc:"Success by school gender"`
}

type CountResponse struct {
	Schools   int64 `json:"schools,omitempty" required:"false" doc:"School count"`
	Directors int64 `json:"directors,omitempty" required:"false" doc:"Director count"`
	Teachers  int64 `json:"teachers,omitempty" required:"false" doc:"Teacher count"`
	Students  int64 `json:"students,omitempty" required:"false" doc:"Student count"`
	Parents   int64 `json:"parents,omitempty" required:"false" doc:"Parent count"`
}

type UsersByYearResponse struct {
	Year  int64 `json:"year,omitempty" required:"false" doc:"Year"`
	Total int64 `json:"total,omitempty" required:"false" doc:"Total"`
}

type SuccessBySchoolResponse struct {
	School  *dataSchool.SchoolResponse `json:"school,omitempty" required:"false" doc:"School"`
	Success int64                      `json:"success,omitempty" required:"false" doc:"Success"`
}

type SuccessBySchoolGenderResponse struct {
	School *dataSchool.SchoolResponse `json:"school,omitempty" required:"false" doc:"School"`
	Boys   int64                      `json:"boys,omitempty" required:"false" doc:"Boys"`
	Girls  int64                      `json:"girls,omitempty" required:"false" doc:"Girls"`
}

type StatisticResponseList struct {
	types.PaginatedResponse
	Data *StatisticResponse `json:"data" required:"false" doc:"Statistics"`
}
