package types

type Filter struct {
	Search  string `json:"search" query:"search" required:"false" doc:"Search keyword"`
	OrderBy string `json:"orderBy" query:"orderBy" required:"false" doc:"Filter by field name(case sensitive)"`
	Sort    string `json:"sort" query:"sort" required:"false" enum:"asc,desc" doc:"Sort asc or desc"`
}

type FilterSchoolYearClassSubjectUnitRequest struct {
	SchoolID       int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	YearID         int64 `json:"yearID" query:"yearID" required:"false" doc:"Year id"`
	ClassSubjectID int64 `json:"classID" query:"classSubjectID" required:"false" doc:"Class subject id"`
	UnitID         int64 `json:"unitID" query:"unitID" required:"false" doc:"Unit id"`
}
