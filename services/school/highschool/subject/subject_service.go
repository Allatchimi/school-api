package subject

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/highschool/subject/model"
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

// Create new subject
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.HighschoolSubject) (result *model.HighschoolSubject, errCode int, err error) {
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
	foundNewSubject, err := service.Repository.GetBySchoolIDName(item.SchoolID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by user school ids from database")
		return
	}
	if foundNewSubject != nil && foundNewSubject.SchoolID == item.SchoolID && foundNewSubject.Name == item.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Subject")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create subject from database")
		return
	}
	return
}

// Update subject
func (service *Service) Update(inputJwtToken *types.JwtToken, subjectID int64, item *model.HighschoolSubject) (result *model.HighschoolSubject, errCode int, err error) {
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

	// Check if subject exists
	foundItem, err := service.Repository.GetById(subjectID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != subjectID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Subject")
		return
	}
	// Check if the school type is highschool
	if foundItem.School.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewSubject, err := service.Repository.GetBySchoolIDName(item.SchoolID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by user school ids from database")
		return
	}
	if foundNewSubject != nil && foundNewSubject.SchoolID == item.SchoolID && foundNewSubject.Name == item.Name {
		if !(foundItem.SchoolID == foundNewSubject.SchoolID && foundItem.Name == foundNewSubject.Name) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage("Subject")
			return
		}
	}

	// Update subject
	result, err = service.Repository.Update(subjectID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update subject from database")
		return
	}
	return
}

// Delete subject with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, subjectID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(subjectID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete subject from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Subject")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple subject from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Subject selection")
		return
	}
	return
}

// Get Returns subject with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, subjectID int64) (result *model.HighschoolSubject, errCode int, err error) {
	result, err = service.Repository.GetById(subjectID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Subject")
		return
	}
	return
}

// GetAll Returns all subjects with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolSubject, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subjects from database")
	}
	return
}
