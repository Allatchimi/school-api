package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/common/student/data"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelLevel "api/services/school/university/level/model"
	modelUser "api/services/user/user/model"
	"time"
)

type StudentPreEnroll struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	ClassID int64                       `gorm:"default:null"`
	Class   *modelClass.HighschoolClass `gorm:"default:null;foreignKey:ClassID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	LevelDomainID int64                             `gorm:"default:null"`
	LevelDomain   *modelLevel.UniversityLevelDomain `gorm:"default:null;foreignKey:LevelDomainID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Status         string `gorm:"default:null"`
	StatusFeedback string `gorm:"default:null;type:text"`

	Message       string     `gorm:"default:null;type:text"`
	Gender        string     `gorm:"default:null"`
	FirstName     string     `gorm:"default:null"`
	LastName      string     `gorm:"default:null"`
	Birthday      *time.Time `gorm:"default:null"`
	BirthLocation string     `gorm:"default:null"`
	Document1     string     `gorm:"default:null"`
	Document2     string     `gorm:"default:null"`
	Document3     string     `gorm:"default:null"`
	Document4     string     `gorm:"default:null"`
	Document5     string     `gorm:"default:null"`
}

func (item *StudentPreEnroll) ToResponse() *data.StudentPreEnrollResponse {
	if item == nil {
		return &data.StudentPreEnrollResponse{}
	}
	resp := &data.StudentPreEnrollResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToResponse()
	resp.LevelDomain = item.LevelDomain.ToResponse()
	resp.Class = item.Class.ToResponse()
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

func ToStudentPreEnrollResponseList(itemList []StudentPreEnroll) []data.StudentPreEnrollResponse {
	resp := make([]data.StudentPreEnrollResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
