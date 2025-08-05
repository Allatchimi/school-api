package data

import (
	"api/common/types"
	"time"
)

type PaymentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Academic payment id"`
}

type PaymentRequest struct {
	SchoolID        int64 `json:"schoolID" required:"true" doc:"School id"`
	StudentEnrollID int64 `json:"studentEnrollID" required:"true" doc:"Student enroll id"`

	Amount   float64    `json:"amount" required:"true" doc:"Amount"`
	Currency string     `json:"currency" required:"true" doc:"Currency"`
	Date     *time.Time `json:"date" required:"false" doc:"Date"`
	Method   string     `json:"method" required:"true" doc:"Method"`
	Status   string     `json:"status" required:"true" enum:"pending,success,failed,canceled,rejected" doc:"Status"`
	Message  string     `json:"message" required:"false" doc:"Message"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassLevelDomainRequest
	types.FilterTeacherStudentParentRequest
	StudentEnrollID int64 `json:"studentEnrollID" query:"studentEnrollID" required:"false" doc:"Student enroll id"`
}
