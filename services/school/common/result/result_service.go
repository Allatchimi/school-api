package result

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	serviceHelperFeature "api/services/helper/feature"
	serviceHelperMessage "api/services/helper/message"
	serviceHelperUser "api/services/helper/user"
	"api/services/school/common/exam"
	dataParent "api/services/school/common/parent/data"
	"api/services/school/common/result/data"
	"api/services/school/common/result/model"
	"api/services/school/common/student"
	dataStudent "api/services/school/common/student/data"

	"go.uber.org/zap"
)

type Service struct {
	Repository     *Repository
	ExamService    *exam.Service
	StudentService *student.Service
}

func NewService(
	repository *Repository,
	examService *exam.Service,
	studentService *student.Service,
) *Service {
	return &Service{
		Repository:     repository,
		ExamService:    examService,
		StudentService: studentService,
	}
}

const MODEL_NAME = "result"
const DEFAULT_ERROR_MESSAGE = "interact with result model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.ResultRequest,
) (result *model.Result, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.Result{
		SchoolID:  newRequest.SchoolID,
		StudentID: newRequest.StudentID,
		ExamID:    newRequest.ExamID,

		Score: newRequest.Score,
	}

	// Check exam
	foundExam, err := service.ExamService.Repository.GetByID(item.ExamID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundExam == nil || foundExam.ID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("exam (not found)")
		return

	}

	// Check if the student is enrolled
	var tempClassID, tempLevelDomainID int64
	if foundExam.ClassSubject != nil {
		tempClassID = foundExam.ClassSubject.ClassID
	}
	if foundExam.Unit != nil {
		tempLevelDomainID = foundExam.Unit.LevelDomainID
	}
	foundStudentEnroll, err := service.StudentService.
		Repository.
		GetStudentEnrollBySchoolIDYearIDClassIDLevelDomainIDStudentID(
			item.SchoolID,
			foundExam.YearID,
			tempClassID,
			tempLevelDomainID,
			item.StudentID,
		)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundStudentEnroll == nil || foundStudentEnroll.ID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("student (not enrolled)")
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

	// Create
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateTable(
	ctxData *types.ContextData,
	request *data.ResultTableRequest,
) (result *model.ResultTable, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.ResultTable{
		SchoolID: newRequest.SchoolID,
		ExamID:   newRequest.ExamID,

		Status: newRequest.Status,
	}

	// Check unique
	foundUnique, err := service.Repository.GetResultTableUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameResultTableUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Create
	result, err = service.Repository.CreateResultTable(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Send message to students and parents
	go func() {
		if result.Exam == nil || result.Exam.ID < 1 || result.Status != constants.RESULT_STATUS_PUBLISHED {
			return
		}
		// Students
		studentEnrollReq := &dataStudent.GetAllStudentEnrollRequest{}
		studentEnrollReq.SchoolID = result.SchoolID
		studentEnrollReq.YearID = result.Exam.YearID
		studentEnrollReq.ClassSubjectID = result.Exam.ClassSubjectID
		studentEnrollReq.UnitID = result.Exam.UnitID
		userStudents, errUsers := serviceHelperUser.GetAllUserForStudentEnroll(studentEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for student enroll", zap.Error(errUsers))
			return
		}
		// Parents
		parentEnrollReq := &dataParent.GetAllParentStudentRequest{}
		parentEnrollReq.SchoolID = result.SchoolID
		parentEnrollReq.YearID = result.Exam.YearID
		parentEnrollReq.ClassSubjectID = result.Exam.ClassSubjectID
		parentEnrollReq.UnitID = result.Exam.UnitID
		userParents, errUsers := serviceHelperUser.GetAllUserForParentStudent(parentEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for parent student", zap.Error(errUsers))
			return
		}
		var title, message string
		switch result.Status {
		case constants.RESULT_STATUS_PUBLISHED:
			title = "Available results for exam"
		}
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.Exam.ClassSubject != nil && result.Exam.ClassSubject.Subject != nil && result.Exam.ClassSubject.Class != nil && result.Exam.Type != nil && result.Exam.Sequence != nil && result.Exam.Year != nil {
				message = fmt.Sprintf("Results exam for subject %s %s: %s %s %s", result.Exam.ClassSubject.Subject.Name, result.Exam.ClassSubject.Class.Name, result.Exam.Type.Name, result.Exam.Sequence.Name, result.Exam.Year.Name)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.Exam.Unit != nil && result.Exam.Type != nil && result.Exam.Unit.Semester != nil && result.Exam.Year != nil {
				message = fmt.Sprintf("Results exam for unit %s: %s %s %s", result.Exam.Unit.Name, result.Exam.Type.Name, result.Exam.Unit.Semester.Name, result.Exam.Year.Name)
			}
		}
		allUsers := append(userStudents, userParents...)
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

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.ResultRequest,
) (result *model.Result, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.Result
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
	item := &model.Result{
		SchoolID:  newRequest.SchoolID,
		StudentID: newRequest.StudentID,
		ExamID:    newRequest.ExamID,

		Score: newRequest.Score,
	}

	// Check exam
	foundExam, err := service.ExamService.Repository.GetByID(item.ExamID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundExam == nil || foundExam.ID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("exam (not found)")
		return

	}

	// Check if the student is enrolled
	var tempClassID, tempLevelDomainID int64
	if foundExam.ClassSubject != nil {
		tempClassID = foundExam.ClassSubject.ClassID
	}
	if foundExam.Unit != nil {
		tempLevelDomainID = foundExam.Unit.LevelDomainID
	}
	foundStudentEnroll, err := service.StudentService.
		Repository.
		GetStudentEnrollBySchoolIDYearIDClassIDLevelDomainIDStudentID(
			item.SchoolID,
			foundExam.YearID,
			tempClassID,
			tempLevelDomainID,
			item.StudentID,
		)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundStudentEnroll == nil || foundStudentEnroll.ID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("student (not enrolled)")
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

func (service *Service) UpdateTable(
	ctxData *types.ContextData,
	id int64,
	request *data.ResultTableRequest,
) (result *model.ResultTable, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ResultTable
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetResultTableByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetResultTableByID(id)
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
	item := &model.ResultTable{
		SchoolID: newRequest.SchoolID,
		ExamID:   newRequest.ExamID,

		Status: newRequest.Status,
	}

	// Check unique
	foundUnique, err := service.Repository.GetResultTableUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameResultTableUniqueObjects(foundUnique, item) && !service.Repository.AreSameResultTableUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateResultTableByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Send message to students and parents
	go func() {
		if result.Exam == nil || result.Exam.ID < 1 || result.Status != constants.RESULT_STATUS_PUBLISHED || result.Status == foundItem.Status {
			return
		}
		// Students
		studentEnrollReq := &dataStudent.GetAllStudentEnrollRequest{}
		studentEnrollReq.SchoolID = result.SchoolID
		studentEnrollReq.YearID = result.Exam.YearID
		studentEnrollReq.ClassSubjectID = result.Exam.ClassSubjectID
		studentEnrollReq.UnitID = result.Exam.UnitID
		userStudents, errUsers := serviceHelperUser.GetAllUserForStudentEnroll(studentEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for student enroll", zap.Error(errUsers))
			return
		}
		// Parents
		parentEnrollReq := &dataParent.GetAllParentStudentRequest{}
		parentEnrollReq.SchoolID = result.SchoolID
		parentEnrollReq.YearID = result.Exam.YearID
		parentEnrollReq.ClassSubjectID = result.Exam.ClassSubjectID
		parentEnrollReq.UnitID = result.Exam.UnitID
		userParents, errUsers := serviceHelperUser.GetAllUserForParentStudent(parentEnrollReq)
		if errUsers != nil {
			helpers.Logger.Error("Error getting users for parent student", zap.Error(errUsers))
			return
		}
		var title, message string
		switch result.Status {
		case constants.RESULT_STATUS_PUBLISHED:
			title = "Available results for exam"
		}
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.Exam.ClassSubject != nil && result.Exam.ClassSubject.Subject != nil && result.Exam.ClassSubject.Class != nil && result.Exam.Type != nil && result.Exam.Sequence != nil && result.Exam.Year != nil {
				message = fmt.Sprintf("Results exam for subject %s %s: %s %s %s", result.Exam.ClassSubject.Subject.Name, result.Exam.ClassSubject.Class.Name, result.Exam.Type.Name, result.Exam.Sequence.Name, result.Exam.Year.Name)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.Exam.Unit != nil && result.Exam.Type != nil && result.Exam.Unit.Semester != nil && result.Exam.Year != nil {
				message = fmt.Sprintf("Results exam for unit %s: %s %s %s", result.Exam.Unit.Name, result.Exam.Type.Name, result.Exam.Unit.Semester.Name, result.Exam.Year.Name)
			}
		}
		allUsers := append(userStudents, userParents...)
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

func (service *Service) Delete(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.Result
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

func (service *Service) DeleteTable(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ResultTable
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetResultTableByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetResultTableByID(id)
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
	affectedRows, err = service.Repository.DeleteResultTableByID(id)
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

func (service *Service) DeleteMultipleTable(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleResultTableByID(list, ctxData.Jwt.SchoolID)
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
) (result *model.Result, errCode int, err error) {
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

func (service *Service) GetTable(
	ctxData *types.ContextData,
	id int64,
) (result *model.ResultTable, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetResultTableByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetResultTableByID(id)
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
) (result []model.Result, errCode int, err error) {
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
		if ctxData.User.Feature != constants.FeatureAdmin && ctxData.User.Feature != constants.FeatureDirector && ctxData.User.Feature != constants.FeatureTeacher {
			newRequest.TableStatusList = []string{constants.EXAM_STATUS_ONLINE, constants.EXAM_STATUS_RESULTS}
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

func (service *Service) GetAllTable(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllResultTableRequest,
) (result []model.ResultTable, errCode int, err error) {
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
		if ctxData.User.Feature != constants.FeatureAdmin && ctxData.User.Feature != constants.FeatureDirector && ctxData.User.Feature != constants.FeatureTeacher {
			newRequest.ExamStatusList = []string{constants.EXAM_STATUS_ONLINE, constants.EXAM_STATUS_RESULTS}
		}
	}

	// Get
	result, err = service.Repository.GetAllResultTable(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
