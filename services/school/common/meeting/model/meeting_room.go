package model

import (
	"api/common/types"
	"api/services/school/common/meeting/data"
	schoolModel "api/services/school/common/school/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
)

type MeetingRoom struct {
	types.BaseGormModel

	ApiRoomID string `gorm:"default:null"`

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                                     `gorm:"default:null"`
	ClassSubject   *modelClassSubject.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *MeetingRoom) ToResponse() *data.MeetingRoomResponse {
	if item == nil {
		return nil
	}
	resp := &data.MeetingRoomResponse{}
	resp.ApiRoomID = item.ApiRoomID

	resp.School = item.School.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *MeetingRoom) ToPublicResponse() *data.MeetingRoomPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.MeetingRoomPublicResponse{}
	resp.ApiRoomID = item.ApiRoomID

	resp.School = item.School.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	return resp
}

func ToResponseList(itemList []MeetingRoom) []data.MeetingRoomResponse {
	resp := make([]data.MeetingRoomResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
