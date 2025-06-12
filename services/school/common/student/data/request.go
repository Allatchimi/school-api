package data

import (
	dataUser "api/services/user/user/data"
	"time"
)

type StudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Student id" example:"1"`
}

type StudentEnrollID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Unit/subject id" example:"1"`
}

type StudentRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`

	UID         string                    `json:"uid" required:"false" doc:"Teacher UID" example:"1"`
	Email       string                    `json:"email" required:"true" minLength:"3" max:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64                    `json:"phoneNumber" required:"false" doc:"Phone number" example:"237690909090"`
	Info        *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information" example:""`
}

type StudentEnrollRequest struct {
	StudentID     int64 `json:"teacherID" required:"false" doc:"Student id" example:"1"`
	YearID        int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`
	ClassID       int64 `json:"classID" required:"false" doc:"Class id" example:"1"`
	LevelDomainID int64 `json:"levelDomainID" required:"false" doc:"Level domain id" example:"1"`

	Email       string `json:"email" required:"false" minLength:"3" max:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number" example:"237690909090"`

	Message       string     `json:"message" required:"false" doc:"Message" example:""`
	Gender        string     `json:"gender" required:"true" enum:"m,f" doc:"Gender" example:"M"`
	FirstName     string     `json:"firstName" required:"true" max:"30" doc:"First name" example:"John"`
	LastName      string     `json:"lastName" required:"true" max:"30" doc:"Last name" example:"Doe"`
	Birthday      *time.Time `json:"birthday" required:"true" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"true" doc:"Birth location"`

	Document1 string `json:"file1" required:"false" doc:"Document1" example:""`
	Document2 string `json:"file2" required:"false" doc:"Document2" example:""`
	Document3 string `json:"file3" required:"false" doc:"Document3" example:""`
	Document4 string `json:"file4" required:"false" doc:"Document4" example:""`
	Document5 string `json:"file5" required:"false" doc:"Document5" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllStudentEnrollRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	StudentID int64 `json:"teacherID" query:"teacherID" required:"false" doc:"Student id" example:"1"`
}
