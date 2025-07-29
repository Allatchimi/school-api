package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataLevel "api/services/school/university/level/data"
	dataUser "api/services/user/user/data"
	"time"
)

type StudentResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserResponse           `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Student UID"`
}

type StudentPublicResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Student UID"`
}

type StudentEnrollResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Year        *dataYear.YearResponse           `json:"year" required:"false" doc:"Year"`
	Class       *dataClass.ClassResponse         `json:"class" required:"false" doc:"Class"`
	LevelDomain *dataLevel.LevelDomainResponse   `json:"levelDomain" required:"false" doc:"Level for domain"`
	Student     *StudentPublicResponse           `json:"student" required:"false" doc:"Student"`

	Origin         string `json:"origin" required:"false" doc:"Origin"`
	OriginFeedback string `json:"originFeedback" required:"false" doc:"Origin feedback"`
}

type StudentPreEnrollResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Year        *dataYear.YearResponse           `json:"year" required:"false" doc:"Year"`
	Class       *dataClass.ClassResponse         `json:"class" required:"false" doc:"Class"`
	LevelDomain *dataLevel.LevelDomainResponse   `json:"levelDomain" required:"false" doc:"Level for domain"`
	User        *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	Birthday       *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation  string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Status         string     `json:"status" required:"false" doc:"Status"`
	StatusFeedback string     `json:"statusFeedback" required:"false" doc:"Status feedback"`
	Message        string     `json:"Message" required:"false" doc:"Message"`
	Gender         string     `json:"gender" required:"false" doc:"Gender"`
	FirstName      string     `json:"firstName" required:"false" doc:"First name"`
	LastName       string     `json:"lastName" required:"false" doc:"Last name or family name"`
	Document1      string     `json:"document1" required:"false" doc:"Document1"`
	Document2      string     `json:"document2" required:"false" doc:"Document2"`
	Document3      string     `json:"document3" required:"false" doc:"Document3"`
	Document4      string     `json:"document4" required:"false" doc:"Document4"`
	Document5      string     `json:"document5" required:"false" doc:"Document5"`
}

type StudentResponseList struct {
	types.PaginatedResponse
	Data []StudentResponse `json:"data" required:"false" doc:"List of student" example:"[]"`
}

type StudentEnrollResponseList struct {
	types.PaginatedResponse
	Data []StudentEnrollResponse `json:"data" required:"false" doc:"List of student enroll" example:"[]"`
}

type StudentPreEnrollResponseList struct {
	types.PaginatedResponse
	Data []StudentPreEnrollResponse `json:"data" required:"false" doc:"List of student pre enroll" example:"[]"`
}
