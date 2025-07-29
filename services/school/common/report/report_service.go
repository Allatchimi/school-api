package report

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/report/data"
	"api/services/school/common/report/model"
	"api/services/school/common/school"
)

type Service struct {
	Repository    *Repository
	SchoolService *school.Service
}

func NewService(
	repository *Repository,
	schoolService *school.Service,
) *Service {
	return &Service{
		Repository:    repository,
		SchoolService: schoolService,
	}
}

const MODEL_NAME = "report"
const DEFAULT_ERROR_MESSAGE = "interact with report model"

func (service *Service) CreateEntry(
	ctxData *types.ContextData,
	request *data.ReportEntryRequest,
) (result *model.ReportEntry, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// TODO
	return
}

func (service *Service) CreateGrade(
	ctxData *types.ContextData,
	request *data.ReportGradeRequest,
) (result *model.ReportGrade, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format
	item := &model.ReportGrade{
		SchoolID: newRequest.SchoolID,

		Type:           newRequest.Type,
		Name:           newRequest.Name,
		Description:    newRequest.Description,
		Minimum:        newRequest.Minimum,
		Maximum:        newRequest.Maximum,
		IncludeMinimum: newRequest.IncludeMinimum,
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObjectReportGrade(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsReportGrade(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Create
	result, err = service.Repository.CreateReportGrade(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateConfig(
	ctxData *types.ContextData,
	request *data.ReportConfigRequest,
) (result *model.ReportConfig, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format
	item := &model.ReportConfig{
		SchoolID: newRequest.SchoolID,

		NotationAverage:               newRequest.NotationAverage,
		NotationReport:                newRequest.NotationReport,
		MinimumRequiredValueToPromote: newRequest.MinimumRequiredValueToPromote,
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObjectReportConfig(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsReportConfig(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Create
	result, err = service.Repository.CreateReportConfig(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateGrade(
	ctxData *types.ContextData,
	id int64,
	request *data.ReportGradeRequest,
) (result *model.ReportGrade, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ReportGrade
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetReportGradeByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetReportGradeByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format request
	item := &model.ReportGrade{
		SchoolID: newRequest.SchoolID,

		Type:           newRequest.Type,
		Name:           newRequest.Name,
		Description:    newRequest.Description,
		Minimum:        newRequest.Minimum,
		Maximum:        newRequest.Maximum,
		IncludeMinimum: newRequest.IncludeMinimum,
		IncludeMaximum: newRequest.IncludeMaximum,
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObjectReportGrade(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsReportGrade(foundUnique, item) &&
		!service.Repository.AreSameUniqueObjectsReportGrade(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateReportGradeByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateConfig(
	ctxData *types.ContextData,
	id int64,
	request *data.ReportConfigRequest,
) (result *model.ReportConfig, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ReportConfig
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetReportConfigByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetReportConfigByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format request
	item := &model.ReportConfig{
		SchoolID: newRequest.SchoolID,

		NotationAverage:               newRequest.NotationAverage,
		NotationReport:                newRequest.NotationReport,
		MinimumRequiredValueToPromote: newRequest.MinimumRequiredValueToPromote,
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObjectReportConfig(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsReportConfig(foundUnique, item) &&
		!service.Repository.AreSameUniqueObjectsReportConfig(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateReportConfigByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) DeleteEntry(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ReportEntry
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetReportEntryByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetReportEntryByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteReportEntryByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteGrade(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ReportGrade
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetReportGradeByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetReportGradeByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteReportGradeByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteConfig(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ReportConfig
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetReportConfigByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetReportConfigByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteReportConfigByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultipleEntry(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleReportEntryByID(list, ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultipleGrade(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleReportGradeByID(list, ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultipleConfig(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleReportConfigByID(list, ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetEntry(
	ctxData *types.ContextData,
	id int64,
) (result *model.ReportEntry, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetReportEntryByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetReportEntryByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetGrade(
	ctxData *types.ContextData,
	id int64,
) (result *model.ReportGrade, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetReportGradeByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetReportGradeByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetConfig(
	ctxData *types.ContextData,
	id int64,
) (result *model.ReportConfig, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetReportConfigByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetReportConfigByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetAllEntry(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllReportEntryRequest,
) (result []model.ReportEntry, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllReportEntry(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllGrade(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllReportGradeRequest,
) (result []model.ReportGrade, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllReportGrade(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllConfig(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllReportConfigRequest,
) (result []model.ReportConfig, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllReportConfig(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllTable(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllReportTableRequest,
) (result []model.ReportTable, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get school
	if newRequest.SchoolID > 0 {
		foundSchool, errFound := service.SchoolService.Repository.GetByID(newRequest.SchoolID)
		if errFound != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		}
		if foundSchool == nil || foundSchool.ID < 1 {
			errCode = http.StatusNotFound
			err = constants.Http404ErrorMessage(MODEL_NAME)
			return
		}

		// Get
		switch foundSchool.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			result, err = service.Repository.GetAllReportTableHighschool(filter, pagination, &newRequest)
			if err != nil {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			result, err = service.Repository.GetAllReportTableUniversity(filter, pagination, &newRequest)
			if err != nil {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			}
		}
		return
	}

	// Get
	result, err = service.Repository.GetAllReportTable(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
