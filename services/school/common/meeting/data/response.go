package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataClass "api/services/school/highschool/class/data"
	dataUnit "api/services/school/university/unit/data"
)

type MeetingRoomResponse struct {
	types.BaseGormModelResponse
	School       *dataSchool.SchoolPublicResponse      `json:"school" required:"false" doc:"School"`
	Unit         *dataUnit.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`
	ClassSubject *dataClass.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`

	ApiRoomID string `json:"apiRoomID" required:"false" doc:"Room id for the API"`
}

type MeetingRoomPublicResponse struct {
	School       *dataSchool.SchoolPublicResponse      `json:"school" required:"false" doc:"School"`
	Unit         *dataUnit.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`
	ClassSubject *dataClass.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Subject for specific class"`

	ApiRoomID string `json:"apiRoomID" required:"false" doc:"Room id for the API"`
}

type MeetingRoomResponseList struct {
	types.PaginatedResponse
	Data []MeetingRoomResponse `json:"data" required:"false" doc:"List of rooms" example:"[]"`
}
