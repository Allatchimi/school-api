package parent

import (
	"net/http"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/config"
	"api/services/school/common/parent/data"
	"api/services/school/common/parent/model"
	"api/services/school/common/school"
	"api/services/user/role"
	"api/services/user/user"
	dataUser "api/services/user/user/data"
)

type Service struct {
	Repository    *Repository
	RoleService   *role.Service
	UserService   *user.Service
	SchoolService *school.Service
}

func NewService(
	repository *Repository,
	roleService *role.Service,
	userService *user.Service,
	schoolService *school.Service,
) *Service {
	return &Service{
		Repository:    repository,
		RoleService:   roleService,
		UserService:   userService,
		SchoolService: schoolService,
	}
}

const MODEL_NAME = "parent"
const DEFAULT_ERROR_MESSAGE = "interact with parent model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.ParentRequest,
) (result *model.Parent, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleParent)
	if errRole != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Get the school
	foundSchool, errSchool := service.SchoolService.Repository.GetByID(newRequest.SchoolID)
	if errSchool != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundSchool == nil || foundSchool.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("school")
		return
	}

	// Check request
	if request == nil || newRequest.Info == nil {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Generate the email
	newEmail := newRequest.Email
	if newRequest.AutoGenerateEmail {
		if foundSchool.Config == nil {
			errCode = http.StatusNotFound
			err = constants.Http404ErrorMessage("school")
			return
		}
		newEmail = helpers.GenerateEmailFromFullName(
			newRequest.Info.FirstName,
			newRequest.Info.LastName,
			foundSchool.Config.UserEmailDomainName,
		)
	}

	// Format request
	var item = &dataUser.UserRequest{
		SchoolID: newRequest.SchoolID,
		RoleID:   userRole.ID,

		Email:       newEmail,
		PhoneNumber: newRequest.PhoneNumber,
		IsActivated: true,
		Status:      newRequest.Status,
		Info: &dataUser.UserInfoRequest{
			Gender:        newRequest.Info.Gender,
			Username:      newRequest.Info.Username,
			FirstName:     newRequest.Info.FirstName,
			LastName:      newRequest.Info.LastName,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}

	// Generate password
	password := helpers.GeneratePasswordFromUser(
		item.Info.FirstName,
		item.Info.LastName,
		item.Info.Birthday,
	)

	// Create user
	createdUser, errCodeCreate, errCreate := service.UserService.Create(ctxData, item, &password)
	if errCreate != nil {
		errCode = errCodeCreate
		err = errCreate
		return
	}
	if createdUser == nil || createdUser.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create
	result, err = service.Repository.Create(&model.Parent{
		SchoolID: newRequest.SchoolID,
		UserID:   createdUser.ID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateParentStudent(
	ctxData *types.ContextData,
	request *data.ParentStudentRequest,
) (result *model.ParentStudent, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.ParentStudent{
		SchoolID:  newRequest.SchoolID,
		ParentID:  newRequest.ParentID,
		StudentID: newRequest.StudentID,
	}

	// Check unique
	foundUnique, err := service.Repository.GetParentStudentUniqueObjectByUserID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreParentStudentSameUniqueObjectsByUserID(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create
	result, err = service.Repository.CreateParentStudent(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateParentAssign(
	ctxData *types.ContextData,
	request *data.ParentAssignRequest,
) (result *model.ParentAssign, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.ParentAssign{
		SchoolID: newRequest.SchoolID,
		UserID:   ctxData.Jwt.UserID,

		StudentListID:  newRequest.StudentListID,
		Status:         newRequest.StudentListID,
		StatusFeedback: newRequest.StudentListID,
		Message:        newRequest.Message,
		Gender:         newRequest.Gender,
		FirstName:      newRequest.FirstName,
		LastName:       newRequest.LastName,
		Birthday:       newRequest.Birthday,
		BirthLocation:  newRequest.BirthLocation,
		Document1:      newRequest.Document1,
		Document2:      newRequest.Document2,
		Document3:      newRequest.Document3,
		Document4:      newRequest.Document4,
		Document5:      newRequest.Document5,
	}

	// Check unique
	foundUnique, err := service.Repository.GetParentAssignUniqueObjectByUserID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreParentAssignSameUniqueObjectsByUserID(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create
	result, err = service.Repository.CreateParentAssign(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.ParentRequest,
) (result *model.Parent, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.Parent
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Get the role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleParent)
	if errRole != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Get the school
	foundSchool, errSchool := service.SchoolService.Repository.GetByID(newRequest.SchoolID)
	if errSchool != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundSchool == nil || foundSchool.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("school")
		return
	}

	// Update user
	userRequest := dataUser.UserRequest{
		SchoolID: newRequest.SchoolID,
		RoleID:   userRole.ID,

		Email:       newRequest.Email,
		PhoneNumber: newRequest.PhoneNumber,
		IsActivated: true,
		Status:      newRequest.Status,
		Info: &dataUser.UserInfoRequest{
			Gender:        newRequest.Info.Gender,
			Username:      newRequest.Info.Username,
			FirstName:     newRequest.Info.FirstName,
			LastName:      newRequest.Info.LastName,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}
	_, errCodeUser, errUser := service.UserService.Update(ctxData, foundItem.UserID, &userRequest)
	if errUser != nil {
		errCode = errCodeUser
		err = errUser
		return
	}

	// Update parent
	result, err = service.Repository.UpdateByID(id, &model.Parent{
		SchoolID: foundItem.SchoolID,
		UserID:   foundItem.UserID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateParentStudent(
	ctxData *types.ContextData,
	id int64,
	request *data.ParentStudentRequest,
) (result *model.ParentStudent, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ParentStudent
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetParentStudentByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetParentStudentByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format request
	item := &model.ParentStudent{
		ParentID:  newRequest.ParentID,
		StudentID: newRequest.StudentID,
	}

	// Check unique
	foundUnique, err := service.Repository.GetParentStudentUniqueObjectByUserID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreParentStudentSameUniqueObjectsByUserID(foundUnique, item) &&
		!service.Repository.AreParentStudentSameUniqueObjectsByUserID(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update
	result, err = service.Repository.UpdateParentStudentByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateParentAssign(
	ctxData *types.ContextData,
	id int64,
	request *data.ParentAssignRequest,
) (result *model.ParentAssign, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ParentAssign
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetParentAssignByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetParentAssignByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format request
	item := &model.ParentAssign{
		SchoolID: newRequest.SchoolID,
		UserID:   foundItem.UserID,

		StudentListID:  newRequest.StudentListID,
		Status:         newRequest.StudentListID,
		StatusFeedback: newRequest.StudentListID,
		Message:        newRequest.Message,
		Gender:         newRequest.Gender,
		FirstName:      newRequest.FirstName,
		LastName:       newRequest.LastName,
		Birthday:       newRequest.Birthday,
		BirthLocation:  newRequest.BirthLocation,
		Document1:      newRequest.Document1,
		Document2:      newRequest.Document2,
		Document3:      newRequest.Document3,
		Document4:      newRequest.Document4,
		Document5:      newRequest.Document5,
	}

	// Check unique
	foundUnique, err := service.Repository.GetParentAssignUniqueObjectByUserID(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreParentAssignSameUniqueObjectsByUserID(foundUnique, item) &&
		!service.Repository.AreParentAssignSameUniqueObjectsByUserID(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update
	result, err = service.Repository.UpdateParentAssignByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}
func (service *Service) UpdateParentAssignStatus(
	ctxData *types.ContextData,
	id int64,
	request *data.ParentAssignStatusRequest,
) (result *model.ParentAssign, errCode int, err error) {
	// Check school
	newRequest := *request
	var newSchoolID int64
	if ctxData.Jwt.SchoolID > 0 {
		newSchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ParentAssign
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetParentAssignByIDSchoolID(id, newSchoolID)
	} else {
		foundItem, err = service.Repository.GetParentAssignByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format request
	item := &model.ParentAssign{
		Status:         newRequest.Status,
		StatusFeedback: newRequest.StatusFeedback,
	}

	// Update
	result, err = service.Repository.UpdateParentAssignStatusByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.Parent
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
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

func (service *Service) DeleteParentStudent(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ParentStudent
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetParentStudentByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetParentStudentByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteParentStudentByID(id)
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

func (service *Service) DeleteParentAssign(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ParentAssign
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetParentAssignByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetParentAssignByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteParentAssignByID(id)
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

func (service *Service) DeleteMultiple(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleByID(list, ctxData.Jwt.SchoolID)
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

func (service *Service) DeleteMultipleParentStudent(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleParentStudentByID(list, ctxData.Jwt.SchoolID)
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

func (service *Service) DeleteMultipleParentAssign(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleParentAssignByID(list, ctxData.Jwt.SchoolID)
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

func (service *Service) Get(
	ctxData *types.ContextData,
	id int64,
) (result *model.Parent, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetByID(id)
	}
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

func (service *Service) GetParentStudent(
	ctxData *types.ContextData,
	id int64,
) (result *model.ParentStudent, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetParentStudentByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetParentStudentByID(id)
	}
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

func (service *Service) GetParentAssign(
	ctxData *types.ContextData,
	id int64,
) (result *model.ParentAssign, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetParentAssignByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetParentAssignByID(id)
	}
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

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Parent, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllParentStudent(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllParentStudentRequest,
) (result []model.ParentStudent, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllParentStudent(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllParentAssign(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllParentAssignRequest,
) (result []model.ParentAssign, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllParentAssign(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
