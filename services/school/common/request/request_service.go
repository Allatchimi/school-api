package request

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	serviceHelperFeature "api/services/helper/feature"
	serviceHelperMessage "api/services/helper/message"
	serviceHelperUser "api/services/helper/user"
	"api/services/school/common/request/data"
	"api/services/school/common/request/model"
	"api/services/school/common/student"
	dataStudent "api/services/school/common/student/data"

	"go.uber.org/zap"
)

type Service struct {
	Repository     *Repository
	StudentService *student.Service
}

func NewService(
	repository *Repository,
	studentService *student.Service,
) *Service {
	return &Service{
		Repository:     repository,
		StudentService: studentService,
	}
}

const MODEL_NAME = "request"
const DEFAULT_ERROR_MESSAGE = "interact with request model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.RequestRequest,
) (result *model.Request, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the user is student
	if ctxData.User.Feature != constants.FeatureStudent {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Format request
	item := &model.Request{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		SequenceID:     newRequest.SequenceID,
		UnitID:         newRequest.UnitID,

		Audience: newRequest.Audience,
		Title:    newRequest.Title,
		Message:  newRequest.Message,

		Document1: newRequest.Document1,
		Document2: newRequest.Document2,
		Document3: newRequest.Document3,
		Document4: newRequest.Document4,
		Document5: newRequest.Document5,
	}
	if newRequest.ClassSubjectID < 1 && newRequest.UnitID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("class subject id or unit id(you should provide one of these fields)")
		return
	}
	if newRequest.ClassSubjectID > 0 && newRequest.SequenceID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("sequence id")
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Get the student and update the item
	foundStudent, err := service.StudentService.Repository.GetByUserIDSchoolID(ctxData.Jwt.UserID, newRequest.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundStudent == nil || foundStudent.ID < 1 {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}
	item.StudentID = foundStudent.ID

	// Create
	result, err = service.Repository.Create(item)
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
	request *data.RequestRequest,
) (result *model.Request, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the user is student
	if ctxData.User.Feature != constants.FeatureStudent {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Check if the item exists
	var foundItem *model.Request
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
	if foundItem.Student.UserID != ctxData.Jwt.UserID {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Format request
	item := &model.Request{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		SequenceID:     newRequest.SequenceID,
		UnitID:         newRequest.UnitID,
		StudentID:      foundItem.StudentID,

		Audience: newRequest.Audience,
		Title:    newRequest.Title,
		Message:  newRequest.Message,

		Document1: newRequest.Document1,
		Document2: newRequest.Document2,
		Document3: newRequest.Document3,
		Document4: newRequest.Document4,
		Document5: newRequest.Document5,
	}
	if newRequest.ClassSubjectID < 1 && newRequest.UnitID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("class subject id or unit id(you should provide one of these fields)")
		return
	}
	if newRequest.ClassSubjectID > 0 && newRequest.SequenceID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("sequence id")
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateStatus(
	ctxData *types.ContextData,
	id int64,
	request *data.RequestStatusRequest,
) (result *model.Request, errCode int, err error) {
	// Check school
	newRequest := *request

	// Check if the item exists
	var foundItem *model.Request
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

	// Format request
	item := &model.Request{
		Status:         newRequest.Status,
		StatusFeedback: newRequest.StatusFeedback,
	}

	// Update request
	result, err = service.Repository.UpdateStatusByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Send message to student
	go func() {
		if result == nil || result.Student == nil || result.StudentID < 1 || result.Status == foundItem.Status {
			return
		}
		if !(result.Status == constants.REQUEST_STATUS_COMPLETED || result.Status == constants.REQUEST_STATUS_REJECTED) {
			return
		}
		// Students
		studentEnrollReq := &dataStudent.GetAllStudentEnrollRequest{}
		studentEnrollReq.SchoolID = result.SchoolID
		studentEnrollReq.YearID = result.YearID
		studentEnrollReq.ClassSubjectID = result.ClassSubjectID
		studentEnrollReq.UnitID = result.UnitID
		userStudents, errUsers := serviceHelperUser.GetAllUserForStudentEnroll(studentEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for student enroll", zap.Error(errUsers))
			return
		}
		var title, message string
		switch result.Status {
		case constants.REQUEST_STATUS_COMPLETED:
			title = "Request Completed"
		case constants.REQUEST_STATUS_REJECTED:
			title = "Request Rejected"
		}
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.ClassSubject != nil && result.ClassSubject.Subject != nil && result.ClassSubject.Class != nil && result.Sequence != nil && result.Year != nil {
				message = fmt.Sprintf("The request for subject %s %s: %s %s has been %s", result.ClassSubject.Subject.Name, result.ClassSubject.Class.Name, result.Sequence.Name, result.Year.Name, result.Status)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.Unit != nil && result.Unit.Semester != nil && result.Year != nil {
				message = fmt.Sprintf("The request for unit %s: %s %s has been %s", result.Unit.Name, result.Unit.Semester.Name, result.Year.Name, result.Year.Name)
			}
		}
		serviceHelperMessage.SendMessage(
			&serviceHelperMessage.MessageRequest{
				PusNotification: true,
				Telegram:        true,
				Whatsapp:        true,
				Mail:            true,
			},
			title,
			message,
			result.School,
			"",
			userStudents,
		)
	}()
	return
}

func (service *Service) Delete(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
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

func (service *Service) Get(
	ctxData *types.ContextData,
	id int64,
) (result *model.Request, errCode int, err error) {
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

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Request, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check feature
	if ctxData.User.Feature != constants.FeatureAdmin {
		var okCheck bool
		var errCheck error
		newRequest.TeacherID,
			newRequest.StudentID,
			newRequest.ParentID,
			okCheck,
			errCheck = serviceHelperFeature.GetUserDataByFeatureName(ctxData)
		if errCheck != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		if !okCheck {
			return
		}
		if ctxData.User.Feature == constants.FeatureTeacher {
			newRequest.Audience = constants.REQUEST_AUDIENCE_TEACHER
		}
	}

	// Get
	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
