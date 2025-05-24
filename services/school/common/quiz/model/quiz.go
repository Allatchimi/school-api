package model

import (
	"api/common/types"
	"api/services/school/common/meeting/data"
	schoolModel "api/services/school/common/school/model"
	userModel "api/services/user/user/model"
)

type Quiz struct {
	types.BaseGormModel

	UserID int64           `gorm:"default:null"`
	User   *userModel.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Quiz) ToResponse() *data.QuizResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizResponse{}
	if resp.User != nil {
		resp.User.Role = nil
		resp.User.Info = nil
		resp.User.Mfa = nil
	}
	if resp.School != nil {
		resp.School.Config = nil
		resp.School.Info = nil
	}

	resp.User = item.User.ToResponse()
	resp.School = item.School.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Quiz) ToPublicResponse() *data.QuizPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizPublicResponse{}
	if resp.User != nil {
		resp.User.Info = nil
	}
	if resp.School != nil {
		resp.School.Info = nil
	}
	return resp
}

func ToResponseList(itemList []Quiz) []data.QuizResponse {
	resp := make([]data.QuizResponse, len(itemList))
	for index, school := range itemList {
		resp[index] = *school.ToResponse()
	}
	return resp
}
