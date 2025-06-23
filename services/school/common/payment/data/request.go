package data

import "time"

type PaymentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Academic payment id"`
}

type PaymentRequest struct {
	SchoolID        int64 `json:"schoolID" required:"true" doc:"School id"`
	StudentEnrollID int64 `json:"studentEnrollID" required:"true" doc:"Student enroll id"`

	Amount        float64    `json:"amount" required:"true" doc:"Amount"`
	Currency      string     `json:"currency" required:"true" doc:"Currency"`
	PaymentDate   *time.Time `json:"paymentDate" required:"false" doc:"Payment date"`
	PaymentMethod string     `json:"paymentMethod" required:"true" doc:"Payment method"`
	PaymentStatus string     `json:"paymentStatus" required:"true" enum:"pending,success,failed,canceled,rejected" doc:"Payment status"`
	PaymentNote   string     `json:"paymentNote" required:"false" doc:"Payment note"`
}

type GetAllRequest struct {
	SchoolID        int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	StudentEnrollID int64 `json:"studentEnrollID" query:"studentEnrollID" required:"false" doc:"Student enroll id"`
}
