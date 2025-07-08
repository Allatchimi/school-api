package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataSpecialty "api/services/school/highschool/specialty/data"
	dataSubject "api/services/school/highschool/subject/data"
	"time"
)

type ClassResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Specialty   *dataSpecialty.SpecialtyResponse `json:"specialty" required:"false" doc:"Specialty"`
	Fees        int64                            `json:"fees" required:"false" doc:"Class fees"`
	Name        string                           `json:"name" required:"false" doc:"Class name"`
	Description string                           `json:"description" required:"false" doc:"Class description"`
}

type ClassSubjectResponse struct {
	types.BaseGormModelResponse
	School  *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Subject *dataSubject.SubjectResponse     `json:"subject" required:"false" doc:"Subject"`
	Class   *ClassResponse                   `json:"class" required:"false" doc:"Class"`

	Coefficient  int    `json:"coefficient" required:"false" doc:"Coefficient"`
	Program      string `json:"program" required:"false" doc:"Program"`
	Requirements string `json:"requirements" required:"false" doc:"Requirements"`

	IsValid     bool       `json:"isValid" required:"false" doc:"Is valid"`
	InvalidDate *time.Time `json:"invalidDate" required:"false" doc:"Invalid date"`
}

type ClassResponseList struct {
	types.PaginatedResponse
	Data []ClassResponse `json:"data" required:"false" doc:"List of classes" example:"[]"`
}

type ClassSubjectResponseList struct {
	types.PaginatedResponse
	Data []ClassSubjectResponse `json:"data" required:"false" doc:"List of subject for matching class" example:"[]"`
}
