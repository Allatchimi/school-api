package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataLevelDomain "api/services/school/university/level/data"
	dataUser "api/services/user/user/data"
	"time"
)

type StudentResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Student UID"`
}

type StudentPublicResponse struct {
	School *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	User   *dataUser.UserPublicResponse     `json:"user" required:"false" doc:"User"`

	UID string `json:"uid" required:"false" doc:"Student UID"`
}

type StudentEnrollResponse struct {
	types.BaseGormModelResponse
	Student     *StudentPublicResponse                     `json:"student" required:"false" doc:"Student"`
	Year        *dataYear.YearPublicResponse               `json:"year" required:"false" doc:"Year"`
	LevelDomain *dataLevelDomain.LevelDomainPublicResponse `json:"levelDomain" required:"false" doc:"Level for domain"`
	Class       *dataClass.ClassPublicResponse             `json:"class" required:"false" doc:"Class"`

	Origin          string  `json:"origin" required:"false" doc:"Origin"`
	IsAccepted      bool    `json:"isAccepted" required:"false" doc:"Is accepted"`
	Payment         float64 `json:"payment" required:"false" doc:"Payment"`
	PaymentCurrency string  `json:"paymentCurrency" required:"false" doc:"Payment currency"`

	Email       string `json:"email" required:"false" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`

	Gender        string     `json:"gender" required:"false" doc:"Gender"`
	FirstName     string     `json:"firstName" required:"false" doc:"First name"`
	LastName      string     `json:"lastName" required:"false" doc:"Last name or family name"`
	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`

	Document1 string `json:"file1" required:"false" doc:"Document1"`
	Document2 string `json:"file2" required:"false" doc:"Document2"`
	Document3 string `json:"file3" required:"false" doc:"Document3"`
	Document4 string `json:"file4" required:"false" doc:"Document4"`
	Document5 string `json:"file5" required:"false" doc:"Document5"`
}

type StudentEnrollPublicResponse struct {
	Student     *StudentPublicResponse                     `json:"student" required:"false" doc:"Student"`
	Year        *dataYear.YearPublicResponse               `json:"year" required:"false" doc:"Year"`
	LevelDomain *dataLevelDomain.LevelDomainPublicResponse `json:"levelDomain" required:"false" doc:"Level for domain"`
	Class       *dataClass.ClassPublicResponse             `json:"class" required:"false" doc:"Class"`

	Origin          string  `json:"origin" required:"false" doc:"Origin"`
	IsAccepted      bool    `json:"isAccepted" required:"false" doc:"Is accepted"`
	Payment         float64 `json:"payment" required:"false" doc:"Payment"`
	PaymentCurrency string  `json:"paymentCurrency" required:"false" doc:"Payment currency"`

	Email       string `json:"email" required:"false" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`

	Gender        string     `json:"gender" required:"false" doc:"Gender"`
	FirstName     string     `json:"firstName" required:"false" doc:"First name"`
	LastName      string     `json:"lastName" required:"false" doc:"Last name or family name"`
	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`

	Document1 string `json:"file1" required:"false" doc:"Document1"`
	Document2 string `json:"file2" required:"false" doc:"Document2"`
	Document3 string `json:"file3" required:"false" doc:"Document3"`
	Document4 string `json:"file4" required:"false" doc:"Document4"`
	Document5 string `json:"file5" required:"false" doc:"Document5"`
}

type StudentResponseList struct {
	types.PaginatedResponse
	Data []StudentResponse `json:"data" required:"false" doc:"List of students" example:"[]"`
}

type StudentEnrollResponseList struct {
	types.PaginatedResponse
	Data []StudentEnrollResponse `json:"data" required:"false" doc:"List of level domain/class for the specified student" example:"[]"`
}
