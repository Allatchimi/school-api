package data

import "time"

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
	SchoolID        int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	StudentEnrollID int64 `json:"studentEnrollID" query:"studentEnrollID" required:"false" doc:"Student enroll id"`
	StudentID       int64 `json:"studentID" query:"studentID" required:"false" doc:"Student id"`
}
