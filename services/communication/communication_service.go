package communication

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/communication/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create Creates a new communication and returns the newly created
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Communication) (result *model.Communication, errCode int, err error) {
	// Insert communication
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create communication from database")
		return
	}
	return
}

// Get Returns communication with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Communication, errCode int, err error) {
	result, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get communication from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("communication")
		return
	}
	return
}

// GetAll Returns communication list with matching for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Communication, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get communication list from database")
	}
	return
}
