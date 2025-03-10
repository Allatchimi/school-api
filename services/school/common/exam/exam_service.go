package exam

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/exam/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new exam type
func (service *Service) CreateType(inputJwtToken *types.JwtToken, item *model.ExamType) (result *model.ExamType, errCode int, err error) {
	// Check if exam type already exists
	foundItem, err := service.Repository.GetTypeByObject(&model.ExamType{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("exam")
		return
	}

	// Insert exam
	result, err = service.Repository.CreateType(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create exam from database")
		return
	}
	return
}

// Create new exam
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Exam) (result *model.Exam, errCode int, err error) {
	// Check if exam already exists
	foundItem, err := service.Repository.GetByObject(&model.Exam{
		SchoolID:       item.SchoolID,
		Type:           item.Type,
		TeachingUnitID: item.TeachingUnitID,
		SubjectID:      item.SubjectID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("exam")
		return
	}

	// Insert exam
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create exam from database")
		return
	}
	return
}

// Update exam type
func (service *Service) UpdateType(inputJwtToken *types.JwtToken, id int64, item *model.ExamType) (result *model.ExamType, errCode int, err error) {
	// Check if exam already exists
	foundExamByID, err := service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundExamByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	foundItem, err := service.Repository.GetTypeByObject(&model.ExamType{
		// SubjectID:   item.SubjectID,
		SchoolID:    item.SchoolID,
		Name:        item.Name,
		Description: item.Description,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("exam")
		return
	}

	// Update exam
	result, err = service.Repository.UpdateType(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update exam from database")
		return
	}
	return
}

// Update exam
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.Exam) (result *model.Exam, errCode int, err error) {
	// Check if exam already exists
	foundExamByID, err := service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundExamByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	foundItem, err := service.Repository.GetByObject(&model.Exam{
		SchoolID:       item.SchoolID,
		Type:           item.Type,
		TeachingUnitID: item.TeachingUnitID,
		SubjectID:      item.SubjectID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("exam")
		return
	}

	// Update exam
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update exam from database")
		return
	}
	return
}

// Delete exam type with matching id and return affected rows
func (service *Service) DeleteType(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteType(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete exam from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	return
}

// Delete exam with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete exam from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	return
}

// GetType Returns exam type with matching id
func (service *Service) GetType(inputJwtToken *types.JwtToken, examID int64) (result *model.ExamType, errCode int, err error) {
	result, err = service.Repository.GetTypeById(examID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	return
}

// Get Returns exam with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, examID int64) (result *model.Exam, errCode int, err error) {
	result, err = service.Repository.GetById(examID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	return
}

// GetAll Returns all exams with support for search, filter and pagination
func (service *Service) GetAllType(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.ExamType, errCode int, err error) {
	// result, err = service.Repository.GetAllExamType(filter, pagination, inputJwtToken.UserID)
	// if err != nil {
	// 	errCode = http.StatusInternalServerError
	// 	err = constants.Http500ErrorMessage("get exams from database")
	// }
	return
}

// GetAll Returns all exams with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Exam, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exams from database")
	}
	return
}
