package model

import (
	"api/common/types"
	"api/services/school/common/quiz/data"
	schoolModel "api/services/school/common/school/model"
	modelYear "api/services/school/common/year/model"
	modelClassSubject "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
)

type Quiz struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                                     `gorm:"default:null"`
	ClassSubject   *modelClassSubject.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Questions []QuizQuestion `gorm:"default:null;foreignKey:QuizID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Status      string `gorm:"default:null"`
}
type QuizWithoutFk struct {
	types.BaseGormModel
	SchoolID       int64 `gorm:"index"`
	YearID         int64 `gorm:"index"`
	ClassSubjectID int64 `gorm:"index"`
	UnitID         int64 `gorm:"index"`

	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Status      string `gorm:"default:null"`
}

func (QuizWithoutFk) TableName() string {
	return "quizzes"
}

func (item *Quiz) ToResponse() *data.QuizResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuizResponse{}
	resp.Title = item.Title
	resp.Description = item.Description
	resp.Status = item.Status

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToPublicResponse()
	resp.ClassSubject = item.ClassSubject.ToClassSubjectPublicResponse()
	resp.Unit = item.Unit.ToPublicResponse()
	resp.Questions = ToQuizQuestionResponseListWithOptions(item.Questions)

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
