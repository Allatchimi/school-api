package document

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/document/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new document
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Document) (result *model.Document, errCode int, err error) {
	// Check if document already exists
	foundItem, err := service.Repository.GetByObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get document by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("document")
		return
	}

	// Insert document
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create document from database")
		return
	}
	return
}

// Delete document with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, documentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(documentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete document from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Document")
		return
	}
	return
}

// Get Returns document with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, documentID int64) (result *model.Document, errCode int, err error) {
	result, err = service.Repository.GetById(documentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get document by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Document")
		return
	}
	return
}

// GetAll Returns all documents with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Document, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get documents from database")
	}
	return
}
