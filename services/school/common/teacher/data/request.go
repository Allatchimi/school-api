package data

type TeacherID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teacher id" example:"1"`
}

type TeacherLevelClassID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teacher level class id" example:"1"`
}

type TeacherRequest struct {
	SchoolID int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	UserID   int64  `json:"userID" required:"true" doc:"User id" example:"1"`
	UID      string `json:"uid" required:"false" doc:"Teacher UID" example:"1"`
}

type TeacherLevelClassRequest struct {
	TeacherID int64 `json:"teacherID" required:"true" doc:"Teacher id" example:"1"`
	YearID    int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`

	DomainID int64 `json:"domainID" required:"false" doc:"Domain id" example:"1"`
	LevelID  int64 `json:"levelID" required:"false" doc:"Level id" example:"1"`

	ClassID int64 `json:"classID" required:"false" doc:"Class id" example:"1"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllLevelClassRequest struct {
	TeacherID int64 `json:"teacherID" query:"teacherID" required:"false" doc:"Teacher id" example:"1"`
}
