package department

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/university/department/model"
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

// Create new department
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversityDepartment) (result *model.UniversityDepartment, errCode int, err error) {
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
		err = constants.Http500ErrorMessage("get department by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Department")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create department from database")
		return
	}
	return
}

// Update department
func (service *Service) Update(inputJwtToken *types.JwtToken, departmentID int64, item *model.UniversityDepartment) (result *model.UniversityDepartment, errCode int, err error) {
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

	// Check if department exists
	foundItem, err := service.Repository.GetById(departmentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != departmentID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
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
		err = constants.Http500ErrorMessage("get department by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Department")
		return
	}

	// Update department
	result, err = service.Repository.Update(departmentID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update department from database")
		return
	}
	return
}

// Delete department with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, departmentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(departmentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete department from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple department from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department selection")
		return
	}
	return
}

// Get Returns department with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, departmentID int64) (result *model.UniversityDepartment, errCode int, err error) {
	result, err = service.Repository.GetById(departmentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	return
}

// GetAll Returns all departments with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityDepartment, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get departments from database")
	}
	return
}
