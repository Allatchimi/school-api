package level

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/university/level/model"
)

type Service struct {
	Repository       *Repository
	SchoolRepository *school.Repository
}

func NewService(repository *Repository, SchoolRepository *school.Repository) *Service {
	return &Service{
		Repository:       repository,
		SchoolRepository: SchoolRepository,
	}
}

// Create new level
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversityLevel) (result *model.UniversityLevel, errCode int, err error) {
	// Check if the school type is university
	foundSchool, err := service.SchoolRepository.GetById(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_UNIVERSITY {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Level")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create level from database")
		return
	}
	return
}

// Update level
func (service *Service) Update(inputJwtToken *types.JwtToken, levelID int64, item *model.UniversityLevel) (result *model.UniversityLevel, errCode int, err error) {
	// Check if the school type is university
	foundSchool, err := service.SchoolRepository.GetById(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_UNIVERSITY {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if level exists
	foundItem, err := service.Repository.GetById(levelID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != levelID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Level")
		return
	}
	// Check if the school type is university
	if foundItem.School.Type != constants.SCHOOL_TYPE_UNIVERSITY {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Level")
		return
	}

	// Update level
	result, err = service.Repository.Update(levelID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update level from database")
		return
	}
	return
}

// Delete level with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, levelID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(levelID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete level from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Level")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple level from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Level selection")
		return
	}
	return
}

// Get Returns level with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, levelID int64) (result *model.UniversityLevel, errCode int, err error) {
	result, err = service.Repository.GetById(levelID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Level")
		return
	}
	return
}

// GetAll Returns all levels with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityLevel, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get levels from database")
	}
	return
}
