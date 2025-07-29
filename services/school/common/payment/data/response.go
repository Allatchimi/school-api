package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	"time"
)

type PaymentResponse struct {
	types.BaseGormModelResponse
	School        *dataSchool.SchoolPublicResponse   `json:"school" required:"false" doc:"School"`
	StudentEnroll *dataStudent.StudentEnrollResponse `json:"studentEnroll" required:"false" doc:"Student Enroll"`

	Amount   float64    `json:"amount" required:"false" doc:"Amount"`
	Currency string     `json:"currency" required:"false" doc:"Currency"`
	Date     *time.Time `json:"date" required:"false" doc:"Date"`
	Method   string     `json:"method" required:"false" doc:"Method"`
	Status   string     `json:"status" required:"false" doc:"Status"`
	Message  string     `json:"message" required:"false" doc:"Message"`
}

type PaymentResponseList struct {
	types.PaginatedResponse
	Data []PaymentResponse `json:"data" required:"false" doc:"List of payment"`
}
