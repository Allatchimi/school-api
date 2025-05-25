package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
	schoolModel "api/services/school/common/school/model"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
	"time"
)

type Quiz struct {
	types.BaseGormModel
	Title       string     `gorm:"default:null"`
	Description string     `gorm:"default:null"`
	StartDate   *time.Time `gorm:"default:null"`
	EndDate     *time.Time `gorm:"default:null"`

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                                     `gorm:"default:null"`
	ClassSubject   *modelClassSubject.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Questions []QuizQuestion `gorm:"foreignKey:QuizID;references:ID"`
}

func (item *Quiz) ToResponse() *data.QuizResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.StartDate = item.StartDate
	resp.EndDate = item.EndDate

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()

	resp.Questions = ToQuizQuestionResponseList(item.Questions)

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToQuizResponseList(itemList []Quiz) []data.QuizResponse {
	resp := make([]data.QuizResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
