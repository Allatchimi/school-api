package model

import (
	"api/common/types"
	"api/services/school/common/schedule/data"
	schoolModel "api/services/school/common/school/model"
	yearModel "api/services/school/common/year/model"
	classModel "api/services/school/highschool/class/model"
	Unitmodel "api/services/school/university/unit/model"
)

type Schedule struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *yearModel.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                              `gorm:"default:null"`
	ClassSubject   *classModel.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *Unitmodel.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Type         string `gorm:"default:null"`
	DayOfTheWeek string `gorm:"default:null"`
	RepeatCount  string `gorm:"default:null"`
	RepeatType   string `gorm:"default:null"`
	StartTime    string `gorm:"default:null"`
	EndTime      string `gorm:"default:null"`
}

func (item *Schedule) ToResponse() *data.ScheduleResponse {
	if item == nil {
		return nil
	}
	resp := &data.ScheduleResponse{}
	resp.Type = item.Type
	resp.DayOfTheWeek = item.DayOfTheWeek
	resp.RepeatCount = item.RepeatCount
	resp.RepeatType = item.RepeatType
	resp.StartTime = item.StartTime
	resp.EndTime = item.EndTime

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Schedule) ToPublicResponse() *data.SchedulePublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.SchedulePublicResponse{}
	resp.Type = item.Type
	resp.DayOfTheWeek = item.DayOfTheWeek
	resp.RepeatCount = item.RepeatCount
	resp.RepeatType = item.RepeatType
	resp.StartTime = item.StartTime
	resp.EndTime = item.EndTime

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	return resp
}

func ToScheduleResponseList(itemList []Schedule) []data.ScheduleResponse {
	resp := make([]data.ScheduleResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
