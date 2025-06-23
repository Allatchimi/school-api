package data

type MeetingRoomID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Meeting room id"`
}

type MeetingRoomRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Subject class id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`
}

type GetAllRequest struct {
	SchoolID      int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	LevelDomainID int64 `json:"levelDomainID" query:"levelDomainID" required:"false" doc:"Level domain id"`
	ClassID       int64 `json:"classID" query:"classID" required:"false" doc:"Class id"`
}
