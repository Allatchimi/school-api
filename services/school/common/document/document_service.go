package document

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	common_svc_permission "api/services/common"
	"api/services/school/common/document/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "document"
const DEFAULT_ERROR_MESSAGE = "interact with document model"

func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Document) (result *model.Document, errCode int, err error) {
	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.RoleID,
		inputJwtToken.UserID,
		item.SchoolID,
		item.YearID,
		item.ClassSubjectID,
		item.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Insert document
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	// Check if the user can access
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.RoleID,
		inputJwtToken.UserID,
		foundItem.SchoolID,
		foundItem.YearID,
		foundItem.ClassSubjectID,
		foundItem.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Perform delete
	affectedRows, err = service.Repository.DeleteByID(id)
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Document, errCode int, err error) {
	result, err = service.Repository.GetByID(id)
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

	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.RoleID,
		inputJwtToken.UserID,
		result.SchoolID,
		result.YearID,
		result.ClassSubjectID,
		result.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}
	return
}

func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	clause *types.FilterlSchoolYearUnitClassSubjectRequest,
) (result []model.Document, errCode int, err error) {
	var schoolID, yearID, classSubjectID, unitID int64
	if clause != nil {
		schoolID = clause.SchoolID
		yearID = clause.YearID
		classSubjectID = clause.ClassSubjectID
		unitID = clause.UnitID
	}
	result, err = service.Repository.GetAll(filter, pagination, schoolID, yearID, classSubjectID, unitID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
