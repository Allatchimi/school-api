package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	modelDomain "api/services/school/university/domain/model"
	"api/services/school/university/level/data"
	"time"
)

type UniversityLevelDomain struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelID int64            `gorm:"default:null"`
	Level   *UniversityLevel `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	DomainID int64                         `gorm:"default:null"`
	Domain   *modelDomain.UniversityDomain `gorm:"default:null;foreignKey:DomainID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Fees         int64      `gorm:"default null"`
	Program      string     `gorm:"default null"`
	Requirements string     `gorm:"default null"`
	IsValid      bool       `gorm:"default:true"`
	InvalidDate  *time.Time `gorm:"default:null"`
}

func (item *UniversityLevelDomain) ToLevelDomainResponse() *data.LevelDomainResponse {
	if item == nil {
		return nil
	}
	resp := &data.LevelDomainResponse{}
	resp.Fees = item.Fees
	resp.Program = item.Program
	resp.Requirements = item.Requirements
	resp.IsValid = item.IsValid
	resp.InvalidDate = item.InvalidDate

	resp.School = item.School.ToPublicResponse()
	resp.Domain = item.Domain.ToResponse()
	resp.Level = item.Level.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToLevelDomainResponseList(itemList []UniversityLevelDomain) []data.LevelDomainResponse {
	resp := make([]data.LevelDomainResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToLevelDomainResponse()
	}
	return resp
}
