package model

import (
	"time"

	"api/common/types"
	"api/services/user/user/data"
)

type UserInfo struct {
	types.BaseGormModel
	Gender        string     `gorm:"default:null"`
	Username      string     `gorm:"default:null"`
	FirstName     string     `gorm:"default:null"`
	LastName      string     `gorm:"default:null"`
	Birthday      *time.Time `gorm:"default:null"`
	BirthLocation string     `gorm:"default:null"`
	Address       string     `gorm:"default:null"`
	Language      string     `gorm:"default:en"`
	Image         string     `gorm:"default:null"`
}

func (item *UserInfo) ToResponse() *data.UserInfoResponse {
	if item == nil {
		return &data.UserInfoResponse{}
	}
	resp := &data.UserInfoResponse{}
	resp.Gender = item.Gender
	resp.Username = item.Username
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.Birthday = item.Birthday
	resp.BirthLocation = item.BirthLocation
	resp.Address = item.Address
	resp.Language = item.Language
	resp.Image = item.Image
	return resp
}

func (item *UserInfo) ToPublicResponse() *data.UserInfoPublicResponse {
	if item == nil {
		return &data.UserInfoPublicResponse{}
	}
	resp := &data.UserInfoPublicResponse{}
	resp.Gender = item.Gender
	resp.Username = item.Username
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.Image = item.Image
	return resp
}
