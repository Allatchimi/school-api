package year

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/year/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new year
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Year) (result *model.Year, errCode int, err error) {
	yearName := fmt.Sprintf("%d-%d", item.StartDate.Year(), item.EndDate.Year())
	item.Name = yearName

	// Check if the year exists
	foundItem, err := service.Repository.GetByNameSchoolID(yearName, item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get year by name from database")
		return
	}
	if foundItem != nil && foundItem.Name == yearName && foundItem.SchoolID == item.SchoolID {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("year")
		return
	}

	// Create
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create year from database")
		return
	}
	return
}

// Update year
func (service *Service) Update(inputJwtToken *types.JwtToken, yearID int64, item *model.Year) (result *model.Year, errCode int, err error) {
	// Check if year exists
	foundItem, err := service.Repository.GetById(yearID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get year by name from database")
		return
	}
	if foundItem == nil || len(foundItem.Name) < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Year")
		return
	}

	yearName := fmt.Sprintf("%d-%d", item.StartDate.Year(), item.EndDate.Year())
	item.Name = yearName
	// Check if the year exists
	foundNewYear, err := service.Repository.GetByNameSchoolID(yearName, item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get year by name from database")
		return
	}
	if foundNewYear != nil && foundNewYear.SchoolID == item.SchoolID && foundNewYear.Name == yearName {
		if !(foundItem.SchoolID == foundNewYear.SchoolID && foundItem.Name == foundNewYear.Name) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage("year")
			return
		}
	}

	// Update year
	result, err = service.Repository.Update(yearID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update year from database")
		return
	}
	return
}

// Delete year with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, yearID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(yearID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete year from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Year")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple year from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Year selection")
		return
	}
	return
}

// Get Returns year with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, yearID int64) (result *model.Year, errCode int, err error) {
	result, err = service.Repository.GetById(yearID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get year by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Year")
		return
	}
	return
}

// GetAll Returns all years with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Year, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get years from database")
	}
	return
}
