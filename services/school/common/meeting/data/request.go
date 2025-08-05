package data

import "api/common/types"

type MeetingRoomID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Meeting room id"`
}

type MeetingRoomRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Subject class id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassLevelDomainRequest
	types.FilterTeacherStudentParentRequest
}
