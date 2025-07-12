package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataStudent "api/services/school/common/student/data"
	"time"
)

type PaymentResponse struct {
	types.BaseGormModelResponse
	Amount        float64    `json:"amount" required:"false" doc:"Amount"`
	Currency      string     `json:"currency" required:"false" doc:"Currency"`
	PaymentDate   *time.Time `json:"paymentDate" required:"false" doc:"Payment date"`
	PaymentMethod string     `json:"paymentMethod" required:"false" doc:"Payment method"`
	PaymentStatus string     `json:"paymentStatus" required:"false" doc:"Payment status"`
	PaymentNote   string     `json:"paymentNote" required:"false" doc:"Payment note"`

	School        *dataSchool.SchoolPublicResponse         `json:"school" required:"false" doc:"School"`
	StudentEnroll *dataStudent.StudentEnrollPublicResponse `json:"studentEnroll" required:"false" doc:"Student Enroll"`
}

type PaymentResponseList struct {
	types.PaginatedResponse
	Data []PaymentResponse `json:"data" required:"false" doc:"List of academic Payments"`
}
