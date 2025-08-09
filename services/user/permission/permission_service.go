package permission

import (
	"net/http"
	"slices"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
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
	ctxData *types.ContextData,
	roleID int64,
	request *data.UpdatePermissionRequest,
) (result *model.Permission, errCode int, err error) {
	// Check table name
	if request.TableName != constants.RESOURCE_TABLE_ALL &&
		!slices.Contains(constants.RESOURCE_TABLE_LIST, request.TableName) {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("table name")
		return
	}

	// Format item
	item := &model.Permission{
		RoleID:    roleID,
		TableName: request.TableName,
		Create:    request.Create,
		Read:      request.Read,
		Update:    request.Update,
		Delete:    request.Delete,
	}

	// Update all
	if item.TableName == constants.RESOURCE_TABLE_ALL {
		// Delete all by roleID
		_, errDelete := service.Repository.DeleteByRoleID(item.RoleID)
		if errDelete != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}

		// Create
		result, err = service.Repository.Create(item)
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
		}
		return
	}

	// Check unique
	foundPermission, err := service.Repository.GetByRoleIDTableName(item.RoleID, item.TableName)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Check if there is one permission on all tables
	foundPermissionAll, errFound := service.Repository.GetByRoleIDTableName(item.RoleID, constants.RESOURCE_TABLE_ALL)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundPermissionAll != nil && foundPermissionAll.ID > 0 {
		// Delete all by roleID
		_, errDelete := service.Repository.DeleteByRoleID(item.RoleID)
		if errDelete != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}

		// Create all
		service.Repository.CreateMultiple(constants.RESOURCE_TABLE_LIST, foundPermissionAll)

		// Update specific
		result, err = service.Repository.UpdateByRoleIDTableName(
			item.RoleID, item.TableName, item,
		)
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
		}
		return
	}

	if foundPermission != nil && foundPermission.ID > 0 {
		// Update
		result, err = service.Repository.UpdateByRoleIDTableName(
			item.RoleID, item.TableName, item,
		)
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
		}
		return
	}

	// Create
	result, err = service.Repository.Create(item)
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
	}
	return
}

func (service *Service) Delete(ctxData *types.ContextData, id int64) (affectedRows int64, errCode int, err error) {
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

func (service *Service) DeleteMultiple(ctxData *types.ContextData, list []int64) (affectedRows int64, errCode int, err error) {
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
	ctxData *types.ContextData,
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
