package data

type ParentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent id" example:"1"`
}

type ParentStudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Parent level class id" example:"1"`
}

type ParentRequest struct {
	UserID int64 `json:"userID" required:"true" doc:"User id" example:"1"`
}

type ParentStudentRequest struct {
	ParentID  int64 `json:"parentID" required:"true" doc:"Parent id" example:"1"`
	StudentID int64 `json:"studentID" required:"true" doc:"Student id" example:"1"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
