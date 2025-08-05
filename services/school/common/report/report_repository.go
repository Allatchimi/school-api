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

func (repository *Repository) CreateReportEntry(item *model.ReportEntry) (result *model.ReportEntry, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.ReportEntry{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

type aggregationKey struct {
	StudentID      int64
	SchoolID       int64
	UnitID         int64 // universitaire
	ClassSubjectID int64 // lycée
	YearID         int64
	SequenceID     int64
}

func (repository *Repository) CreateReportGrade(item *model.ReportGrade) (result *model.ReportGrade, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.ReportGrade{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateReportCorrespondence(item *model.ReportCorrespondence) (result *model.ReportCorrespondence, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.ReportCorrespondence{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateReportConfig(item *model.ReportConfig) (result *model.ReportConfig, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.ReportConfig{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateReportAverage(item *model.ReportAverage) (result *model.ReportAverage, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.ReportAverage{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateReportTable(item *model.ReportTable) (result *model.ReportTable, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.ReportTable{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateReportGradeByID(id int64, item *model.ReportGrade) (result *model.ReportGrade, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,

		"type":            item.Type,
		"name":            item.Name,
		"description":     item.Description,
		"minimum":         item.Minimum,
		"maximum":         item.Maximum,
		"include_minimum": item.IncludeMinimum,
		"include_maximum": item.IncludeMaximum,
	}
	err = repository.Db.
		Model(&model.ReportGrade{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.ReportGrade{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateReportCorrespondenceByID(id int64, item *model.ReportCorrespondence) (result *model.ReportCorrespondence, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,

		"minimum":         item.Minimum,
		"maximum":         item.Maximum,
		"include_minimum": item.IncludeMinimum,
		"include_maximum": item.IncludeMaximum,
		"new_score":       item.NewScore,
	}
	err = repository.Db.
		Model(&model.ReportCorrespondence{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.ReportCorrespondence{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateReportConfigByID(id int64, item *model.ReportConfig) (result *model.ReportConfig, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":               item.SchoolID,
		"report_grade_to_fail_id": item.ReportGradeToFailID,

		"notation_average":                  item.NotationAverage,
		"notation_report":                   item.NotationReport,
		"minimum_required_score_to_promote": item.MinimumRequiredScoreToPromote,
		"only_failed_exams":                 item.OnlyFailedExams,
	}
	err = repository.Db.
		Model(&model.ReportConfig{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.ReportConfig{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateReportTableByID(id int64, item *model.ReportTable) (result *model.ReportTable, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":       item.SchoolID,
		"year_id":         item.YearID,
		"class_id":        nil,
		"level_domain_id": nil,

		"period_type": item.PeriodType,
		"period_name": item.PeriodName,
		"status":      item.Status,
	}
	if item.ClassID > 0 {
		fields["class_id"] = item.ClassID
	}
	if item.LevelDomainID > 0 {
		fields["level_domain_id"] = item.LevelDomainID
	}
	err = repository.Db.
		Model(&model.ReportTable{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.ReportTable{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteReportEntryByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ReportEntry{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteReportGradeByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ReportGrade{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteReportCorrespondenceByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ReportCorrespondence{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteReportConfigByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ReportConfig{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleReportEntryByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ReportEntry{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleReportGradeByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ReportGrade{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleReportCorrespondenceByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ReportCorrespondence{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleReportConfigByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ReportConfig{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetReportEntryByID(id int64) (*model.ReportEntry, error) {
	result := &model.ReportEntry{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Class.School").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.School").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Unit.Semester").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetReportGradeByID(id int64) (*model.ReportGrade, error) {
	result := &model.ReportGrade{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetReportCorrespondenceByID(id int64) (*model.ReportCorrespondence, error) {
	result := &model.ReportCorrespondence{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetReportConfigByID(id int64) (*model.ReportConfig, error) {
	result := &model.ReportConfig{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetReportEntryByIDSchoolID(id int64, schoolID int64) (*model.ReportEntry, error) {
	result := &model.ReportEntry{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Class.School").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.School").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Unit.Semester").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetReportGradeByIDSchoolID(id int64, schoolID int64) (*model.ReportGrade, error) {
	result := &model.ReportGrade{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetReportCorrespondenceByIDSchoolID(id int64, schoolID int64) (*model.ReportCorrespondence, error) {
	result := &model.ReportCorrespondence{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetReportConfigByIDSchoolID(id int64, schoolID int64) (*model.ReportConfig, error) {
	result := &model.ReportConfig{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetReportConfigBySchoolID(schoolID int64) (*model.ReportConfig, error) {
	result := &model.ReportConfig{}
	return result, repository.Db.Preload(clause.Associations).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectReportEntry(item *model.ReportEntry) (*model.ReportEntry, error) {
	result := &model.ReportEntry{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportEntry{
		SchoolID:       item.SchoolID,
		YearID:         item.YearID,
		ClassSubjectID: item.ClassSubjectID,
		SequenceID:     item.SequenceID,
		UnitID:         item.UnitID,
		StudentID:      item.StudentID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsReportEntry(item1 *model.ReportEntry, item2 *model.ReportEntry) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.ClassSubjectID == item2.ClassSubjectID &&
			item1.SequenceID == item2.SequenceID &&
			item1.UnitID == item2.UnitID &&
			item1.StudentID == item2.StudentID) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectReportGrade(item *model.ReportGrade) (*model.ReportGrade, error) {
	result := &model.ReportGrade{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportGrade{
		SchoolID: item.SchoolID,
		Type:     item.Type,
		Name:     item.Name,
		Minimum:  item.Maximum,
		Maximum:  item.Maximum,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsReportGrade(item1 *model.ReportGrade, item2 *model.ReportGrade) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Type == item2.Type &&
			item1.Name == item2.Name &&
			item1.Minimum == item2.Minimum &&
			item1.Maximum == item2.Maximum) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectReportCorrespondence(item *model.ReportCorrespondence) (*model.ReportCorrespondence, error) {
	result := &model.ReportCorrespondence{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportCorrespondence{
		SchoolID: item.SchoolID,
		Minimum:  item.Maximum,
		Maximum:  item.Maximum,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsReportCorrespondence(item1 *model.ReportCorrespondence, item2 *model.ReportCorrespondence) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Minimum == item2.Minimum &&
			item1.Maximum == item2.Maximum) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectReportConfig(item *model.ReportConfig) (*model.ReportConfig, error) {
	result := &model.ReportConfig{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportConfig{
		SchoolID: item.SchoolID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsReportConfig(item1 *model.ReportConfig, item2 *model.ReportConfig) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectReportAverage(item *model.ReportAverage) (*model.ReportAverage, error) {
	result := &model.ReportAverage{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportAverage{
		SchoolID:      item.SchoolID,
		YearID:        item.YearID,
		ClassID:       item.ClassID,
		LevelDomainID: item.LevelDomainID,
		PeriodType:    item.PeriodType,
		PeriodName:    item.PeriodName,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsReportAverage(item1 *model.ReportAverage, item2 *model.ReportAverage) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.ClassID == item2.ClassID &&
			item1.LevelDomainID == item2.LevelDomainID &&
			item1.PeriodType == item2.PeriodType &&
			item1.PeriodName == item2.PeriodName) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectReportTable(item *model.ReportTable) (*model.ReportTable, error) {
	result := &model.ReportTable{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ReportTable{
		SchoolID:      item.SchoolID,
		YearID:        item.YearID,
		ClassID:       item.ClassID,
		LevelDomainID: item.LevelDomainID,
		PeriodType:    item.PeriodType,
		PeriodName:    item.PeriodName,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsReportTable(item1 *model.ReportTable, item2 *model.ReportTable) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.ClassID == item2.ClassID &&
			item1.LevelDomainID == item2.LevelDomainID &&
			item1.PeriodType == item2.PeriodType &&
			item1.PeriodName == item2.PeriodName) {
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
			years.name ILIKE ? OR
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
		Preload("ClassSubject.Class.School").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.School").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Unit.Semester").
		Preload("Student.User").
		Preload("Student.User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT report_entries.* 
				FROM report_entries 
				LEFT JOIN schools ON report_entries.school_id = schools.id
				LEFT JOIN years ON report_entries.year_id = years.id
				LEFT JOIN highschool_class_subjects ON report_entries.class_subject_id = highschool_class_subjects.id
				LEFT JOIN highschool_sequences ON report_entries.sequence_id = highschool_sequences.id
				LEFT JOIN university_units ON report_entries.unit_id = university_units.id
				LEFT JOIN students ON report_entries.student_id = students.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id
				LEFT JOIN university_semesters ON university_units.semester_id = university_semesters.id
				
				LEFT JOIN teacher_class_subject_units ON courses.class_subject_id = teacher_class_subject_units.class_subject_id AND courses.unit_id = teacher_class_subject_units.unit_id
				LEFT JOIN student_enrolls ON highschool_class_subjects.class_id = student_enrolls.class_id AND university_units.level_domain_id = student_enrolls.level_domain_id
				LEFT JOIN parent_students ON student_enrolls.student_id = parent_students.student_id`,
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
		if len(request.Type) > 0 {
			where = helpers.AppendWhereClause(where, "report_grades.type = ?")
			args = append(args, request.Type)
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
			report_grades.type ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT report_grades.* 
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

func (repository *Repository) GetAllReportCorrespondence(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportCorrespondenceRequest,
) (result []model.ReportCorrespondence, err error) {
	result = make([]model.ReportCorrespondence, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_correspondences.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_correspondences.id AS TEXT) = ? OR
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
				`SELECT DISTINCT report_correspondences.* 
				FROM report_correspondences 
				LEFT JOIN schools ON report_correspondences.school_id = schools.id`,
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
				`SELECT DISTINCT report_configs.* 
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

func (repository *Repository) GetAllReportAverage(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportAverageRequest,
) (result []model.ReportAverage, err error) {
	result = make([]model.ReportAverage, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_averages.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "report_averages.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "report_averages.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "report_averages.level_domain = ?")
			args = append(args, request.LevelDomainID)
		}
		if len(request.PeriodType) > 0 {
			where = helpers.AppendWhereClause(where, "report_averages.period_type = ?")
			args = append(args, request.PeriodType)
		}
		if len(request.PeriodName) > 0 {
			where = helpers.AppendWhereClause(where, "report_averages.period_name = ?")
			args = append(args, request.PeriodName)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "teacher_class_subject_units.teacher_id = ?")
			args = append(args, request.TeacherID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.parent_id = ?")
			args = append(args, request.ParentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_averages.id AS TEXT) = ? OR
			report_averages.period_type ILIKE ? OR
			report_averages.period_name ILIKE ? OR
			report_averages.status ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			university_levels.name ILIKE ? OR
			university_levels.description ILIKE ? OR
			university_domains.name ILIKE ? OR
			university_domains.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.School").
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.School").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT report_averages.* 
				FROM report_averages 
				LEFT JOIN schools ON report_averages.school_id = schools.id
				LEFT JOIN years ON report_averages.year_id = years.id
				LEFT JOIN highschool_classes ON report_averages.class_id = highschool_classes.id
				LEFT JOIN university_level_domains ON report_averages.level_domain_id = university_level_domains.id
				LEFT JOIN university_levels ON university_level_domains.level_id = university_levels.id
				LEFT JOIN university_domains ON university_level_domains.domain_id = university_domains.id
				LEFT JOIN highschool_class_subjects ON report_averages.class_id = highschool_class_subjects.class_id
				LEFT JOIN university_units ON report_averages.level_domain_id = university_units.level_domain_id
				
				LEFT JOIN teacher_class_subject_units ON report_averages.school_id = teacher_class_subject_units.school_id
				AND (
				(teacher_class_subject_units.class_subject_id IS NOT NULL AND highschool_class_subjects.id = teacher_class_subject_units.class_subject_id)
				OR
				(teacher_class_subject_units.unit_id IS NOT NULL AND university_units.id = teacher_class_subject_units.unit_id)
				)
				LEFT JOIN student_enrolls ON report_averages.school_id = student_enrolls.school_id
				AND (
				(student_enrolls.class_id IS NOT NULL AND report_averages.class_id = student_enrolls.class_id)
				OR
				(student_enrolls.level_domain_id IS NOT NULL AND report_averages.level_domain_id = student_enrolls.level_domain_id)
				)
				LEFT JOIN parent_students ON student_enrolls.student_id = parent_students.student_id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllReportTable(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportTableRequest,
) (result []model.ReportTable, err error) {
	result = make([]model.ReportTable, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.level_domain = ?")
			args = append(args, request.LevelDomainID)
		}
		if len(request.PeriodType) > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.period_type = ?")
			args = append(args, request.PeriodType)
		}
		if len(request.PeriodName) > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.period_name = ?")
			args = append(args, request.PeriodName)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "teacher_class_subject_units.teacher_id = ?")
			args = append(args, request.TeacherID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.parent_id = ?")
			args = append(args, request.ParentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_tables.id AS TEXT) = ? OR
			report_tables.period_type ILIKE ? OR
			report_tables.period_name ILIKE ? OR
			report_tables.status ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			university_levels.name ILIKE ? OR
			university_levels.description ILIKE ? OR
			university_domains.name ILIKE ? OR
			university_domains.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.School").
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.School").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT report_tables.* 
				FROM report_tables 
				LEFT JOIN schools ON report_tables.school_id = schools.id
				LEFT JOIN years ON report_tables.year_id = years.id
				LEFT JOIN highschool_classes ON report_tables.class_id = highschool_classes.id
				LEFT JOIN university_level_domains ON report_tables.level_domain_id = university_level_domains.id
				LEFT JOIN university_levels ON university_level_domains.level_id = university_levels.id
				LEFT JOIN university_domains ON university_level_domains.domain_id = university_domains.id
				LEFT JOIN highschool_class_subjects ON report_tables.class_id = highschool_class_subjects.class_id
				LEFT JOIN university_units ON report_tables.level_domain_id = university_units.level_domain_id
				
				
				LEFT JOIN teacher_class_subject_units ON report_tables.school_id = teacher_class_subject_units.school_id
				AND (
				(teacher_class_subject_units.class_subject_id IS NOT NULL AND highschool_class_subjects.id = teacher_class_subject_units.class_subject_id)
				OR
				(teacher_class_subject_units.unit_id IS NOT NULL AND university_units.id = teacher_class_subject_units.unit_id)
				)
				LEFT JOIN student_enrolls ON report_tables.school_id = student_enrolls.school_id
				AND (
				(student_enrolls.class_id IS NOT NULL AND report_tables.class_id = student_enrolls.class_id)
				OR
				(student_enrolls.level_domain_id IS NOT NULL AND report_tables.level_domain_id = student_enrolls.level_domain_id)
				)
				LEFT JOIN parent_students ON student_enrolls.student_id = parent_students.student_id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllReportTableHighschool(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportTableRequest,
) (result []model.ReportTable, err error) {
	result = make([]model.ReportTable, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.class_id = ?")
			args = append(args, request.ClassID)
		}
		if len(request.PeriodType) > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.period_type = ?")
			args = append(args, request.PeriodType)
		}
		if len(request.PeriodName) > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.period_name = ?")
			args = append(args, request.PeriodName)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "teacher_class_subject_units.teacher_id = ?")
			args = append(args, request.TeacherID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.parent_id = ?")
			args = append(args, request.ParentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_tables.id AS TEXT) = ? OR
			report_tables.period_type ILIKE ? OR
			report_tables.period_name ILIKE ? OR
			report_tables.status ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			students.uid ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.School").
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.School").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT report_tables.* 
				FROM report_tables 
				LEFT JOIN schools ON report_tables.school_id = schools.id
				LEFT JOIN years ON report_tables.year_id = years.id
				LEFT JOIN highschool_classes ON report_tables.class_id = highschool_classes.id
				LEFT JOIN highschool_class_subjects ON report_tables.class_id = highschool_class_subjects.class_id
				
				
				LEFT JOIN teacher_class_subject_units ON report_tables.school_id = teacher_class_subject_units.school_id
				AND (
				teacher_class_subject_units.class_subject_id IS NOT NULL AND highschool_class_subjects.id = teacher_class_subject_units.class_subject_id
				)
				LEFT JOIN student_enrolls ON report_tables.school_id = student_enrolls.school_id
				AND (
				student_enrolls.class_id IS NOT NULL AND report_tables.class_id = student_enrolls.class_id
				)
				LEFT JOIN parent_students ON student_enrolls.student_id = parent_students.student_id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllReportTableUniversity(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllReportTableRequest,
) (result []model.ReportTable, err error) {
	result = make([]model.ReportTable, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.class_id = ?")
			args = append(args, request.ClassID)
		}
		if len(request.PeriodType) > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.period_type = ?")
			args = append(args, request.PeriodType)
		}
		if len(request.PeriodName) > 0 {
			where = helpers.AppendWhereClause(where, "report_tables.period_name = ?")
			args = append(args, request.PeriodName)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "teacher_class_subject_units.teacher_id = ?")
			args = append(args, request.TeacherID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.parent_id = ?")
			args = append(args, request.ParentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(report_tables.id AS TEXT) = ? OR
			report_tables.period_type ILIKE ? OR
			report_tables.period_name ILIKE ? OR
			report_tables.status ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			university_levels.name ILIKE ? OR
			university_levels.description ILIKE ? OR
			university_domains.name ILIKE ? OR
			university_domains.description ILIKE ? OR
			students.uid ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.School").
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.School").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT report_tables.* 
				FROM report_tables 
				LEFT JOIN schools ON report_tables.school_id = schools.id
				LEFT JOIN years ON report_tables.year_id = years.id
				LEFT JOIN university_level_domains ON report_tables.level_domain_id = university_level_domains.id
				LEFT JOIN university_levels ON university_level_domains.level_id = university_levels.id
				LEFT JOIN university_domains ON university_level_domains.domain_id = university_domains.id
				LEFT JOIN university_units ON report_tables.level_domain_id = university_units.level_domain_id
				
				
				LEFT JOIN teacher_class_subject_units ON report_tables.school_id = teacher_class_subject_units.school_id
				AND (
				teacher_class_subject_units.unit_id IS NOT NULL AND university_units.id = teacher_class_subject_units.unit_id
				)
				LEFT JOIN student_enrolls ON report_tables.school_id = student_enrolls.school_id
				AND (
				student_enrolls.level_domain_id IS NOT NULL AND report_tables.level_domain_id = student_enrolls.level_domain_id
				)
				LEFT JOIN parent_students ON student_enrolls.student_id = parent_students.student_id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
