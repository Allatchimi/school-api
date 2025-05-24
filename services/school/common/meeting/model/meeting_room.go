package model

import (
	"api/common/types"
	"api/services/school/common/meeting/data"
	schoolModel "api/services/school/common/school/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
)

type MeetingRoom struct {
	types.BaseGormModel

	ApiRoomID string `gorm:"default:null"`

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                              `gorm:"default:null"`
	Class   *modelClassSubject.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *MeetingRoom) ToResponse() *data.MeetingRoomResponse {
	if item == nil {
		return nil
	}
	resp := &data.MeetingRoomResponse{}
	resp.ApiRoomID = item.ApiRoomID

	resp.School = item.School.ToPublicResponse()
	resp.LevelDomain = item.LevelDomain.ToLevelDomainPublicResponse()
	resp.Class = item.Class.ToPublicResponse()

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
	resp.LevelDomain = item.LevelDomain.ToLevelDomainPublicResponse()
	resp.Class = item.Class.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []MeetingRoom) []data.MeetingRoomResponse {
	resp := make([]data.MeetingRoomResponse, len(itemList))
	for index, school := range itemList {
		resp[index] = *school.ToResponse()
	}
	return resp
}
