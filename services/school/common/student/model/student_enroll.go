package model

import (
	"api/common/types"
	"api/services/school/common/student/data"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
	"time"
)

type StudentEnroll struct {
	types.BaseGormModel

	StudentID int64    `gorm:"default:null"`
	Student   *Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassID int64                              `gorm:"default:null"`
	Class   *modelClassSubject.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Email       string `gorm:"default:null"`
	PhoneNumber uint64 `gorm:"default:null"`

	Message    string `gorm:"default:null"`
	Origin     string `gorm:"default:null"`
	IsAccepted bool   `gorm:"default:null"`

	Gender        string     `gorm:"default:null"`
	FirstName     string     `gorm:"default:null"`
	LastName      string     `gorm:"default:null"`
	Birthday      *time.Time `gorm:"default:null"`
	BirthLocation string     `gorm:"default:null"`

	Document1 string `gorm:"default:null"`
	Document2 string `gorm:"default:null"`
	Document3 string `gorm:"default:null"`
	Document4 string `gorm:"default:null"`
	Document5 string `gorm:"default:null"`
}

func (item *StudentEnroll) ToStudentEnrollResponse() *data.StudentEnrollResponse {
	if item == nil {
		return nil
	}
	resp := &data.StudentEnrollResponse{}
	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber

	resp.Message = item.Message
	resp.Origin = item.Origin
	resp.IsAccepted = item.IsAccepted

	resp.Gender = item.Gender
	resp.FirstName = item.FirstName
	resp.Birthday = item.Birthday
	resp.BirthLocation = item.BirthLocation

	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Document3 = item.Document3
	resp.Document4 = item.Document4
	resp.Document5 = item.Document5

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
	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber

	resp.Message = item.Message
	resp.Origin = item.Origin
	resp.IsAccepted = item.IsAccepted

	resp.Gender = item.Gender
	resp.FirstName = item.FirstName
	resp.Birthday = item.Birthday
	resp.BirthLocation = item.BirthLocation

	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Document3 = item.Document3
	resp.Document4 = item.Document4
	resp.Document5 = item.Document5

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
