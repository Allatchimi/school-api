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
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	User   *dataUser.UserResponse           `json:"user,omitempty" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Teacher UID"`
}

type TeacherPublicResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user,omitempty" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Teacher UID"`
}

type TeacherClassSubjectUnitResponse struct {
	types.BaseGormModelResponse
	School       *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	Year         *dataYear.YearResponse           `json:"year,omitempty" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectResponse  `json:"classSubject,omitempty" required:"false" doc:"Subject for specific class"`
	Unit         *dataUnit.UnitResponse           `json:"unit,omitempty" required:"false" doc:"Unit"`
	Teacher      *TeacherPublicResponse           `json:"teacher,omitempty" required:"false" doc:"Teacher"`
}

type TeacherResponseList struct {
	types.PaginatedResponse
	Data []TeacherResponse `json:"data" required:"false" doc:"List of teacher" example:"[]"`
}

type TeacherClassSubjectUnitResponseList struct {
	types.PaginatedResponse
	Data []TeacherClassSubjectUnitResponse `json:"data" required:"false" doc:"List of subject for teacher" example:"[]"`
}
