package model

import (
	"api/common/types"
	"api/services/school/common/exam/data"
	modelSchool "api/services/school/common/school/model"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelSequence "api/services/school/highschool/sequence/model"
	modelUnit "api/services/school/university/unit/model"
	"time"
)

type Exam struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	TypeID int64     `gorm:"default:null"`
	Type   *ExamType `gorm:"default:null;foreignKey:TypeID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                              `gorm:"default:null"`
	ClassSubject   *modelClass.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SequenceID int64                             `gorm:"default:null"`
	Sequence   *modelSequence.HighschoolSequence `gorm:"default:null;foreignKey:SequenceID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Status          string     `gorm:"default:null"`
	Percentage      int        `gorm:"default:null"`
	Description     string     `gorm:"default:null"`
	LocationType    string     `gorm:"default:null"`
	LocationDetails string     `gorm:"default:null"`
	Requirements    string     `gorm:"default:null"`
	AllowedItems    string     `gorm:"default:null"`
	StartDate       *time.Time `gorm:"default:null"`
	EndDate         *time.Time `gorm:"default:null"`
}

func (item *Exam) ToResponse() *data.ExamResponse {
	if item == nil {
		return nil
	}
	resp := &data.ExamResponse{}
	resp.Status = item.Status
	resp.Percentage = item.Percentage
	resp.Description = item.Description
	resp.LocationType = item.LocationType
	resp.LocationDetails = item.LocationDetails
	resp.Requirements = item.Requirements
	resp.AllowedItems = item.AllowedItems
	resp.StartDate = item.StartDate
	resp.EndDate = item.EndDate

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToResponse()
	resp.Type = item.Type.ToResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectResponse()
	resp.Sequence = item.Sequence.ToResponse()
	resp.Unit = item.Unit.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToExamResponseList(itemList []Exam) []data.ExamResponse {
	resp := make([]data.ExamResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
