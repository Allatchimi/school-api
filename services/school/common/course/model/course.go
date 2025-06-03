package model

import (
	"api/common/types"
	"api/services/school/common/course/data"
	schoolModel "api/services/school/common/school/model"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
)

type Course struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                                     `gorm:"default:null"`
	ClassSubject   *modelClassSubject.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Content     string `gorm:"default:null"`

	Documents []CourseDocument `gorm:"foreignKey:CourseID;references:ID"`
	Videos    []CourseVideo    `gorm:"foreignKey:CourseID;references:ID"`
}

func (item *Course) ToResponse() *data.CourseResponse {
	if item == nil {
		return nil
	}
	resp := &data.CourseResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.Content = item.Content

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()

	resp.Documents = ToCourseDocumentResponseList(item.Documents)
	resp.Videos = ToCourseVideoResponseList(item.Videos)

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Course) ToPublicResponse() *data.CoursePublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.CoursePublicResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.Content = item.Content

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()

	return resp
}

func ToResponseList(itemList []Course) []data.CoursePublicResponse {
	resp := make([]data.CoursePublicResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToPublicResponse()
	}
	return resp
}
