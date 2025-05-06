package contact

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/contact/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create Creates a new contact and returns the newly created
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Contact) (result *model.Contact, errCode int, err error) {
	// Insert contact
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create contact from database")
		return
	}
	return
}

// Get Returns contact with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Contact, errCode int, err error) {
	result, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get contact from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("contact")
		return
	}
	return
}

// GetAll Returns contact list with matching for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Contact, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get contact list from database")
	}
	return
}
