package teacher

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/teacher/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "teacher"
const DEFAULT_ERROR_MESSAGE = "interact with teacher model"

func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Teacher) (result *model.Teacher, errCode int, err error) {
	// Check unique by user id
	foundUnique1, err := service.Repository.GetUniqueObjectByUserID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsByUserID(foundUnique1, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}
	// Check unique by uid
	foundUnique2, err := service.Repository.GetUniqueObjectByUID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsByUID(foundUnique2, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Insert teacher
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, item *model.TeacherClassSubjectUnit) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	// Check unique
	foundUnique, err := service.Repository.GetTeacherClassSubjectUnitUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreTeacherClassSubjectUnitSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert teacher unit/subject
	result, err = service.Repository.CreateTeacherClassSubjectUnit(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.Teacher) (result *model.Teacher, errCode int, err error) {
	// Check if teacher exists
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check unique by user id
	foundUnique1, err := service.Repository.GetUniqueObjectByUserID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsByUserID(foundUnique1, item) && !service.Repository.AreSameUniqueObjectsByUserID(foundUnique1, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}
	// Check unique by uid
	foundUnique2, err := service.Repository.GetUniqueObjectByUID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsByUID(foundUnique2, item) && !service.Repository.AreSameUniqueObjectsByUID(foundUnique2, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update teacher
	result, err = service.Repository.UpdateByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, id int64, item *model.TeacherClassSubjectUnit) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	// Check if teacher unit/subject exists
	foundItem, err := service.Repository.GetTeacherClassSubjectUnitByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetTeacherClassSubjectUnitUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreTeacherClassSubjectUnitSameUniqueObjects(foundUnique, item) && !service.Repository.AreTeacherClassSubjectUnitSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update teacher
	result, err = service.Repository.UpdateTeacherClassSubjectUnitByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteTeacherClassSubjectUnitByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Teacher, errCode int, err error) {
	result, err = service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, id int64) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	result, err = service.Repository.GetTeacherClassSubjectUnitByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Teacher, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64, teacherID int64) (result []model.TeacherClassSubjectUnit, errCode int, err error) {
	result, err = service.Repository.GetAllTeacherClassSubjectUnit(filter, pagination, schoolID, teacherID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
