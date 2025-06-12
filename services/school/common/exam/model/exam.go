package model

import (
	"api/common/types"
	"api/services/school/common/exam/data"
	schoolModel "api/services/school/common/school/model"
	yearModel "api/services/school/common/year/model"
	classModel "api/services/school/highschool/class/model"
	sequenceModel "api/services/school/highschool/sequence/model"
	Unitmodel "api/services/school/university/unit/model"
	"time"
)

type Exam struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *yearModel.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	TypeID int64     `gorm:"default:null"`
	Type   *ExamType `gorm:"default:null;foreignKey:TypeID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *Unitmodel.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                              `gorm:"default:null"`
	ClassSubject   *classModel.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SequenceID int64                             `gorm:"default:null"`
	Sequence   *sequenceModel.HighschoolSequence `gorm:"default:null;foreignKey:SequenceID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

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
	resp.Percentage = item.Percentage
	resp.Description = item.Description
	resp.LocationType = item.LocationType
	resp.LocationDetails = item.LocationDetails
	resp.Requirements = item.Requirements
	resp.AllowedItems = item.AllowedItems
	resp.StartDate = item.StartDate
	resp.EndDate = item.EndDate

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.Type = item.Type.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Sequence = item.Sequence.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Exam) ToPublicResponse() *data.ExamPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ExamPublicResponse{}
	resp.Percentage = item.Percentage
	resp.Description = item.Description
	resp.LocationType = item.LocationType
	resp.LocationDetails = item.LocationDetails
	resp.Requirements = item.Requirements
	resp.AllowedItems = item.AllowedItems
	resp.StartDate = item.StartDate
	resp.EndDate = item.EndDate

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.Type = item.Type.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Sequence = item.Sequence.ToPublicResponse()
	return resp
}

func ToExamResponseList(itemList []Exam) []data.ExamResponse {
	resp := make([]data.ExamResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
