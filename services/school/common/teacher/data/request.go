package data

import dataUser "api/services/user/user/data"

type TeacherID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teacher id" example:"1"`
}

type UnitSubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Unit/subject id" example:"1"`
}

type TeacherRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`

	UID         string                    `json:"uid" required:"false" doc:"Teacher UID" example:"1"`
	Email       string                    `json:"email" required:"true" minLength:"3" max:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64                    `json:"phoneNumber" required:"false" doc:"Phone number" example:"237690909090"`
	Info        *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information" example:""`
}

type TeacherClassSubjectUnitRequest struct {
	TeacherID int64 `json:"teacherID" required:"true" doc:"Teacher id" example:"1"`
	YearID    int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`

	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Subject class id" example:"1"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id" example:"1"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllTeacherClassSubjectUnitRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	TeacherID int64 `json:"teacherID" query:"teacherID" required:"false" doc:"Teacher id" example:"1"`
}
