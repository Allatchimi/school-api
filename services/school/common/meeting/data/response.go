package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataClass "api/services/school/highschool/class/data"
	dataUnit "api/services/school/university/unit/data"
)

type MeetingRoomResponse struct {
	types.BaseGormModelResponse
	School       *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	ClassSubject *dataClass.ClassSubjectResponse  `json:"classSubject" required:"false" doc:"Subject for specific class"`
	Unit         *dataUnit.UnitResponse           `json:"unit" required:"false" doc:"Unit"`

	ApiRoomID string `json:"apiRoomID" required:"false" doc:"Room id for the API"`
}

type MeetingRoomResponseList struct {
	types.PaginatedResponse
	Data []MeetingRoomResponse `json:"data" required:"false" doc:"List of rooms"`
}
