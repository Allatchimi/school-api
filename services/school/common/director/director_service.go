package director

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/config"
	"api/services/school/common/director/model"
	"api/services/user/user"
)

type Service struct {
	Repository     *Repository
	UserRepository *user.Repository
}

func NewService(repository *Repository, userRepository *user.Repository) *Service {
	return &Service{
		Repository:     repository,
		UserRepository: userRepository,
	}
}

// Create new director
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Director) (result *model.Director, errCode int, err error) {
	// Check unique
	foundNewDirector, err := service.Repository.GetByUserSchoolIDs(item.UserID, item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get director by user school ids from database")
		return
	}
	if foundNewDirector != nil && foundNewDirector.UserID == item.UserID {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Director")
		return
	}

	// Check if the user exists and role is director
	foundUser, err := service.UserRepository.GetByID(item.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get user by id from database")
		return
	}
	if foundUser == nil || foundUser.ID != item.UserID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	if foundUser.Role.Name != config.Env.RoleDirector {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create director from database")
		return
	}
	return
}

// Update director
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.Director) (result *model.Director, errCode int, err error) {
	// Check if director exists
	foundItem, err := service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get director by name from database")
		return
	}
	if foundItem == nil || foundItem.UserID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Director")
		return
	}

	// Check unique
	foundNewDirector, err := service.Repository.GetByUserSchoolIDs(item.UserID, item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get director by user school ids from database")
		return
	}
	if foundNewDirector != nil && foundNewDirector.UserID == item.UserID {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Director")
		return
	}

	// Check if the user exists and role is director
	foundUser, err := service.UserRepository.GetByID(item.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get user by id from database")
		return
	}
	if foundUser == nil || foundUser.ID != item.UserID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	if foundUser.Role.Name != config.Env.RoleDirector {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Update director
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update director from database")
		return
	}
	return
}

// Delete director with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete director from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Director")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple director from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Director selection")
		return
	}
	return
}

// Get Returns director with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Director, errCode int, err error) {
	result, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get director by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Director")
		return
	}
	return
}

// GetAll Returns all directors with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Director, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get directors from database")
	}
	return
}
