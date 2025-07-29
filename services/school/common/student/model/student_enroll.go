package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/common/student/data"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
)

type StudentEnroll struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                       `gorm:"default:null"`
	Class   *modelClass.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelDomainID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentID int64    `gorm:"default:null"`
	Student   *Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Origin         string `gorm:"default:null"`
	OriginFeedback string `gorm:"default:null;type:text"`
}

func (item *StudentEnroll) ToResponse() *data.StudentEnrollResponse {
	if item == nil {
		return nil
	}
	resp := &data.StudentEnrollResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToResponse()
	resp.Class = item.Class.ToResponse()
	resp.LevelDomain = item.LevelDomain.ToResponse()
	resp.Student = item.Student.ToPublicResponse()

	resp.Origin = item.Origin
	resp.OriginFeedback = item.OriginFeedback
	return resp
}

func ToStudentEnrollResponseList(itemList []StudentEnroll) []data.StudentEnrollResponse {
	resp := make([]data.StudentEnrollResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
