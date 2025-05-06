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

// Create new teacher
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Teacher) (result *model.Teacher, errCode int, err error) {
	// Check unique
	foundItem, err := service.Repository.GetByObject(&model.Teacher{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
		UID:      item.UID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher by name from database")
		return
	}
	if foundItem != nil && foundItem.UserID == item.UserID {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teacher")
		return
	}

	// Insert teacher
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create teacher from database")
		return
	}
	return
}

// Create new teacher
func (service *Service) CreateTeacherTUSubject(inputJwtToken *types.JwtToken, item *model.TeacherTeachingUnitSubject) (result *model.TeacherTeachingUnitSubject, errCode int, err error) {
	// Check if teacher teaching unit/subject already exists
	foundItem, err := service.Repository.GetTeacherTUSubjectByObject(&model.TeacherTeachingUnitSubject{
		TeacherID: item.TeacherID,
		YearID:    item.YearID,

		TeachingUnitID: item.TeachingUnitID,
		SubjectID:      item.SubjectID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher teaching unit/subject by name from database")
		return
	}
	if foundItem != nil && foundItem.TeacherID == item.TeacherID && foundItem.YearID == item.YearID {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teacher teaching unit/subject")
		return
	}

	// Insert teacher teaching unit/subject
	result, err = service.Repository.CreateTeacherTUSubject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create teacher teaching unit/subject from database")
		return
	}
	return
}

// Update teacher
func (service *Service) Update(inputJwtToken *types.JwtToken, teacherID int64, item *model.Teacher) (result *model.Teacher, errCode int, err error) {
	// Check unique
	foundTeacherByID, err := service.Repository.GetById(teacherID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher by name from database")
		return
	}
	if foundTeacherByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher")
		return
	}
	foundItem, err := service.Repository.GetByObject(&model.Teacher{
		SchoolID: foundTeacherByID.SchoolID,
		UserID:   foundTeacherByID.UserID,
		UID:      item.UID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teacher")
		return
	}

	// Update teacher
	result, err = service.Repository.Update(teacherID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update teacher from database")
		return
	}
	return
}

// Update teacher teaching unit/subject
func (service *Service) UpdateTeacherTUSubject(inputJwtToken *types.JwtToken, teacherTUSubjectID int64, item *model.TeacherTeachingUnitSubject) (result *model.TeacherTeachingUnitSubject, errCode int, err error) {
	// Check unique
	foundTeacherByID, err := service.Repository.GetTeacherTUSubjectById(teacherTUSubjectID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher teaching unit/subject by name from database")
		return
	}
	if foundTeacherByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher teaching unit/subject")
		return
	}
	foundItem, err := service.Repository.GetTeacherTUSubjectByObject(&model.TeacherTeachingUnitSubject{
		TeacherID: item.TeacherID,
		YearID:    item.YearID,

		TeachingUnitID: item.TeachingUnitID,
		SubjectID:      item.SubjectID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher teaching unit/subject by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teacher teaching unit/subject")
		return
	}

	// Update teacher
	result, err = service.Repository.UpdateTeacherTUSubject(teacherTUSubjectID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update teacher teaching unit/subject from database")
		return
	}
	return
}

// Delete teacher with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, teacherID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(teacherID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete teacher from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher")
		return
	}
	return
}

// Delete teacher teaching unit/subject with matching id and return affected rows
func (service *Service) DeleteTeacherTUSubject(inputJwtToken *types.JwtToken, teacherTUSubjectID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteTeacherTUSubject(teacherTUSubjectID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete teacher teaching unit/subject from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher teaching unit/subject")
		return
	}
	return
}

// Get Returns teacher with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, teacherID int64) (result *model.Teacher, errCode int, err error) {
	result, err = service.Repository.GetById(teacherID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher")
		return
	}
	return
}

// Get Returns teacher teaching unit/subject with matching id
func (service *Service) GetTeacherTUSubject(inputJwtToken *types.JwtToken, teacherTUSubjectID int64) (result *model.TeacherTeachingUnitSubject, errCode int, err error) {
	result, err = service.Repository.GetTeacherTUSubjectById(teacherTUSubjectID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher teaching unit/subject by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher teaching unit/subject")
		return
	}
	return
}

// GetAll Returns all teachers with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Teacher, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teachers from database")
	}
	return
}

// GetAll Returns all teachers teaching unit/subject with support for search, filter and pagination
func (service *Service) GetAllTeacherTUSubject(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, teacherID int64) (result []model.TeacherTeachingUnitSubject, errCode int, err error) {
	result, err = service.Repository.GetAllTeacherTUSubject(filter, pagination, teacherID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teachers teaching unit/subject from database")
	}
	return
}
