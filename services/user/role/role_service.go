package role

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/user/role/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create Creates role
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Role) (result *model.Role, errCode int, err error) {
	result, err = service.Repository.Create(item)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage("role")
				return
			}
		}

		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create role from database")
		return
	}
	return
}

// Update Updates role
func (service *Service) Update(inputJwtToken *types.JwtToken, roleID int64, item *model.Role) (result *model.Role, errCode int, err error) {
	// Check if the role exists
	foundRole, err := service.Repository.GetByID(roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("find role from database")
		return
	}
	if foundRole == nil || foundRole.ID != roleID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role")
		return
	}

	// Check if there is some role with the same name
	foundRoleToUpdate, err := service.Repository.GetByName(item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("find role from database")
		return
	}
	if foundRoleToUpdate != nil && foundRoleToUpdate.ID > 0 && foundRoleToUpdate.Name != foundRole.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("role")
		return
	}

	// Update now
	result, err = service.Repository.Update(roleID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update role from database")
		return
	}
	return
}

// Delete Deletes role
func (service *Service) Delete(inputJwtToken *types.JwtToken, roleID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete role from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple role from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role selection")
		return
	}
	return
}

// Get Returns role
func (service *Service) GetByID(inputJwtToken *types.JwtToken, roleID int64) (result *model.Role, errCode int, err error) {
	result, err = service.Repository.GetByID(roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get role by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role")
		return
	}
	return
}

// GetAll Returns role list
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Role, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get roles from database")
	}
	return
}
