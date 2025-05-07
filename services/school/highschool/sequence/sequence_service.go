package sequence

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/highschool/sequence/model"
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

// Create new sequence
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.HighschoolSequence) (result *model.HighschoolSequence, errCode int, err error) {
	// Check if the school type is highschool
	foundSchool, err := service.SchoolRepository.GetById(item.SchoolID)
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

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get sequence by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Sequence")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create sequence from database")
		return
	}
	return
}

// Update sequence
func (service *Service) Update(inputJwtToken *types.JwtToken, sequenceID int64, item *model.HighschoolSequence) (result *model.HighschoolSequence, errCode int, err error) {
	// Check if the school type is highschool
	foundSchool, err := service.SchoolRepository.GetById(item.SchoolID)
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

	// Check if sequence exists
	foundItem, err := service.Repository.GetById(sequenceID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get sequence by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != sequenceID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Sequence")
		return
	}
	// Check if the school type is highschool
	if foundItem.School.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get sequence by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Sequence")
		return
	}

	// Update sequence
	result, err = service.Repository.Update(sequenceID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update sequence from database")
		return
	}
	return
}

// Delete sequence with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, sequenceID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(sequenceID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete sequence from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Sequence")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple sequence from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Sequence selection")
		return
	}
	return
}

// Get Returns sequence with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, sequenceID int64) (result *model.HighschoolSequence, errCode int, err error) {
	result, err = service.Repository.GetById(sequenceID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get sequence by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Sequence")
		return
	}
	return
}

// GetAll Returns all sequences with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolSequence, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get sequences from database")
	}
	return
}
