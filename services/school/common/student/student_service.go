package student

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/student/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new student
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Student) (result *model.Student, errCode int, err error) {
	// Check if student already exists
	foundItem, err := service.Repository.GetByObject(&model.Student{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
		UID:      item.UID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("student")
		return
	}

	// Insert student
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create student from database")
		return
	}
	return
}

// Create new student
func (service *Service) CreateLevelClass(inputJwtToken *types.JwtToken, item *model.StudentLevelClass) (result *model.StudentLevelClass, errCode int, err error) {
	// Check if student level/class already exists
	foundItem, err := service.Repository.GetLevelClassByObject(&model.StudentLevelClass{
		StudentID: item.StudentID,
		YearID:    item.YearID,

		DomainID: item.DomainID,
		LevelID:  item.LevelID,

		ClassID: item.ClassID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student level/class by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("student level/class")
		return
	}

	// Insert student level/class
	result, err = service.Repository.CreateLevelClass(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create student level/class from database")
		return
	}
	return
}

// Update student
func (service *Service) Update(inputJwtToken *types.JwtToken, studentID int64, item *model.Student) (result *model.Student, errCode int, err error) {
	// Check if student already exists
	foundStudentByID, err := service.Repository.GetById(studentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by name from database")
		return
	}
	if foundStudentByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student")
		return
	}
	foundItem, err := service.Repository.GetByObject(&model.Student{
		SchoolID: foundStudentByID.SchoolID,
		UserID:   foundStudentByID.UserID,
		UID:      item.UID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("student")
		return
	}

	// Update student
	result, err = service.Repository.Update(studentID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update student from database")
		return
	}
	return
}

// Update student level/class
func (service *Service) UpdateLevelClass(inputJwtToken *types.JwtToken, studentLevelClassID int64, item *model.StudentLevelClass) (result *model.StudentLevelClass, errCode int, err error) {
	// Check if student already exists
	foundStudentByID, err := service.Repository.GetLevelClassById(studentLevelClassID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student level/class by name from database")
		return
	}
	if foundStudentByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student level/class")
		return
	}
	foundItem, err := service.Repository.GetLevelClassByObject(&model.StudentLevelClass{
		StudentID: item.StudentID,
		YearID:    item.YearID,

		DomainID: item.DomainID,
		LevelID:  item.LevelID,

		ClassID: item.ClassID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student level/class by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("student level/class")
		return
	}

	// Update student
	result, err = service.Repository.UpdateLevelClass(studentLevelClassID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update student level/class from database")
		return
	}
	return
}

// Delete student with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, studentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(studentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete student from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student")
		return
	}
	return
}

// Delete student level/class with matching id and return affected rows
func (service *Service) DeleteLevelClass(inputJwtToken *types.JwtToken, studentLevelClassID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteLevelClass(studentLevelClassID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete student level/class from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student level/class")
		return
	}
	return
}

// Get Returns student with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, studentID int64) (result *model.Student, errCode int, err error) {
	result, err = service.Repository.GetById(studentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student")
		return
	}
	return
}

// Get Returns student level/class with matching id
func (service *Service) GetLevelClass(inputJwtToken *types.JwtToken, studentLevelClassID int64) (result *model.StudentLevelClass, errCode int, err error) {
	result, err = service.Repository.GetLevelClassById(studentLevelClassID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student level/class by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student level/class")
		return
	}
	return
}

// GetAll Returns all students with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Student, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get students from database")
	}
	return
}

// GetAll Returns all students level/class with support for search, filter and pagination
func (service *Service) GetAllLevelClass(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, studentID int64) (result []model.StudentLevelClass, errCode int, err error) {
	result, err = service.Repository.GetAllLevelClass(filter, pagination, studentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get students level/class from database")
	}
	return
}
