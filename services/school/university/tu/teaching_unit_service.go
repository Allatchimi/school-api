package tu

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/tu/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new teaching unit
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversityTeachingUnit) (result *model.UniversityTeachingUnit, errCode int, err error) {
	// Check if teaching unit already exists
	foundItem, err := service.Repository.GetByObject(&model.UniversityTeachingUnit{
		SchoolID: item.SchoolID,
		DomainID: item.DomainID,
		Name:     item.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teaching unit by name from database")
		return
	}
	if foundItem != nil {
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
	// Check if teaching unit already exists
	foundTeachingUnitByID, err := service.Repository.GetById(teachingUnitID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teaching unit by name from database")
		return
	}
	if foundTeachingUnitByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("teaching unit")
		return
	}
	foundItem, err := service.Repository.GetByObject(&model.UniversityTeachingUnit{
		SchoolID: foundTeachingUnitByID.SchoolID,
		DomainID: foundTeachingUnitByID.DomainID,
		Name:     item.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teaching unit by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teaching unit")
		return
	}

	// Update teaching unit
	result, err = service.Repository.Update(teachingUnitID, inputJwtToken.UserID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update teaching unit from database")
		return
	}
	return
}

// Delete teaching unit with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, teachingUnitID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(teachingUnitID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete teaching unit from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("teaching unit")
		return
	}
	return
}

// Get Returns teaching unit with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, teachingUnitID int64) (result *model.UniversityTeachingUnit, errCode int, err error) {
	result, err = service.Repository.GetById(teachingUnitID, inputJwtToken.UserID)
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
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.UniversityTeachingUnit, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
