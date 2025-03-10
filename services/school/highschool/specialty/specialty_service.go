package specialty

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/highschool/specialty/model"
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

// Create new specialty
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.HighschoolSpecialty) (result *model.HighschoolSpecialty, errCode int, err error) {
	// Check if the school type is highschool
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewSpecialty, err := service.Repository.GetBySchoolIDSectionIDName(item.SchoolID, item.SectionID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get specialty by user school ids from database")
		return
	}
	if foundNewSpecialty != nil && foundNewSpecialty.SchoolID == item.SchoolID && foundNewSpecialty.SectionID == item.SectionID && foundNewSpecialty.Name == item.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Specialty")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create specialty from database")
		return
	}
	return
}

// Update specialty
func (service *Service) Update(inputJwtToken *types.JwtToken, specialtyID int64, item *model.HighschoolSpecialty) (result *model.HighschoolSpecialty, errCode int, err error) {
	// Check if the school type is highschool
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if specialty exists
	foundItem, err := service.Repository.GetById(specialtyID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get specialty by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != specialtyID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Specialty")
		return
	}
	// Check if the school type is highschool
	if foundItem.School.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewSpecialty, err := service.Repository.GetBySchoolIDSectionIDName(item.SchoolID, item.SectionID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get specialty by user school ids from database")
		return
	}
	if foundNewSpecialty != nil && foundNewSpecialty.SchoolID == item.SchoolID && foundNewSpecialty.SectionID == item.SectionID && foundNewSpecialty.Name == item.Name {
		if !(foundItem.SchoolID == foundNewSpecialty.SchoolID && foundItem.SectionID == foundNewSpecialty.SectionID && foundItem.Name == foundNewSpecialty.Name) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage("Section")
			return
		}
	}

	// Update specialty
	result, err = service.Repository.Update(specialtyID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update specialty from database")
		return
	}
	return
}

// Delete specialty with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, specialtyID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(specialtyID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete specialty from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Specialty")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple specialty from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Specialty selection")
		return
	}
	return
}

// Get Returns specialty with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, specialtyID int64) (result *model.HighschoolSpecialty, errCode int, err error) {
	result, err = service.Repository.GetById(specialtyID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get specialty by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Specialty")
		return
	}
	return
}

// GetAll Returns all specialties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolSpecialty, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get specialties from database")
	}
	return
}
