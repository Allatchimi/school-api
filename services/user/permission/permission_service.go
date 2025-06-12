package permission

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/permission/data"
	"api/services/user/permission/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "permission"
const DEFAULT_ERROR_MESSAGE = "interact with permission model"

func (service *Service) Update(
	inputJwtToken *types.JwtToken,
	roleID int64,
	tableName string,
	item *model.Permission,
) (result *model.Permission, errCode int, err error) {
	// Check unique
	foundPermission, err := service.Repository.GetByRoleIDTableName(roleID, tableName)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundPermission == nil || foundPermission.RoleID != roleID {
		// Create new ones
		result, err = service.Repository.Create(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		}
		return
	}

	// Update now
	result, err = service.Repository.UpdateByID(
		roleID, tableName, item,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, roleID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteByID(roleID)
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

func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleByID(list)
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

func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Permission, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
