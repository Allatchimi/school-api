package quarter

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/highschool/quarter/model"
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

// Create new quarter
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.HighschoolQuarter) (result *model.HighschoolQuarter, errCode int, err error) {
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
		err = constants.Http500ErrorMessage("get quarter by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Quarter")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create quarter from database")
		return
	}
	return
}

// Update quarter
func (service *Service) Update(inputJwtToken *types.JwtToken, quarterID int64, item *model.HighschoolQuarter) (result *model.HighschoolQuarter, errCode int, err error) {
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

	// Check if quarter exists
	foundItem, err := service.Repository.GetById(quarterID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get quarter by name from database")
		return
	}
	if foundItem == nil || foundItem.ID != quarterID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Quarter")
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
		err = constants.Http500ErrorMessage("get quarter by user school ids from database")
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Quarter")
		return
	}

	// Update quarter
	result, err = service.Repository.Update(quarterID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update quarter from database")
		return
	}
	return
}

// Delete quarter with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, quarterID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(quarterID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete quarter from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Quarter")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple quarter from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Quarter selection")
		return
	}
	return
}

// Get Returns quarter with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, quarterID int64) (result *model.HighschoolQuarter, errCode int, err error) {
	result, err = service.Repository.GetById(quarterID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get quarter by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Quarter")
		return
	}
	return
}

// GetAll Returns all quarters with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolQuarter, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get quarters from database")
	}
	return
}
