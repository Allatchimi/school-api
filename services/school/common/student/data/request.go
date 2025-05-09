package data

type StudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Student id" example:"1"`
}

type StudentEnrollID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Unit/subject id" example:"1"`
}

type StudentRequest struct {
	SchoolID int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	UserID   int64  `json:"userID" required:"true" doc:"User id" example:"1"`
	UID      string `json:"uid" required:"false" doc:"Student UID" example:"1"`
}

type StudentEnrollRequest struct {
	StudentID int64 `json:"teacherID" required:"true" doc:"Student id" example:"1"`
	YearID    int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`

	LevelDomainID int64 `json:"levelDomainID" required:"false" doc:"Level domain id" example:"1"`
	ClassID       int64 `json:"classID" required:"false" doc:"Class id" example:"1"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllStudentEnrollRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	StudentID int64 `json:"teacherID" query:"teacherID" required:"false" doc:"Student id" example:"1"`
}
