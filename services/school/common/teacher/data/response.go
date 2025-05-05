package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataDomain "api/services/school/university/domain/data"
	dataLevel "api/services/school/university/level/data"
	dataUser "api/services/user/user/data"
)

type TeacherResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserResponse     `json:"userID" required:"false" doc:"User id"`

	UID string `json:"uid" required:"false" doc:"Teacher UID"`
}

type TeacherLevelClassResponse struct {
	types.BaseGormModelResponse
	Teacher *TeacherResponse       `json:"teacher" required:"false" doc:"Teacher"`
	Year    *dataYear.YearResponse `json:"year" required:"false" doc:"Year"`

	Domain *dataDomain.DomainResponse `json:"domain" required:"false" doc:"Domain"`
	Level  *dataLevel.LevelResponse   `json:"level" required:"false" doc:"Level"`

	Class *dataClass.ClassResponse `json:"class" required:"false" doc:"Class"`
}

type TeacherResponseList struct {
	types.PaginatedResponse
	Data []TeacherResponse `json:"data" required:"false" doc:"List of teachers" example:"[]"`
}

type TeacherLevelClassResponseList struct {
	types.PaginatedResponse
	Data []TeacherLevelClassResponse `json:"data" required:"false" doc:"List of teachers levels/classes" example:"[]"`
}
