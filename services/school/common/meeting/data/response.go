package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataClass "api/services/school/highschool/class/data"
	dataLevelDomain "api/services/school/university/level/data"
)

type MeetingRoomResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse           `json:"school" required:"false" doc:"School"`
	LevelDomain *dataLevelDomain.LevelDomainPublicResponse `json:"levelDomain" required:"false" doc:"Level for domain"`
	Class       *dataClass.ClassPublicResponse             `json:"class" required:"false" doc:"Class"`

	ApiRoomID string `json:"apiRoomID" required:"false" doc:"Room id for the API"`
}

type MeetingRoomPublicResponse struct {
	School      *dataSchool.SchoolPublicResponse           `json:"school" required:"false" doc:"School"`
	LevelDomain *dataLevelDomain.LevelDomainPublicResponse `json:"levelDomain" required:"false" doc:"Level for domain"`
	Class       *dataClass.ClassPublicResponse             `json:"class" required:"false" doc:"Class"`

	ApiRoomID string `json:"apiRoomID" required:"false" doc:"Room id for the API"`
}

type MeetingRoomResponseList struct {
	types.PaginatedResponse
	Data []MeetingRoomResponse `json:"data" required:"false" doc:"List of rooms" example:"[]"`
}
