package model

import (
	"api/common/types"
	"api/services/school/common/request/data"
	schoolModel "api/services/school/common/school/model"
	studentmodel "api/services/school/common/student/model"
	yearModel "api/services/school/common/year/model"
	classModel "api/services/school/highschool/class/model"
	sequenceModel "api/services/school/highschool/sequence/model"
	unitmodel "api/services/school/university/unit/model"
)

type Request struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *yearModel.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                              `gorm:"default:null"`
	ClassSubject   *classModel.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SequenceID int64                             `gorm:"default:null"`
	Sequence   *sequenceModel.HighschoolSequence `gorm:"default:null;foreignKey:SequenceID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *unitmodel.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentID int64                 `gorm:"default:null"`
	Student   *studentmodel.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Status         string `gorm:"default:null"`
	StatusFeedback string `gorm:"default:null"`
	Audience       string `gorm:"default:null"`
	Title          string `gorm:"default:null"`
	Message        string `gorm:"default:null"`

	Document1 string `gorm:"default:null"`
	Document2 string `gorm:"default:null"`
	Document3 string `gorm:"default:null"`
	Document4 string `gorm:"default:null"`
	Document5 string `gorm:"default:null"`
}

func (item *Request) ToResponse() *data.RequestResponse {
	if item == nil {
		return nil
	}
	resp := &data.RequestResponse{}
	resp.Status = item.Status
	resp.StatusFeedback = item.StatusFeedback
	resp.Audience = item.Audience
	resp.Title = item.Title
	resp.Message = item.Message

	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Document3 = item.Document3
	resp.Document4 = item.Document4
	resp.Document5 = item.Document5

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Sequence = item.Sequence.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.Student = item.Student.ToStudentPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Request) ToPublicResponse() *data.RequestPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.RequestPublicResponse{}
	resp.Status = item.Status
	resp.StatusFeedback = item.StatusFeedback
	resp.Audience = item.Audience
	resp.Title = item.Title
	resp.Message = item.Message

	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Document3 = item.Document3
	resp.Document4 = item.Document4
	resp.Document5 = item.Document5

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Sequence = item.Sequence.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.Student = item.Student.ToStudentPublicResponse()
	return resp
}

func ToRequestResponseList(itemList []Request) []data.RequestResponse {
	resp := make([]data.RequestResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
