package data

type TeacherID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teacher id" example:"1"`
}

type TUSubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teaching unit/subject id" example:"1"`
}

type TeacherRequest struct {
	SchoolID int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	UserID   int64  `json:"userID" required:"true" doc:"User id" example:"1"`
	UID      string `json:"uid" required:"false" doc:"Teacher UID" example:"1"`
}

type TeacherTUSubjectRequest struct {
	TeacherID int64 `json:"teacherID" required:"true" doc:"Teacher id" example:"1"`
	YearID    int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`

	TeachingUnitID int64 `json:"teachingUnitID" required:"false" doc:"Teaching unit id" example:"1"`
	SubjectID      int64 `json:"subjectID" required:"false" doc:"Subject id" example:"1"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllTeacherTUSubjectRequest struct {
	TeacherID int64 `json:"teacherID" path:"id" required:"true" doc:"Teacher id" example:"1"`
}
