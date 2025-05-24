package data

type MeetingRoomID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Meeting room id" example:"1"`
}

type MeetingRoomRequest struct {
	SchoolID      int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	LevelDomainID int64 `json:"levelDomainID" required:"false" doc:"Level domain id" example:"1"`
	ClassID       int64 `json:"classID" required:"false" doc:"Class id" example:"1"`
}
