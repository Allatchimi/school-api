package domain

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/university/domain/model"
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

// Create new domain
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversityDomain) (result *model.UniversityDomain, errCode int, err error) {
	// Check if the school type is university
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
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

	// Check if the new one exists
	foundNewDomain, err := service.Repository.GetBySchoolIDDepartmentIDName(item.SchoolID, item.DepartmentID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by user school ids from database")
		return
	}
	if foundNewDomain != nil && foundNewDomain.SchoolID == item.SchoolID && foundNewDomain.DepartmentID == item.DepartmentID && foundNewDomain.Name == item.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Domain")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create domain from database")
		return
	}
	return
}

// Update domain
func (service *Service) Update(inputJwtToken *types.JwtToken, domainID int64, item *model.UniversityDomain) (result *model.UniversityDomain, errCode int, err error) {
	// Check if the school type is university
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
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

	// Check if domain exists
	foundItem, err := service.Repository.GetById(domainID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != domainID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Domain")
		return
	}
	// Check if the school type is university
	if foundItem.School.Type != constants.SCHOOL_TYPE_UNIVERSITY {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewDomain, err := service.Repository.GetBySchoolIDDepartmentIDName(item.SchoolID, item.DepartmentID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by user school ids from database")
		return
	}
	if foundNewDomain != nil && foundNewDomain.SchoolID == item.SchoolID && foundNewDomain.DepartmentID == item.DepartmentID && foundNewDomain.Name == item.Name {
		if !(foundItem.SchoolID == foundNewDomain.SchoolID && foundItem.DepartmentID == foundNewDomain.DepartmentID && foundItem.Name == foundNewDomain.Name) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage("Department")
			return
		}
	}

	// Update domain
	result, err = service.Repository.Update(domainID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update domain from database")
		return
	}
	return
}

// Delete domain with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, domainID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(domainID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete domain from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Domain")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple domain from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Domain selection")
		return
	}
	return
}

// Get Returns domain with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, domainID int64) (result *model.UniversityDomain, errCode int, err error) {
	result, err = service.Repository.GetById(domainID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Domain")
		return
	}
	return
}

// GetAll Returns all domains with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityDomain, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domains from database")
	}
	return
}
