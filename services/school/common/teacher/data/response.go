package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataUnit "api/services/school/university/unit/data"
	dataUser "api/services/user/user/data"
)

type TeacherResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Teacher UID"`
}

type TeacherPublicResponse struct {
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Teacher UID"`
}

type TeacherUnitSubjectResponse struct {
	types.BaseGormModelResponse
	Teacher *TeacherPublicResponse       `json:"teacher" required:"false" doc:"Teacher"`
	Year    *dataYear.YearPublicResponse `json:"year" required:"false" doc:"Year"`

	Unit         *dataUnit.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`
	ClassSubject *dataClass.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`
}

type TeacherUnitSubjectPublicResponse struct {
	Teacher *TeacherPublicResponse       `json:"teacher" required:"false" doc:"Teacher"`
	Year    *dataYear.YearPublicResponse `json:"year" required:"false" doc:"Year"`

	Unit         *dataUnit.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`
	ClassSubject *dataClass.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`
}

type TeacherResponseList struct {
	types.PaginatedResponse
	Data []TeacherResponse `json:"data" required:"false" doc:"List of teachers" example:"[]"`
}

type TeacherUnitSubjectResponseList struct {
	types.PaginatedResponse
	Data []TeacherUnitSubjectResponse `json:"data" required:"false" doc:"List of unit/subject for the specified teacher" example:"[]"`
}
