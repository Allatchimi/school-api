package data

type UnitID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Unit id"`
}

type UnitRequest struct {
	SchoolID      int64 `json:"schoolID" required:"true" doc:"School id"`
	LevelDomainID int64 `json:"levelDomainID" required:"true" doc:"Level domain id"`
	SemesterID    int64 `json:"semesterID" required:"true" doc:"Semester id"`

	Name         string `json:"name" required:"true" doc:"Name"`
	Description  string `json:"description" required:"false" doc:"Description"`
	Credit       int    `json:"credit" required:"true" doc:"Credit"`
	Program      string `json:"program" required:"false" doc:"Program"`
	Requirements string `json:"requirements" required:"false" doc:"Requirements"`
	IsValid      bool   `json:"isValid" required:"false" doc:"Is valid"`
}

type GetAllRequest struct {
	SchoolID      int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	LevelDomainID int64 `json:"levelDomainID" query:"levelDomainID" required:"false" doc:"Level domain id"`
	SemesterID    int64 `json:"semesterID" query:"semesterID" required:"false" doc:"Semester id"`
}
