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
	// Check if teacher already exists
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
	if foundItem != nil {
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
func (service *Service) CreateLevelClass(inputJwtToken *types.JwtToken, item *model.TeacherLevelClass) (result *model.TeacherLevelClass, errCode int, err error) {
	// Check if teacher level/class already exists
	foundItem, err := service.Repository.GetLevelClassByObject(&model.TeacherLevelClass{
		TeacherID: item.TeacherID,
		YearID:    item.YearID,

		DomainID: item.DomainID,
		LevelID:  item.LevelID,

		ClassID: item.ClassID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher level/class by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teacher level/class")
		return
	}

	// Insert teacher level/class
	result, err = service.Repository.CreateLevelClass(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create teacher level/class from database")
		return
	}
	return
}

// Update teacher
func (service *Service) Update(inputJwtToken *types.JwtToken, teacherID int64, item *model.Teacher) (result *model.Teacher, errCode int, err error) {
	// Check if teacher already exists
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

// Update teacher level/class
func (service *Service) UpdateLevelClass(inputJwtToken *types.JwtToken, teacherLevelClassID int64, item *model.TeacherLevelClass) (result *model.TeacherLevelClass, errCode int, err error) {
	// Check if teacher already exists
	foundTeacherByID, err := service.Repository.GetLevelClassById(teacherLevelClassID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher level/class by name from database")
		return
	}
	if foundTeacherByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher level/class")
		return
	}
	foundItem, err := service.Repository.GetLevelClassByObject(&model.TeacherLevelClass{
		TeacherID: item.TeacherID,
		YearID:    item.YearID,

		DomainID: item.DomainID,
		LevelID:  item.LevelID,

		ClassID: item.ClassID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher level/class by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("teacher level/class")
		return
	}

	// Update teacher
	result, err = service.Repository.UpdateLevelClass(teacherLevelClassID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update teacher level/class from database")
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

// Delete teacher level/class with matching id and return affected rows
func (service *Service) DeleteLevelClass(inputJwtToken *types.JwtToken, teacherLevelClassID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteLevelClass(teacherLevelClassID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete teacher level/class from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher level/class")
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

// Get Returns teacher level/class with matching id
func (service *Service) GetLevelClass(inputJwtToken *types.JwtToken, teacherLevelClassID int64) (result *model.TeacherLevelClass, errCode int, err error) {
	result, err = service.Repository.GetLevelClassById(teacherLevelClassID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher level/class by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Teacher level/class")
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

// GetAll Returns all teachers level/class with support for search, filter and pagination
func (service *Service) GetAllLevelClass(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, teacherID int64) (result []model.TeacherLevelClass, errCode int, err error) {
	result, err = service.Repository.GetAllLevelClass(filter, pagination, teacherID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teachers level/class from database")
	}
	return
}
