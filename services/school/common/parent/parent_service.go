package parent

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/parent/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new parent
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.Parent) (result *model.Parent, errCode int, err error) {
	// Check if parent already exists
	foundItem, err := service.Repository.GetByObject(&model.Parent{
		UserID: item.UserID,
		UID:    item.UID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("parent")
		return
	}

	// Insert parent
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create parent from database")
		return
	}
	return
}

// Create new parent
func (service *Service) CreateParentStudent(inputJwtToken *types.JwtToken, item *model.ParentStudent) (result *model.ParentStudent, errCode int, err error) {
	// Check if parent level/class already exists
	foundItem, err := service.Repository.GetParentStudentByObject(&model.ParentStudent{
		ParentID:  item.ParentID,
		StudentID: item.StudentID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent level/class by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("parent level/class")
		return
	}

	// Insert parent level/class
	result, err = service.Repository.CreateParentStudent(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create parent level/class from database")
		return
	}
	return
}

// Update parent
func (service *Service) Update(inputJwtToken *types.JwtToken, parentID int64, item *model.Parent) (result *model.Parent, errCode int, err error) {
	// Check if parent already exists
	foundParentByID, err := service.Repository.GetById(parentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent by name from database")
		return
	}
	if foundParentByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Parent")
		return
	}
	foundItem, err := service.Repository.GetByObject(&model.Parent{
		UserID: foundParentByID.UserID,
		UID:    item.UID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("parent")
		return
	}

	// Update parent
	result, err = service.Repository.Update(parentID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update parent from database")
		return
	}
	return
}

// Update parent level/class
func (service *Service) UpdateParentStudent(inputJwtToken *types.JwtToken, parentParentStudentID int64, item *model.ParentStudent) (result *model.ParentStudent, errCode int, err error) {
	// Check if parent already exists
	foundParentByID, err := service.Repository.GetParentStudentById(parentParentStudentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent level/class by name from database")
		return
	}
	if foundParentByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Parent level/class")
		return
	}
	foundItem, err := service.Repository.GetParentStudentByObject(&model.ParentStudent{
		ParentID:  item.ParentID,
		StudentID: item.StudentID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent level/class by name from database")
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("parent level/class")
		return
	}

	// Update parent
	result, err = service.Repository.UpdateParentStudent(parentParentStudentID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update parent level/class from database")
		return
	}
	return
}

// Delete parent with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, parentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(parentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete parent from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Parent")
		return
	}
	return
}

// Delete parent level/class with matching id and return affected rows
func (service *Service) DeleteParentStudent(inputJwtToken *types.JwtToken, parentParentStudentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteParentStudent(parentParentStudentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete parent level/class from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Parent level/class")
		return
	}
	return
}

// Get Returns parent with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, parentID int64) (result *model.Parent, errCode int, err error) {
	result, err = service.Repository.GetById(parentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Parent")
		return
	}
	return
}

// Get Returns parent level/class with matching id
func (service *Service) GetParentStudent(inputJwtToken *types.JwtToken, parentParentStudentID int64) (result *model.ParentStudent, errCode int, err error) {
	result, err = service.Repository.GetParentStudentById(parentParentStudentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent level/class by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Parent level/class")
		return
	}
	return
}

// GetAll Returns all parents with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Parent, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parents from database")
	}
	return
}

// GetAll Returns all parents level/class with support for search, filter and pagination
func (service *Service) GetAllParentStudent(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.ParentStudent, errCode int, err error) {
	result, err = service.Repository.GetAllParentStudent(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parents level/class from database")
	}
	return
}
