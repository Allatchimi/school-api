package report

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/report/data"
	"api/services/school/common/report/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) CreateReportEntry(data *model.ReportEntry) (*model.ReportEntry, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateReportGrade(data *model.ReportGrade) (*model.ReportGrade, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateReportConfig(data *model.ReportConfig) (*model.ReportConfig, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateReportEntryByID(id int64, data *model.ReportEntry) (*model.ReportEntry, error) {
	foundItem, err := repository.GetReportEntryByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}

	result := &model.ReportEntry{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":        data.SchoolID,
			"year_id":          data.YearID,
			"class_subject_id": data.ClassSubjectID,
			"sequence_id":      data.SequenceID,
			"unit_id":          data.UnitID,
			"student_id":       data.StudentID,

			"coefficient": data.Coefficient,
			"credit":      data.Credit,
			"value":       data.Value,
			"notation":    data.Notation,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateReportGradeByID(id int64, data *model.ReportGrade) (*model.ReportGrade, error) {
	foundItem, err := repository.GetReportGradeByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}

	result := &model.ReportGrade{}
	return result, repository.Db.Model(&model.ReportGrade{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id": data.SchoolID,

			"name":        data.Name,
			"description": data.Description,

			"minimum_result":         data.MinimumResult,
			"maximum_result":         data.MaximumResult,
			"include_minimum_result": data.IncludeMinimumResult,
			"include_maximum_result": data.IncludeMaximumResult,
			"correspondence":         data.Correspondence,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateReportConfigByID(id int64, data *model.ReportConfig) (*model.ReportConfig, error) {
	foundItem, err := repository.GetReportConfigByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}

	result := &model.ReportConfig{}
	return result, repository.Db.Model(&model.ReportConfig{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id": data.SchoolID,

			"notation":                 data.Notation,
			"notation_minimum_success": data.NotationMinimumSuccess,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteReportEntryByID(id int64) (int64, error) {
	foundItem, err := repository.GetReportEntryByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.ReportEntry{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteReportGradeByID(id int64) (int64, error) {
	foundItem, err := repository.GetReportGradeByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.ReportGrade{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteReportConfigByID(id int64) (int64, error) {
	foundItem, err := repository.GetReportConfigByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.ReportConfig{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleReportEntryByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.ReportEntry{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleReportGradeByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.ReportGrade{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleReportConfigByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.ReportConfig{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetReportEntryByID(id int64) (*model.ReportEntry, error) {
	result := &model.ReportEntry{}
	return result, repository.Db.Model(&model.ReportEntry{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetReportGradeByID(id int64) (*model.ReportGrade, error) {
	result := &model.ReportGrade{}
	return result, repository.Db.Model(&model.ReportGrade{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetReportConfigByID(id int64) (*model.ReportConfig, error) {
	result := &model.ReportConfig{}
	return result, repository.Db.Model(&model.ReportConfig{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectReportGrade(item *model.ReportGrade) (*model.ReportGrade, error) {
	result := &model.ReportGrade{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportGrade{
		SchoolID: item.SchoolID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectReportConfig(item *model.ReportConfig) (*model.ReportConfig, error) {
	result := &model.ReportConfig{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportConfig{
		SchoolID: item.SchoolID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsReportGrade(item1 *model.ReportGrade, item2 *model.ReportGrade) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID) {
		return true
	}
	return false
}

func (repository *Repository) AreSameUniqueObjectsReportConfig(item1 *model.ReportConfig, item2 *model.ReportConfig) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID) {
		return true
	}
	return false
}

func (repository *Repository) GetAllReportEntry(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportEntryRequest,
) (result []model.ReportEntry, err error) {
	result = make([]model.ReportEntry, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_entries.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "report_entries.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, "report_entries.class_subject_id = ?")
			args = append(args, request.ClassSubjectID)
		}
		if request.SequenceID > 0 {
			where = helpers.AppendWhereClause(where, "report_entries.sequence_id = ?")
			args = append(args, request.SequenceID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "report_entries.unit_id = ?")
			args = append(args, request.UnitID)
		}
		if request.SemesterID > 0 {
			where = helpers.AppendWhereClause(where, "university_units.semester_id = ?")
			args = append(args, request.SemesterID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "report_entries.student_id = ?")
			args = append(args, request.StudentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_entries.id AS TEXT) = ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			report_configs.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ? OR
			highschool_sequences.name ILIKE ? OR
			highschool_sequences.description ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ? OR
			university_semesters.name ILIKE ? OR
			university_semesters.description ILIKE ? OR
			students.uid ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.Domain").
		Preload("Unit.Domain.Department").
		Preload("Unit.Level").
		Preload("Unit.Semester").
		Preload("Student.User").
		Preload("Student.User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT report_entries.* 
				FROM report_entries 
				LEFT JOIN schools ON report_entries.school_id = schools.id
				LEFT JOIN report_configs ON report_entries.year_id = report_configs.id
				LEFT JOIN exam_types ON report_entries.type_id = exam_types.id
				LEFT JOIN highschool_class_subjects ON report_entries.class_subject_id = highschool_class_subjects.id
				LEFT JOIN highschool_sequences ON report_entries.sequence_id = highschool_sequences.id
				LEFT JOIN university_units ON report_entries.unit_id = university_units.id
				LEFT JOIN students ON report_entries.student_id = students.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id
				LEFT JOIN university_semesters ON university_units.semester_id = university_semesters.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllReportGrade(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportGradeRequest,
) (result []model.ReportGrade, err error) {
	result = make([]model.ReportGrade, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_grades.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_grades.id AS TEXT) = ? OR
			report_grades.name ILIKE ? OR
			report_grades.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT report_grades.* 
				FROM report_grades 
				LEFT JOIN schools ON report_grades.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllReportConfig(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportConfigRequest,
) (result []model.ReportConfig, err error) {
	result = make([]model.ReportConfig, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_configs.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_configs.id AS TEXT) = ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT report_configs.* 
				FROM report_configs 
				LEFT JOIN schools ON report_configs.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
