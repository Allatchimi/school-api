package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
)

type MonitoringResponse struct {
	Count                 *CountResponse                  `json:"count,omitempty" required:"false" doc:"Count of schools, directors, teachers, students, and parents"`
	UsersByFeature        []UsersByFeatureResponse        `json:"usersByFeature,omitempty" required:"false" doc:"Users by feature"`
	UsersByGender         []UsersByGenderResponse         `json:"usersByGender,omitempty" required:"false" doc:"Users by gender"`
	UsersByMonth          []UsersByMonthResponse          `json:"usersByMonth,omitempty" required:"false" doc:"Users by month"`
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
type UsersByFeatureResponse struct {
	Feature string `json:"feature,omitempty" required:"false" doc:"Feature"`
	Count   int64  `json:"count,omitempty" required:"false" doc:"Count"`
}

type UsersByGenderResponse struct {
	Gender string `json:"gender,omitempty" required:"false" doc:"Gender"`
	Count  int64  `json:"count,omitempty" required:"false" doc:"Count"`
}

type UsersByMonthResponse struct {
	Month int64 `json:"month,omitempty" required:"false" doc:"Month"`
	Count int64 `json:"count,omitempty" required:"false" doc:"Count"`
}

type UsersByYearResponse struct {
	Year  int64 `json:"year,omitempty" required:"false" doc:"Year"`
	Count int64 `json:"count,omitempty" required:"false" doc:"Count"`
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

type MonitoringResponseList struct {
	types.PaginatedResponse
	Data *MonitoringResponse `json:"data" required:"false" doc:"Statistics"`
}
