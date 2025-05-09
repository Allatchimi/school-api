package data

type TeacherID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teacher id" example:"1"`
}

type UnitSubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Unit/subject id" example:"1"`
}

type TeacherRequest struct {
	SchoolID int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	UserID   int64  `json:"userID" required:"true" doc:"User id" example:"1"`
	UID      string `json:"uid" required:"false" doc:"Teacher UID" example:"1"`
}

type TeacherUnitSubjectRequest struct {
	TeacherID int64 `json:"teacherID" required:"true" doc:"Teacher id" example:"1"`
	YearID    int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`

	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Subject class id" example:"1"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllTeacherUnitSubjectRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	TeacherID int64 `json:"teacherID" query:"teacherID" required:"false" doc:"Teacher id" example:"1"`
}
