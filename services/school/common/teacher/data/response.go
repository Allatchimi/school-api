package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataSubject "api/services/school/highschool/subject/data"
	dataUnit "api/services/school/university/unit/data"
	dataUser "api/services/user/user/data"
)

type TeacherResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Teacher UID"`
}

type TeacherUnitSubjectResponse struct {
	types.BaseGormModelResponse
	Teacher *TeacherResponse       `json:"teacher" required:"false" doc:"Teacher"`
	Year    *dataYear.YearResponse `json:"year" required:"false" doc:"Year"`

	Unit    *dataUnit.UnitResponse       `json:"domain" required:"false" doc:"Unit"`
	Subject *dataSubject.SubjectResponse `json:"subject" required:"false" doc:"Subject"`
}

type TeacherResponseList struct {
	types.PaginatedResponse
	Data []TeacherResponse `json:"data" required:"false" doc:"List of teachers" example:"[]"`
}

type TeacherUnitSubjectResponseList struct {
	types.PaginatedResponse
	Data []TeacherUnitSubjectResponse `json:"data" required:"false" doc:"List of unit/subject for the specified teacher" example:"[]"`
}
