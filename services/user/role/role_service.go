package role

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/user/role/data"
	"api/services/user/role/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "role"
const DEFAULT_ERROR_MESSAGE = "interact with role model"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.RoleRequest) (result *model.Role, errCode int, err error) {
	// Format item
	item := &model.Role{
		Name:        request.Name,
		Feature:     request.Feature,
		Description: request.Description,
	}

	// Create
	result, err = service.Repository.Create(item)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage(MODEL_NAME)
				return
			}
			if pgState == constants.PG_ERROR_CONSTRAINT_COLUMN {
				errCode = http.StatusConflict
				err = constants.Http409ConflictErrorMessage()
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.RoleRequest) (result *model.Role, errCode int, err error) {
	// Check unique
	foundRole, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundRole == nil || foundRole.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format item
	item := &model.Role{
		Name:        request.Name,
		Feature:     request.Feature,
		Description: request.Description,
	}

	// Check if there is some role with the same name
	foundRoleToUpdate, err := service.Repository.GetByName(item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundRoleToUpdate != nil && foundRoleToUpdate.ID > 0 && foundRoleToUpdate.Name != foundRole.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateByID(id, item)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_CONSTRAINT_COLUMN {
				errCode = http.StatusConflict
				err = constants.Http409ConflictErrorMessage()
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
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

func (service *Service) GetByID(inputJwtToken *types.JwtToken, id int64) (result *model.Role, errCode int, err error) {
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
	return
}

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Role, errCode int, err error) {
	result, err = service.Repository.GetAll(inputJwtToken, filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
