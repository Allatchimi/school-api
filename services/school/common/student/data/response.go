package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataLevelDomain "api/services/school/university/level/data"
	dataUser "api/services/user/user/data"
)

type StudentResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Student UID"`
}

type StudentPublicResponse struct {
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Student UID"`
}

type StudentEnrollResponse struct {
	types.BaseGormModelResponse
	Student *StudentPublicResponse       `json:"student" required:"false" doc:"Student"`
	Year    *dataYear.YearPublicResponse `json:"year" required:"false" doc:"Year"`

	LevelDomain *dataLevelDomain.LevelDomainPublicResponse `json:"levelDomain" required:"false" doc:"Level for domain"`
	Class       *dataClass.ClassPublicResponse             `json:"class" required:"false" doc:"Class"`
}

type StudentEnrollPublicResponse struct {
	Student *StudentPublicResponse       `json:"student" required:"false" doc:"Student"`
	Year    *dataYear.YearPublicResponse `json:"year" required:"false" doc:"Year"`

	LevelDomain *dataLevelDomain.LevelDomainPublicResponse `json:"levelDomain" required:"false" doc:"Level for domain"`
	Class       *dataClass.ClassPublicResponse             `json:"class" required:"false" doc:"Class"`
}

type StudentResponseList struct {
	types.PaginatedResponse
	Data []StudentResponse `json:"data" required:"false" doc:"List of students" example:"[]"`
}

type StudentEnrollResponseList struct {
	types.PaginatedResponse
	Data []StudentEnrollResponse `json:"data" required:"false" doc:"List of level domain/class for the specified student" example:"[]"`
}
