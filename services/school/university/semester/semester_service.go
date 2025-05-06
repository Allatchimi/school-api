package semester

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/university/semester/model"
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

// Create new semester
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversitySemester) (result *model.UniversitySemester, errCode int, err error) {
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

	// Check unique
	foundItem, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get semester by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundItem, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Semester")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create semester from database")
		return
	}
	return
}

// Update semester
func (service *Service) Update(inputJwtToken *types.JwtToken, semesterID int64, item *model.UniversitySemester) (result *model.UniversitySemester, errCode int, err error) {
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

	// Check if semester exists
	foundItem, err := service.Repository.GetById(semesterID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get semester by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != semesterID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Semester")
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
		err = constants.Http500ErrorMessage("get semester by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Semester")
		return
	}

	// Update semester
	result, err = service.Repository.Update(semesterID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update semester from database")
		return
	}
	return
}

// Delete semester with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, semesterID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(semesterID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete semester from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Semester")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple semester from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Semester selection")
		return
	}
	return
}

// Get Returns semester with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, semesterID int64) (result *model.UniversitySemester, errCode int, err error) {
	result, err = service.Repository.GetById(semesterID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get semester by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Semester")
		return
	}
	return
}

// GetAll Returns all semesters with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversitySemester, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get semesters from database")
	}
	return
}
