package model

import (
	"api/common/types"
	"api/services/school/common/meeting/data"
	modelSchool "api/services/school/common/school/model"
	modelClass "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
)

type MeetingRoom struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	ClassSubjectID int64                              `gorm:"default:null"`
	ClassSubject   *modelClass.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	ApiRoomID string `gorm:"default:null"`
}

func (item *MeetingRoom) ToResponse() *data.MeetingRoomResponse {
	if item == nil {
		return &data.MeetingRoomResponse{}
	}
	resp := &data.MeetingRoomResponse{}
	resp.ApiRoomID = item.ApiRoomID

	resp.School = item.School.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToResponse()
	resp.Unit = item.Unit.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []MeetingRoom) []data.MeetingRoomResponse {
	resp := make([]data.MeetingRoomResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
