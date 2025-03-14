package subject

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/highschool/subject/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new subject
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Subject) (result *model.Subject, errCode int, err error) {
	// Check if subject already exists
	foundItem, err := service.Repository.GetByObject(&model.Subject{
		SchoolID:     item.SchoolID,
		ClassID:      item.ClassID,
		Name:         item.Name,
		Description:  item.Description,
		Coefficient:  item.Coefficient,
		Program:      item.Program,
		Requirements: item.Requirements,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("subject")
		return
	}

	// Insert subject
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create subject from database")
		return
	}
	return
}

// AddProfessor adds new professor
func (service *Service) AddProfessor(inputJwtToken *types.JwtToken, item *model.SubjectProfessor) (result *model.SubjectProfessor, errCode int, err error) {
	// Check if professor already exists
	foundItem, err := service.Repository.GetProfessorById(item.SubjectID, item.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get professor by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("professor")
		return
	}

	// Insert professor
	result, err = service.Repository.AddProfessor(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create professor from database")
		return
	}
	return
}

// Update subject
func (service *Service) Update(inputJwtToken *types.JwtToken, subjectID int64, item *model.Subject) (result *model.Subject, errCode int, err error) {
	// Check if subject already exists
	foundSubjectByID, err := service.Repository.GetById(subjectID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by name from database")
		return
	}
	if foundSubjectByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("subject")
		return
	}
	foundItem, err := service.Repository.GetByObject(&model.Subject{
		SchoolID: item.SchoolID,
		ClassID:  item.ClassID,
		Name:     item.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("subject")
		return
	}

	// Update subject
	result, err = service.Repository.Update(subjectID, inputJwtToken.UserID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update subject from database")
		return
	}
	return
}

// Delete subject with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, subjectID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(subjectID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete subject from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("subject")
		return
	}
	return
}

// Delete professor with matching id and return affected rows
func (service *Service) DeleteProfessor(inputJwtToken *types.JwtToken, subjectProfessorID int64, userID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteProfessor(subjectProfessorID, userID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete professor from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("professor")
		return
	}
	return
}

// Get Returns subject with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, subjectID int64) (result *model.Subject, errCode int, err error) {
	result, err = service.Repository.GetById(subjectID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get subject by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("subject")
		return
	}
	return
}

// GetAll Returns all subjects with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Subject, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
