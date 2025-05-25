package data

type MeetingRoomID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Meeting room id" example:"1"`
}

type MeetingRoomRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`

	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Subject class id" example:"1"`
}
