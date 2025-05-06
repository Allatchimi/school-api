package tu

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/university/tu/model"
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

// Create new teaching unit
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversityTeachingUnit) (result *model.UniversityTeachingUnit, errCode int, err error) {
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
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teaching unit by name from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teaching unit")
		return
	}

	// Insert teaching unit
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create teaching unit from database")
		return
	}
	return
}

// Update teaching unit
func (service *Service) Update(inputJwtToken *types.JwtToken, teachingUnitID int64, item *model.UniversityTeachingUnit) (result *model.UniversityTeachingUnit, errCode int, err error) {
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

	// Check if tu exists
	foundItem, err := service.Repository.GetById(teachingUnitID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teaching unit by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != teachingUnitID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teaching unit")
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
		err = constants.Http500ErrorMessage("get teaching unit by name from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teaching unit")
		return
	}

	// Update teaching unit
	result, err = service.Repository.Update(teachingUnitID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update teaching unit from database")
		return
	}
	return
}

// Delete teaching unit with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, teachingUnitID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(teachingUnitID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete teaching unit from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("teaching unit" + fmt.Sprint(teachingUnitID))
		return
	}
	return
}

// Get Returns teaching unit with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, teachingUnitID int64) (result *model.UniversityTeachingUnit, errCode int, err error) {
	result, err = service.Repository.GetById(teachingUnitID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teaching unit by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("teaching unit")
		return
	}
	return
}

// GetAll Returns all teaching units with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityTeachingUnit, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teaching units from database")
	}
	return
}
