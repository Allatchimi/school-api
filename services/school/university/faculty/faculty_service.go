package faculty

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/university/faculty/model"
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

// Create new faculty
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversityFaculty) (result *model.UniversityFaculty, errCode int, err error) {
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
		err = constants.Http500ErrorMessage("get faculty by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Faculty")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create faculty from database")
		return
	}
	return
}

// Update faculty
func (service *Service) Update(inputJwtToken *types.JwtToken, domainID int64, item *model.UniversityFaculty) (result *model.UniversityFaculty, errCode int, err error) {
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

	// Check if faculty exists
	foundItem, err := service.Repository.GetById(domainID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculty by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != domainID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Faculty")
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
		err = constants.Http500ErrorMessage("get faculty by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Faculty")
		return
	}

	// Update faculty
	result, err = service.Repository.Update(domainID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update faculty from database")
		return
	}
	return
}

// Delete faculty with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, domainID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(domainID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete faculty from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Faculty")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple faculty from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Faculty selection")
		return
	}
	return
}

// Get Returns faculty with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, domainID int64) (result *model.UniversityFaculty, errCode int, err error) {
	result, err = service.Repository.GetById(domainID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculty by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Faculty")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityFaculty, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
