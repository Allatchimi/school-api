package data

import (
	"api/common/types"
	"time"
)

type SchoolResponse struct {
	types.BaseGormModelResponse
	Name string `json:"name" doc:"School name"`
	Type string `json:"type" doc:"Type"`
	Logo string `json:"logo" doc:"School logo"`

	Info   *SchoolInfoResponse   `json:"info" doc:"Information"`
	Config *SchoolConfigResponse `json:"config" doc:"Configuration"`
}

type SchoolPublicResponse struct {
	Name string `json:"name" doc:"School name"`
	Type string `json:"type" doc:"Type"`
	Logo string `json:"logo" doc:"School logo"`

	Info *SchoolInfoResponse `json:"info" doc:"Information"`
}

type SchoolInfoResponse struct {
	FullName    string `json:"fullName" doc:"School name"`
	Description string `json:"description" doc:"Description"`
	Slogan      string `json:"slogan" doc:"Slogan"`

	PhoneNumber1 int64 `json:"phoneNumber1" doc:"Phone number 1"`
	PhoneNumber2 int64 `json:"phoneNumber2" doc:"Phone number 2"`
	PhoneNumber3 int64 `json:"phoneNumber3" doc:"Phone number 3"`

	Email1 string `json:"email1" doc:"Email 1"`
	Email2 string `json:"email2" doc:"Email 2"`
	Email3 string `json:"email3" doc:"Email 3"`

	Founder   string     `json:"founder" doc:"Founder name"`
	FoundedAt *time.Time `json:"foundedAt" doc:"Founded date time"`

	Address           string  `json:"address" doc:"Address"`
	LocationLongitude float64 `json:"locationLongitude" doc:"Location longitude"`
	LocationLatitude  float64 `json:"locationLatitude" doc:"Location latitude"`

	Image1 string `json:"image1" doc:"Image 1"`
	Image2 string `json:"image2" doc:"Image 2"`
	Image3 string `json:"image3" doc:"Image 3"`
	Image4 string `json:"image4" doc:"Image 4"`
}

type SchoolConfigResponse struct {
	EmailDomain  string `json:"emailDomain" doc:"Email domain"`
	ColorPrimary string `json:"colorPrimary" doc:"Color primary"`
}

type SchoolResponseList struct {
	types.PaginatedResponse
	Data []SchoolResponse `json:"data" required:"false" doc:"List of schools" example:"[]"`
}
