package model

import (
	"api/common/types"
	"api/services/school/common/payment/data"
	modelSchool "api/services/school/common/school/model"
	modelStudent "api/services/school/common/student/model"
	"time"
)

type Payment struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentEnrollID int64                       `gorm:"default:null"`
	StudentEnroll   *modelStudent.StudentEnroll `gorm:"default:null;foreignKey:StudentEnrollID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Amount        float64    `gorm:"default:null"`
	Currency      string     `gorm:"default:null"`
	PaymentDate   *time.Time `gorm:"default:null"`
	PaymentMethod string     `gorm:"default:null"`
	PaymentStatus string     `gorm:"default:null"`
	PaymentNote   string     `gorm:"default:null;type:text"`
}

func (item *Payment) ToResponse() *data.PaymentResponse {
	if item == nil {
		return nil
	}
	resp := &data.PaymentResponse{}
	resp.Amount = item.Amount
	resp.Currency = item.Currency
	resp.PaymentDate = item.PaymentDate
	resp.PaymentMethod = item.PaymentMethod
	resp.PaymentStatus = item.PaymentStatus
	resp.PaymentNote = item.PaymentNote

	resp.School = item.School.ToPublicResponse()
	resp.StudentEnroll = item.StudentEnroll.ToStudentEnrollPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Payment) []data.PaymentResponse {
	resp := make([]data.PaymentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
