package data

import "time"

type PaymentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Academic payment id" example:"1"`
}

type PaymentRequest struct {
	SchoolID        int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	StudentEnrollID int64 `json:"studentEnrollID" required:"true" doc:"Student enroll id" example:"1"`

	Amount        float64    `json:"amount" required:"true" doc:"Amount" example:"1000"`
	Currency      string     `json:"currency" required:"true" doc:"Currency" example:""`
	PaymentDate   *time.Time `json:"paymentDate" required:"false" doc:"Payment date" example:""`
	PaymentMethod string     `json:"paymentMethod" required:"true" doc:"Payment method" example:""`
	PaymentStatus string     `json:"paymentStatus" required:"true" enum:"pending,success,failed,canceled,rejected" doc:"Payment status" example:"pending"`
	PaymentNote   string     `json:"paymentNote" required:"false" doc:"Payment note" example:""`
}

type GetAllRequest struct {
	SchoolID        int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	StudentEnrollID int64 `json:"studentEnrollID" query:"studentEnrollID" required:"false" doc:"Student enroll id" example:"1"`
}
