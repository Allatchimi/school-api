package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	modelLevel "api/services/school/university/level/model"
	modelSemester "api/services/school/university/semester/model"
	"api/services/school/university/unit/data"
	"time"
)

type UniversityUnit struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelDomainID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SemesterID int64                             `gorm:"default:null"`
	Semester   *modelSemester.UniversitySemester `gorm:"default:null;foreignKey:SemesterID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name         string     `gorm:"default:null"`
	Description  string     `gorm:"default:null"`
	Credit       int        `gorm:"default:null"`
	Program      string     `gorm:"default null"`
	Requirements string     `gorm:"default null"`
	IsValid      bool       `gorm:"default:true"`
	InvalidDate  *time.Time `gorm:"default:null"`
}

func (item *UniversityUnit) ToResponse() *data.UnitResponse {
	if item == nil {
		return nil
	}
	resp := &data.UnitResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.LevelDomain = item.LevelDomain.ToResponse()
	resp.Semester = item.Semester.ToResponse()

	resp.Name = item.Name
	resp.Description = item.Description
	resp.Credit = item.Credit
	resp.Program = item.Program
	resp.Requirements = item.Requirements
	resp.IsValid = item.IsValid
	resp.InvalidDate = item.InvalidDate
	return resp
}

func ToUnitResponseList(itemList []UniversityUnit) []data.UnitResponse {
	resp := make([]data.UnitResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
