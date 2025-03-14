package data

type SubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Subject id" example:"1"`
}

type CreateSubjectRequest struct {
	SchoolID     int64  `json:"schoolID" required:"false" doc:"School id" example:"1"`
	ClassID      int64  `json:"classID" required:"false" doc:"Class id" example:"1"`
	Name         string `json:"name" required:"false" doc:"Name" example:"Biologic"`
	Description  string `json:"description" required:"false" doc:"Description" example:""`
	Coefficient  int    `json:"Coefficient" required:"false" doc:"Coefficient" example:""`
	Program      string `json:"Program" required:"false" doc:"Program" example:""`
	Requirements string `json:"Requirements" required:"false" doc:"Requirements" example:""`
}

type SubjectProfessorRequest struct {
	UserID int64 `json:"userID" required:"true" doc:"User id" example:"1"`
}

type UpdateSubjectRequest struct {
	Name         string `json:"name" required:"false" doc:"Name" example:"Biologic"`
	Description  string `json:"description" required:"false" doc:"Description" example:""`
	Coefficient  int    `json:"Coefficient" required:"false" doc:"Coefficient" example:""`
	Program      string `json:"Program" required:"false" doc:"Program" example:""`
	Requirements string `json:"Requirements" required:"false" doc:"Requirements" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
