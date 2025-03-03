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
	User   *dataUser.UserResponse     `json:"userID" required:"true" doc:"User id"`
	Year   *dataYear.YearResponse     `json:"levelID" required:"true" doc:"Level id"`
}

type StudentLevelResponse struct {
	types.BaseGormModelResponse
	Year  *dataYear.YearResponse   `json:"year" required:"true" doc:"Year"`
	Level *dataLevel.LevelResponse `json:"level" required:"true" doc:"Level"`
}

type StudentResponseList struct {
	types.PaginatedResponse
	Data []StudentResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
