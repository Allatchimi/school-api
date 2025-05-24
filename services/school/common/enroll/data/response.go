package data

import (
	"api/common/types"
	"time"
)

type SchoolResponse struct {
	types.BaseGormModelResponse
	Name     string `json:"name" required:"false" doc:"School name"`
	Type     string `json:"type" required:"false" doc:"Type"`
	Logo     string `json:"logo" required:"false" doc:"School logo"`
	Currency string `json:"currency" required:"false" doc:"Currency"`

	Info   *SchoolInfoResponse   `json:"info" required:"false" doc:"Information"`
	Config *SchoolConfigResponse `json:"config" required:"false" doc:"Configuration"`
}

type SchoolPublicResponse struct {
	Name     string `json:"name" required:"false" doc:"School name"`
	Type     string `json:"type" required:"false" doc:"Type"`
	Logo     string `json:"logo" required:"false" doc:"School logo"`
	Currency string `json:"currency" required:"false" doc:"Currency"`

	Info *SchoolInfoResponse `json:"info" required:"false" doc:"Information"`
}

type SchoolInfoResponse struct {
	FullName    string `json:"fullName" required:"false" doc:"School name"`
	Description string `json:"description" required:"false" doc:"Description"`
	Slogan      string `json:"slogan" required:"false" doc:"Slogan"`

	PhoneNumber1 int64 `json:"phoneNumber1" required:"false" doc:"Phone number 1"`
	PhoneNumber2 int64 `json:"phoneNumber2" required:"false" doc:"Phone number 2"`
	PhoneNumber3 int64 `json:"phoneNumber3" required:"false" doc:"Phone number 3"`

	Email1 string `json:"email1" required:"false" doc:"Email 1"`
	Email2 string `json:"email2" required:"false" doc:"Email 2"`
	Email3 string `json:"email3" required:"false" doc:"Email 3"`

	Founder   string     `json:"founder" required:"false" doc:"Founder name"`
	FoundedAt *time.Time `json:"foundedAt" required:"false" doc:"Founded date time"`

	Address           string  `json:"address" required:"false" doc:"Address"`
	PoBox             string  `json:"poBox" required:"false" doc:"PO Box"`
	LocationLongitude float64 `json:"locationLongitude" required:"false" doc:"Location longitude"`
	LocationLatitude  float64 `json:"locationLatitude" required:"false" doc:"Location latitude"`

	Image1 string `json:"image1" required:"false" doc:"Image 1"`
	Image2 string `json:"image2" required:"false" doc:"Image 2"`
	Image3 string `json:"image3" required:"false" doc:"Image 3"`
	Image4 string `json:"image4" required:"false" doc:"Image 4"`
}

type SchoolConfigResponse struct {
	EmailDomain  string `json:"emailDomain" required:"false" doc:"Email domain"`
	ColorPrimary string `json:"colorPrimary" required:"false" doc:"Color primary"`
}

type SchoolResponseList struct {
	types.PaginatedResponse
	Data []SchoolResponse `json:"data" required:"false" doc:"List of schools" example:"[]"`
}
