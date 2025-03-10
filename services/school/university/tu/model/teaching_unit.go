package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	modelDomain "api/services/school/university/domain/model"
	modelLevel "api/services/school/university/level/model"
	"api/services/school/university/tu/data"
	"time"
)

type TeachingUnit struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	DomainID int64                         `gorm:"default:null"`
	Domain   *modelDomain.UniversityDomain `gorm:"default:null;foreignKey:DomainID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelID int64                       `gorm:"default:null"`
	Level   *modelLevel.UniversityLevel `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name         string `gorm:"default:null"`
	Description  string `gorm:"default:null"`
	Credit       int    `gorm:"default:1"`
	Semester     int    `gorm:"default:1"`
	Program      string `gorm:"default null"`
	Requirements string `gorm:"default null"`

	IsValid     bool       `gorm:"default:true"`
	InvalidDate *time.Time `gorm:"default:null"`

	TeachingUnitProfessors []TeachingUnitProfessor `gorm:"default:null;foreignKey:TeachingUnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *TeachingUnit) ToResponse() *data.TeachingUnitResponse {
	resp := &data.TeachingUnitResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.SchoolID = item.SchoolID
	resp.DomainID = item.DomainID
	resp.LevelID = item.LevelID

	resp.Name = item.Name
	resp.Description = item.Description
	resp.Credit = item.Credit
	resp.Semester = item.Semester
	resp.Program = item.Program
	resp.Requirements = item.Requirements

	resp.IsValid = item.IsValid
	resp.InvalidDate = item.InvalidDate

	resp.TeachingUnitProfessors = ToTeachingUnitProfessorResponseList(item.TeachingUnitProfessors)
	return resp
}

func ToTeachingUnitResponseList(itemList []TeachingUnit) []data.TeachingUnitResponse {
	resp := make([]data.TeachingUnitResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
