package data

type LevelID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Level id"`
}

type LevelDomainID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Level domain id"`
}

type LevelRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	Name        string `json:"name" required:"true" doc:"Level name"`
	Description string `json:"description" required:"false" doc:"Level description"`
}

type LevelDomainRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`
	LevelID  int64 `json:"levelID" required:"true" doc:"Level id"`
	DomainID int64 `json:"domainID" required:"true" doc:"Domain id"`

	Program      string `json:"program" required:"true" doc:"Program"`
	Requirements string `json:"requirements" required:"false" doc:"Requirements"`
	IsValid      bool   `json:"isValid" required:"false" doc:"Is valid"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllLevelDomainRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	LevelID  int64 `json:"levelID" query:"levelID" required:"false" doc:"Level id"`
	DomainID int64 `json:"domainID" query:"domainID" required:"false" doc:"Domain id"`
}
