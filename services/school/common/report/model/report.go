package model

import (
	"api/common/types"
	"api/services/school/common/report/data"
	modelSchool "api/services/school/common/school/model"
	modelStudent "api/services/school/common/student/model"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelSequence "api/services/school/highschool/sequence/model"
	modelUnit "api/services/school/university/unit/model"
)

type Report struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                                     `gorm:"default:null"`
	ClassSubject   *modelClassSubject.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SequenceID int64                             `gorm:"default:null"`
	Sequence   *modelSequence.HighschoolSequence `gorm:"default:null;foreignKey:SequenceID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentID int64                 `gorm:"default:null"`
	Student   *modelStudent.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Coefficient int     `gorm:"default:null"`
	Credit      int     `gorm:"default:null"`
	Value       float64 `gorm:"default:null"`
	Notation    float64 `gorm:"default:null"`
}

func (item *Report) ToResponse() *data.ReportResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportResponse{}
	resp.Coefficient = item.Coefficient
	resp.Credit = item.Credit
	resp.Value = item.Value
	resp.Notation = item.Notation

	resp.Student = item.Student.ToStudentResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectResponse()
	resp.Sequence = item.Sequence.ToResponse()
	resp.Unit = item.Unit.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Report) ToPublicResponse() *data.ReportPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportPublicResponse{}
	resp.Coefficient = item.Coefficient
	resp.Credit = item.Credit
	resp.Value = item.Value
	resp.Notation = item.Notation

	resp.Student = item.Student.ToStudentResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectResponse()
	resp.Sequence = item.Sequence.ToResponse()
	resp.Unit = item.Unit.ToResponse()
	return resp
}

func ToReportResponseList(itemList []Report) []data.ReportResponse {
	resp := make([]data.ReportResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
