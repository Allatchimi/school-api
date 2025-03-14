package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataLevel "api/services/school/university/level/data"
	dataUser "api/services/user/user/data"
)

type StudentResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserResponse     `json:"userID" required:"false" doc:"User id"`

	UID string `json:"uid" required:"false" doc:"Student UID"`
}

type StudentLevelClassResponse struct {
	types.BaseGormModelResponse
	Student *StudentResponse         `json:"student" required:"false" doc:"Student"`
	Year    *dataYear.YearResponse   `json:"year" required:"false" doc:"Year"`
	Level   *dataLevel.LevelResponse `json:"level" required:"false" doc:"Level"`
	Class   *dataClass.ClassResponse `json:"class" required:"false" doc:"Class"`
}

type StudentResponseList struct {
	types.PaginatedResponse
	Data []StudentResponse `json:"data" required:"false" doc:"List of students" example:"[]"`
}

type StudentLevelClassResponseList struct {
	types.PaginatedResponse
	Data []StudentLevelClassResponse `json:"data" required:"false" doc:"List of students levels/classes" example:"[]"`
}
