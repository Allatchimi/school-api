package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataLevel "api/services/school/university/level/data"
	dataUser "api/services/user/user/data"
)

type StudentResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserResponse     `json:"userID" required:"false" doc:"User id"`
}

type StudentLevelResponse struct {
	types.BaseGormModelResponse
	Student *StudentResponse         `json:"student" required:"false" doc:"Student"`
	Year    *dataYear.YearResponse   `json:"year" required:"false" doc:"Year"`
	Level   *dataLevel.LevelResponse `json:"level" required:"false" doc:"Level"`
}

type StudentResponseList struct {
	types.PaginatedResponse
	Data []StudentResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
