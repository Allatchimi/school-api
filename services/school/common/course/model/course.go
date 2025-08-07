package model

import (
	"api/common/types"
	"api/services/school/common/course/data"
	modelSchool "api/services/school/common/school/model"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
)

type Course struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	ClassSubjectID int64                                     `gorm:"default:null"`
	ClassSubject   *modelClassSubject.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Documents []CourseDocument `gorm:"foreignKey:CourseID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`
	Videos    []CourseVideo    `gorm:"foreignKey:CourseID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Content     string `gorm:"default:null"`
}

func (item *Course) ToResponse() *data.CourseResponse {
	if item == nil {
		return &data.CourseResponse{}
	}
	resp := &data.CourseResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.Content = item.Content

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToResponse()
	resp.ClassSubject = item.ClassSubject.ToResponse()
	resp.Unit = item.Unit.ToResponse()
	resp.Documents = ToCourseDocumentResponseList(item.Documents)
	resp.Videos = ToCourseVideoResponseList(item.Videos)

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Course) []data.CourseResponse {
	resp := make([]data.CourseResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
