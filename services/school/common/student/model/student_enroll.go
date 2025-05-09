package model

import (
	"api/common/types"
	"api/services/school/common/student/data"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
)

type StudentEnroll struct {
	types.BaseGormModel

	StudentID int64    `gorm:"default:null"`
	Student   *Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                              `gorm:"default:null"`
	Class   *modelClassSubject.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	PaymentID int64    `gorm:"default:null"`
	Payment   *Payment `gorm:"default:null;foreignKey:PaymentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Status string `gorm:"default:null"`
	Status string `gorm:"default:null"`

	File1 string `gorm:"default:null"`
	File2 string `gorm:"default:null"`
	File3 string `gorm:"default:null"`
	File4 string `gorm:"default:null"`
	File5 string `gorm:"default:null"`
}

func (item *StudentEnroll) ToStudentEnrollResponse() *data.StudentEnrollResponse {
	if item == nil {
		return nil
	}
	resp := &data.StudentEnrollResponse{}
	resp.Student = item.Student.ToStudentPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.LevelDomain = item.LevelDomain.ToLevelDomainPublicResponse()
	resp.Class = item.Class.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *StudentEnroll) ToStudentEnrollPublicResponse() *data.StudentEnrollPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.StudentEnrollPublicResponse{}
	resp.Student = item.Student.ToStudentPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.LevelDomain = item.LevelDomain.ToLevelDomainPublicResponse()
	resp.Class = item.Class.ToPublicResponse()
	return resp
}

func ToStudentEnrollResponseList(itemList []StudentEnroll) []data.StudentEnrollResponse {
	resp := make([]data.StudentEnrollResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToStudentEnrollResponse()
	}
	return resp
}
