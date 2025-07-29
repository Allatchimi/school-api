package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
	modelSchool "api/services/school/common/school/model"
	modelUser "api/services/user/user/model"
	"time"
)

type ParentAssign struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentListID  string     `gorm:"default:null"`
	Status         string     `gorm:"default:null"`
	StatusFeedback string     `gorm:"default:null;type:text"`
	Message        string     `gorm:"default:null;type:text"`
	Gender         string     `gorm:"default:null"`
	FirstName      string     `gorm:"default:null"`
	LastName       string     `gorm:"default:null"`
	Birthday       *time.Time `gorm:"default:null"`
	BirthLocation  string     `gorm:"default:null"`
	Document1      string     `gorm:"default:null"`
	Document2      string     `gorm:"default:null"`
	Document3      string     `gorm:"default:null"`
	Document4      string     `gorm:"default:null"`
	Document5      string     `gorm:"default:null"`
}

func (item *ParentAssign) ToResponse() *data.ParentAssignResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentAssignResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.StudentListID = item.StudentListID

	resp.School = item.School.ToPublicResponse()
	resp.User = item.User.ToPublicResponse()

	resp.Status = item.Status
	resp.StatusFeedback = item.StatusFeedback
	resp.Message = item.Message
	resp.Gender = item.Gender
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.Birthday = item.Birthday
	resp.BirthLocation = item.BirthLocation
	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Document3 = item.Document3
	resp.Document4 = item.Document4
	resp.Document5 = item.Document5
	return resp
}

func ToParentAssignResponseList(itemList []ParentAssign) []data.ParentAssignResponse {
	resp := make([]data.ParentAssignResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
