package data

type UnitID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Unit id" example:"1"`
}

type UnitRequest struct {
	SchoolID      int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	LevelDomainID int64 `json:"levelDomainID" required:"true" doc:"Level domain id" example:"1"`
	SemesterID    int64 `json:"semesterID" required:"true" doc:"Semester id" example:"1"`

	Name         string `json:"name" required:"true" doc:"Name" example:"MATH110"`
	Description  string `json:"description" required:"false" doc:"Description" example:""`
	Credit       int    `json:"credit" required:"true" min:"1" doc:"Credit" example:"1"`
	Program      string `json:"program" required:"false" doc:"Program" example:""`
	Requirements string `json:"requirements" required:"false" doc:"Requirements" example:""`

	IsValid bool `json:"isValid" required:"false" doc:"Is valid" example:"true"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
