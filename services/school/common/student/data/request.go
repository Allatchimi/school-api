package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
	"time"
)

type StudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Student id"`
}

type StudentEnrollID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Student enroll id"`
}

type StudentPreEnrollID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Student pre enroll id"`
}

type StudentRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	UID               string                    `json:"uid" required:"false" doc:"User UID"`
	AutoGenerateEmail bool                      `json:"autoGenerateEmail" required:"false" doc:"Auto generate email"`
	Email             string                    `json:"email" required:"false" format:"email" doc:"Email"`
	PhoneNumber       uint64                    `json:"phoneNumber" required:"false" doc:"Phone number"`
	Status            string                    `json:"status" required:"true" enum:"enabled,disabled" doc:"Status"`
	Info              *dataUser.UserInfoRequest `json:"info" required:"true" doc:"Information"`
}

type StudentEnrollRequest struct {
	SchoolID      int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID        int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassID       int64 `json:"classID" required:"false" doc:"Class id"`
	LevelDomainID int64 `json:"levelDomainID" required:"false" doc:"Level domain id"`
	StudentID     int64 `json:"studentID" required:"true" doc:"Student id"`

	Origin         string `json:"origin" required:"false" doc:"Origin"`
	OriginFeedback string `json:"originFeedback" required:"false" doc:"Origin feedback"`
}

type StudentPreEnrollRequest struct {
	SchoolID      int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID        int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassID       int64 `json:"classID" required:"false" doc:"Class id"`
	LevelDomainID int64 `json:"levelDomainID" required:"false" doc:"Level domain id"`

	Message       string     `json:"message" required:"false" doc:"Message"`
	Gender        string     `json:"gender" required:"true" enum:"male,female" doc:"Gender"`
	FirstName     string     `json:"firstName" required:"true" doc:"First name"`
	LastName      string     `json:"lastName" required:"true" doc:"Last name"`
	Birthday      *time.Time `json:"birthday" required:"true" doc:"Birthday"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Document1     string     `json:"document1" required:"false" doc:"Document1"`
	Document2     string     `json:"document2" required:"false" doc:"Document2"`
	Document3     string     `json:"document3" required:"false" doc:"Document3"`
	Document4     string     `json:"document4" required:"false" doc:"Document4"`
	Document5     string     `json:"document5" required:"false" doc:"Document5"`
}

type StudentPreEnrollStatusRequest struct {
	Status         string `json:"status" required:"true" enum:"initiated,pending,enrolled,rejected" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"false" doc:"Status feedback"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	ExamID   int64 `json:"examID" query:"examID" required:"false" doc:"Exam id"`
}

type GetAllStudentEnrollRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentParentRequest
	ClassID       int64 `json:"classID" query:"classID" required:"false" doc:"Class id"`
	LevelDomainID int64 `json:"levelDomainID" query:"levelDomainID" required:"false" doc:"Level domain id"`
}

type GetAllStudentPreEnrollRequest struct {
	types.FilterSchoolYearClassLevelDomainRequest
	UserID int64 `json:"userID" query:"userID" required:"false" doc:"User id"`
}
