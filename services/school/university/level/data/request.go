package data

type LevelID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Level id" example:"1"`
}

type LevelDomainID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Level domain id" example:"1"`
}

type LevelRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`

	Name        string `json:"name" required:"true" doc:"Level name" example:"Sciences"`
	Description string `json:"description" required:"false" doc:"Level description" example:""`
}

type LevelDomainRequest struct {
	LevelID  int64 `json:"levelID" required:"true" doc:"Level id" example:"1"`
	DomainID int64 `json:"domainID" required:"true" doc:"Domain id" example:"1"`

	Program      string `json:"program" required:"true" doc:"Program" example:""`
	Requirements string `json:"requirements" required:"false" doc:"Requirements" example:""`
	IsValid      bool   `json:"isValid" required:"false" doc:"Is valid" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllLevelDomainRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	LevelID  int64 `json:"levelID" query:"levelID" required:"false" doc:"Level id" example:"1"`
	DomainID int64 `json:"domainID" query:"domainID" required:"false" doc:"Domain id" example:"1"`
}
