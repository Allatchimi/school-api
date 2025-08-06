package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
)

type TeacherID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teacher id"`
}

type UnitSubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Unit/subject id"`
}

type TeacherRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	UID               string                    `json:"uid" required:"false" doc:"User UID"`
	AutoGenerateEmail bool                      `json:"autoGenerateEmail" required:"false" doc:"Auto generate email"`
	Email             string                    `json:"email" required:"false" format:"email" doc:"Email"`
	PhoneNumber       uint64                    `json:"phoneNumber" required:"false" doc:"Phone number"`
	Status            string                    `json:"status" required:"true" enum:"enabled,disabled" doc:"Status"`
	Info              *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type TeacherClassSubjectUnitRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Subject class id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`
	TeacherID      int64 `json:"teacherID" required:"true" doc:"Teacher id"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllTeacherClassSubjectUnitRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentParentRequest
	ClassID       int64 `json:"classID" query:"classID" required:"false" doc:"Class id"`
	LevelDomainID int64 `json:"levelDomainID" query:"levelDomainID" required:"false" doc:"Level domain id"`
}
