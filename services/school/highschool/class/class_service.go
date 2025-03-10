package class

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/highschool/class/model"
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

// Create new class
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.HighschoolClass) (result *model.HighschoolClass, errCode int, err error) {
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
	foundNewClass, err := service.Repository.GetBySchoolIDSpecialtyIDName(item.SchoolID, item.SpecialtyID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by user school ids from database")
		return
	}
	if foundNewClass != nil && foundNewClass.SchoolID == item.SchoolID && foundNewClass.SpecialtyID == item.SpecialtyID && foundNewClass.Name == item.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Class")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create class from database")
		return
	}
	return
}

// Update class
func (service *Service) Update(inputJwtToken *types.JwtToken, classID int64, item *model.HighschoolClass) (result *model.HighschoolClass, errCode int, err error) {
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

	// Check if class exists
	foundItem, err := service.Repository.GetById(classID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != classID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Class")
		return
	}
	// Check if the school type is highschool
	if foundItem.School.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewClass, err := service.Repository.GetBySchoolIDSpecialtyIDName(item.SchoolID, item.SpecialtyID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by user school ids from database")
		return
	}
	if foundNewClass != nil && foundNewClass.SchoolID == item.SchoolID && foundNewClass.SpecialtyID == item.SpecialtyID && foundNewClass.Name == item.Name {
		if !(foundItem.SchoolID == foundNewClass.SchoolID && foundItem.SpecialtyID == foundNewClass.SpecialtyID && foundItem.Name == foundNewClass.Name) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage("Specialty")
			return
		}
	}

	// Update class
	result, err = service.Repository.Update(classID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update class from database")
		return
	}
	return
}

// Delete class with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, classID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(classID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete class from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Class")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple class from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Class selection")
		return
	}
	return
}

// Get Returns class with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, classID int64) (result *model.HighschoolClass, errCode int, err error) {
	result, err = service.Repository.GetById(classID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Class")
		return
	}
	return
}

// GetAll Returns all classes with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolClass, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get classes from database")
	}
	return
}
