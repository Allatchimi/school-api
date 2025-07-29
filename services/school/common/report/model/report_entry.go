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

type ReportEntry struct {
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

	Coefficient      int     `gorm:"default:null"`
	Credit           int     `gorm:"default:null"`
	Value            float64 `gorm:"default:null"`
	Notation         float64 `gorm:"default:null"`
	Grade            string  `gorm:"default:null"`
	GradeDescription string  `gorm:"default:null"`
	IsRetry          bool    `gorm:"default:null"`
	RetryCount       int64   `gorm:"default:null"`
	RetryDetails     string  `gorm:"default:null"`
}

func (item *ReportEntry) ToResponse() *data.ReportEntryResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportEntryResponse{}
	resp.Coefficient = item.Coefficient
	resp.Credit = item.Credit
	resp.Value = item.Value
	resp.Notation = item.Notation
	resp.Grade = item.Grade
	resp.GradeDescription = item.GradeDescription
	resp.IsRetry = item.IsRetry
	resp.RetryCount = item.RetryCount
	resp.RetryDetails = item.RetryDetails

	resp.Student = item.Student.ToResponse()
	resp.ClassSubject = item.ClassSubject.ToResponse()
	resp.Sequence = item.Sequence.ToResponse()
	resp.Unit = item.Unit.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToReportEntryResponseList(itemList []ReportEntry) []data.ReportEntryResponse {
	resp := make([]data.ReportEntryResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
