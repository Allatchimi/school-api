package model

import (
	"api/common/types"
	"api/services/school/common/schedule/data"
	modelSchool "api/services/school/common/school/model"
	modelYear "api/services/school/common/year/model"
	modelClass "api/services/school/highschool/class/model"
	modelUnit "api/services/school/university/unit/model"
	"time"
)

type Schedule struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	YearID int64           `gorm:"default:null"`
	Year   *modelYear.Year `gorm:"default:null;foreignKey:YearID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ClassSubjectID int64                              `gorm:"default:null"`
	ClassSubject   *modelClass.HighschoolClassSubject `gorm:"default:null;foreignKey:ClassSubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UnitID int64                     `gorm:"default:null"`
	Unit   *modelUnit.UniversityUnit `gorm:"default:null;foreignKey:UnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Type           string     `gorm:"default:null"`
	DayOfTheWeek   string     `gorm:"default:null"`
	RepeatCount    int        `gorm:"default:null"`
	RepeatType     string     `gorm:"default:null"`
	StartTime      string     `gorm:"default:null"`
	EndTime        string     `gorm:"default:null"`
	StartCountDate *time.Time `gorm:"default:null"`
	IsValid        bool       `gorm:"default:null"`
	InvalidDate    *time.Time `gorm:"default:null"`
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
	resp.StartCountDate = item.StartCountDate
	resp.IsValid = item.IsValid
	resp.InvalidDate = item.InvalidDate

	resp.School = item.School.ToPublicResponse()
	resp.Year = item.Year.ToResponse()
	resp.ClassSubject = item.ClassSubject.ToResponse()
	resp.Unit = item.Unit.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToScheduleResponseList(itemList []Schedule) []data.ScheduleResponse {
	resp := make([]data.ScheduleResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}

func ToScheduleWeeklyViewResponseList(itemList []Schedule) []data.ScheduleWeeklyViewResponse {
	type key struct {
		StartTime string
		EndTime   string
	}

	groupMap := make(map[key]*data.ScheduleWeeklyViewResponse)

	for _, schedule := range itemList {
		k := key{StartTime: schedule.StartTime, EndTime: schedule.EndTime}

		// Check if the time group exists
		if _, exists := groupMap[k]; !exists {
			groupMap[k] = &data.ScheduleWeeklyViewResponse{
				StartTime: k.StartTime,
				EndTime:   k.EndTime,
			}
		}

		// Add schedule to the correct day
		switch schedule.DayOfTheWeek {
		case "monday":
			groupMap[k].Monday = append(groupMap[k].Monday, *schedule.ToResponse())
		case "tuesday":
			groupMap[k].Tuesday = append(groupMap[k].Tuesday, *schedule.ToResponse())
		case "wednesday":
			groupMap[k].Wednesday = append(groupMap[k].Wednesday, *schedule.ToResponse())
		case "thursday":
			groupMap[k].Thursday = append(groupMap[k].Thursday, *schedule.ToResponse())
		case "friday":
			groupMap[k].Friday = append(groupMap[k].Friday, *schedule.ToResponse())
		case "saturday":
			groupMap[k].Saturday = append(groupMap[k].Saturday, *schedule.ToResponse())
		case "sunday":
			groupMap[k].Sunday = append(groupMap[k].Sunday, *schedule.ToResponse())
		}
	}

	// Conversion de map en slice
	result := make([]data.ScheduleWeeklyViewResponse, 0, len(groupMap))
	for _, v := range groupMap {
		result = append(result, *v)
	}

	return result
}

func ListAppendCommonSchedules(dest []Schedule, src []ScheduleCommon) []Schedule {
	defaultSize := len(dest)
	commonSize := len(src)
	result := make([]Schedule, defaultSize+commonSize)
	copy(result, dest)
	for index := range src {
		schedule := ConvertScheduleCommonToSchedule(&src[index])
		if schedule != nil {
			result[defaultSize+index] = *schedule
		}
	}
	return result
}

func ConvertScheduleCommonToSchedule(item *ScheduleCommon) (result *Schedule) {
	if item == nil {
		return
	}
	result = &Schedule{
		SchoolID: item.SchoolID,
		School:   item.School,
		YearID:   item.YearID,
		Year:     item.Year,

		Type:         item.Type,
		DayOfTheWeek: item.DayOfTheWeek,
		RepeatCount:  item.RepeatCount,
		RepeatType:   item.RepeatType,
		StartTime:    item.StartTime,
		EndTime:      item.EndTime,
		IsValid:      item.IsValid,
		InvalidDate:  item.InvalidDate,
	}
	result.ID = item.ID
	result.CreatedAt = item.CreatedAt
	result.UpdatedAt = item.UpdatedAt
	return
}
