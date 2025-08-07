package exam

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	serviceHelperFeature "api/services/helper/feature"
	serviceHelperMessage "api/services/helper/message"
	serviceHelperUser "api/services/helper/user"
	"api/services/school/common/exam/data"
	"api/services/school/common/exam/model"
	dataParent "api/services/school/common/parent/data"
	dataStudent "api/services/school/common/student/data"
	dataTeacher "api/services/school/common/teacher/data"

	"go.uber.org/zap"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "exam/type"
const DEFAULT_ERROR_MESSAGE = "interact with exam/type model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.ExamRequest,
) (result *model.Exam, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.Exam{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		SequenceID:     newRequest.SequenceID,
		UnitID:         newRequest.UnitID,
		TypeID:         newRequest.TypeID,

		Status:          newRequest.Status,
		Notation:        newRequest.Notation,
		Percentage:      newRequest.Percentage,
		Description:     newRequest.Description,
		LocationType:    newRequest.LocationType,
		LocationDetails: newRequest.LocationDetails,
		Requirements:    newRequest.Requirements,
		AllowedItems:    newRequest.AllowedItems,
		StartDate:       newRequest.StartDate,
		EndDate:         newRequest.EndDate,

		IsRetry: newRequest.IsRetry,
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

	// Check retry count
	var retryCount int64 = 0
	if newRequest.IsRetry {
		var notRetryCount int64 = 0
		notRetryCount, err = service.Repository.CountAllUniqueIsNotRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		if notRetryCount < 1 {
			errCode = http.StatusLocked
			err = constants.Http423LockedErrorMessage()
			return
		}
		retryCount, err = service.Repository.CountAllUniqueIsRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		retryCount++
	}
	item.RetryCount = retryCount

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

	// Insert exam
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Send message to teachers, students and parents
	go func() {
		// Teachers
		teacherEnrollReq := &dataTeacher.GetAllTeacherClassSubjectUnitRequest{}
		teacherEnrollReq.SchoolID = result.SchoolID
		teacherEnrollReq.YearID = result.YearID
		teacherEnrollReq.ClassSubjectID = result.ClassSubjectID
		teacherEnrollReq.UnitID = result.UnitID
		userTeachers, errUsers := serviceHelperUser.GetAllUserForTeacherClassSubjectUnit(teacherEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for teacher class subject unit", zap.Error(errUsers))
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
		// Parents
		parentEnrollReq := &dataParent.GetAllParentStudentRequest{}
		parentEnrollReq.SchoolID = result.SchoolID
		parentEnrollReq.YearID = result.YearID
		parentEnrollReq.ClassSubjectID = result.ClassSubjectID
		parentEnrollReq.UnitID = result.UnitID
		userParents, errUsers := serviceHelperUser.GetAllUserForParentStudent(parentEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for parent student", zap.Error(errUsers))
			return
		}
		var title, message string
		switch result.Status {
		case constants.EXAM_STATUS_DRAFT:
			title = "New added exam to draft"
		case constants.EXAM_STATUS_ONLINE:
			title = "Published exam is now online"
		case constants.EXAM_STATUS_RESULTS:
			title = "Exam results are available"
		}
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.ClassSubject != nil && result.ClassSubject.Subject != nil && result.ClassSubject.Class != nil && result.Type != nil && result.Sequence != nil && result.Year != nil {
				message = fmt.Sprintf("Exam for subject %s %s: %s %s %s", result.ClassSubject.Subject.Name, result.ClassSubject.Class.Name, result.Type.Name, result.Sequence.Name, result.Year.Name)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.Unit != nil && result.Type != nil && result.Unit.Semester != nil && result.Year != nil {
				message = fmt.Sprintf("Exam for unit %s: %s %s %s", result.Unit.Name, result.Type.Name, result.Unit.Semester.Name, result.Year.Name)
			}
		}
		if result.Status == constants.EXAM_STATUS_DRAFT {
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
				userTeachers,
			)
			return
		}
		allUsers := append(userTeachers, append(userStudents, userParents...)...)
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
			allUsers,
		)

	}()
	return
}

func (service *Service) CreateType(
	ctxData *types.ContextData,
	request *data.ExamTypeRequest,
) (result *model.ExamType, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.ExamType{
		SchoolID:    newRequest.SchoolID,
		Name:        newRequest.Name,
		Description: newRequest.Description,
	}

	// Check unique
	foundUnique, err := service.Repository.GetExamTypeUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameExamTypeUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Insert exam
	result, err = service.Repository.CreateType(item)
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
	request *data.ExamRequest,
) (result *model.Exam, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.Exam
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

	// Format request
	item := &model.Exam{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		SequenceID:     newRequest.SequenceID,
		UnitID:         newRequest.UnitID,
		TypeID:         newRequest.TypeID,

		Status:          newRequest.Status,
		Notation:        newRequest.Notation,
		Percentage:      newRequest.Percentage,
		Description:     newRequest.Description,
		LocationType:    newRequest.LocationType,
		LocationDetails: newRequest.LocationDetails,
		Requirements:    newRequest.Requirements,
		AllowedItems:    newRequest.AllowedItems,
		StartDate:       newRequest.StartDate,
		EndDate:         newRequest.EndDate,

		IsRetry: newRequest.IsRetry,
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

	// Check retry count
	var retryCount int64 = 0
	if newRequest.IsRetry {
		var notRetryCount int64 = 0
		notRetryCount, err = service.Repository.CountAllUniqueIsNotRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		if notRetryCount < 1 {
			errCode = http.StatusLocked
			err = constants.Http423LockedErrorMessage()
			return
		}
		retryCount, err = service.Repository.CountAllUniqueIsRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		retryCount++
	}
	item.RetryCount = retryCount

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) &&
		!service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update exam
	result, err = service.Repository.UpdateByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Send message to teachers, students and parents
	go func() {
		if foundItem.Status == result.Status {
			return
		}
		// Teachers
		teacherEnrollReq := &dataTeacher.GetAllTeacherClassSubjectUnitRequest{}
		teacherEnrollReq.SchoolID = result.SchoolID
		teacherEnrollReq.YearID = result.YearID
		teacherEnrollReq.ClassSubjectID = result.ClassSubjectID
		teacherEnrollReq.UnitID = result.UnitID
		userTeachers, errUsers := serviceHelperUser.GetAllUserForTeacherClassSubjectUnit(teacherEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for teacher class subject unit", zap.Error(errUsers))
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
		// Parents
		parentEnrollReq := &dataParent.GetAllParentStudentRequest{}
		parentEnrollReq.SchoolID = result.SchoolID
		parentEnrollReq.YearID = result.YearID
		parentEnrollReq.ClassSubjectID = result.ClassSubjectID
		parentEnrollReq.UnitID = result.UnitID
		userParents, errUsers := serviceHelperUser.GetAllUserForParentStudent(parentEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for parent student", zap.Error(errUsers))
			return
		}
		var title, message string
		switch result.Status {
		case constants.EXAM_STATUS_DRAFT:
			title = "New added exam to draft"
		case constants.EXAM_STATUS_ONLINE:
			title = "Published exam is now online"
		case constants.EXAM_STATUS_RESULTS:
			title = "Exam results are available"
		}
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.ClassSubject != nil && result.ClassSubject.Subject != nil && result.ClassSubject.Class != nil && result.Type != nil && result.Sequence != nil && result.Year != nil {
				message = fmt.Sprintf("Exam for subject %s %s: %s %s %s", result.ClassSubject.Subject.Name, result.ClassSubject.Class.Name, result.Type.Name, result.Sequence.Name, result.Year.Name)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.Unit != nil && result.Type != nil && result.Unit.Semester != nil && result.Year != nil {
				message = fmt.Sprintf("Exam for unit %s: %s %s %s", result.Unit.Name, result.Type.Name, result.Unit.Semester.Name, result.Year.Name)
			}
		}
		if result.Status == constants.EXAM_STATUS_DRAFT {
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
				userTeachers,
			)
			return
		}
		allUsers := append(userTeachers, append(userStudents, userParents...)...)
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
			allUsers,
		)

	}()
	return
}

func (service *Service) UpdateType(
	ctxData *types.ContextData,
	id int64,
	request *data.ExamTypeRequest,
) (result *model.ExamType, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ExamType
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetExamTypeByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetExamTypeByID(id)
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
	item := &model.ExamType{
		SchoolID:    newRequest.SchoolID,
		Name:        newRequest.Name,
		Description: newRequest.Description,
	}

	// Check unique
	foundUnique, err := service.Repository.GetExamTypeUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameExamTypeUniqueObjects(foundUnique, item) &&
		!service.Repository.AreSameExamTypeUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update exam
	result, err = service.Repository.UpdateExamTypeByID(id, item)
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
	var foundItem *model.Exam
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

func (service *Service) DeleteType(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ExamType
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetExamTypeByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetExamTypeByID(id)
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
	affectedRows, err = service.Repository.DeleteExamTypeByID(id)
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

func (service *Service) DeleteMultipleType(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleExamTypeByID(list, ctxData.Jwt.SchoolID)
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
) (result *model.Exam, errCode int, err error) {
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

func (service *Service) GetType(
	ctxData *types.ContextData,
	id int64,
) (result *model.ExamType, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetExamTypeByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetExamTypeByID(id)
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
) (result []model.Exam, errCode int, err error) {
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
		if ctxData.User.Feature == constants.FeatureParent && newRequest.StudentID < 1 {
			return
		}
		if ctxData.User.Feature != constants.FeatureAdmin && ctxData.User.Feature != constants.FeatureDirector && ctxData.User.Feature != constants.FeatureTeacher {
			newRequest.StatusList = []string{constants.EXAM_STATUS_ONLINE, constants.EXAM_STATUS_RESULTS}
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

func (service *Service) GetAllExamType(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllExamTypeRequest,
) (result []model.ExamType, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllExamType(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
