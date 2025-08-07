package class

import (
	"net/http"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/meeting"
	dataMeeting "api/services/school/common/meeting/data"
	"api/services/school/common/school"
	"api/services/school/highschool/class/data"
	"api/services/school/highschool/class/model"
)

type Service struct {
	Repository     *Repository
	SchoolService  *school.Service
	MeetingService *meeting.Service
}

func NewService(
	repository *Repository,
	schoolService *school.Service,
	meetingService *meeting.Service,
) *Service {
	return &Service{
		Repository:     repository,
		SchoolService:  schoolService,
		MeetingService: meetingService,
	}
}

const MODEL_NAME = "class"
const DEFAULT_ERROR_MESSAGE = "interact with class model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.ClassRequest,
) (result *model.HighschoolClass, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.HighschoolClass{
		SchoolID:    newRequest.SchoolID,
		SpecialtyID: newRequest.SpecialtyID,

		Name:         newRequest.Name,
		Description:  newRequest.Description,
		Fees:         newRequest.Fees,
		Program:      newRequest.Program,
		Requirements: newRequest.Requirements,
		IsValid:      newRequest.IsValid,
	}

	// Check if the school type is highschool
	foundSchool, err := service.SchoolService.Repository.GetByID(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
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

	// Check invalid date
	if !item.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateClassSubject(
	ctxData *types.ContextData,
	request *data.ClassSubjectRequest,
) (result *model.HighschoolClassSubject, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.HighschoolClassSubject{
		SchoolID:  newRequest.SchoolID,
		ClassID:   newRequest.ClassID,
		SubjectID: newRequest.SubjectID,

		Coefficient:  newRequest.Coefficient,
		Program:      newRequest.Program,
		Requirements: newRequest.Requirements,
		IsValid:      newRequest.IsValid,
	}

	// Check unique
	foundUnique, err := service.Repository.GetClassSubjectUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameClassSubjectUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Check invalid date
	if !item.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	}

	// Create
	result, err = service.Repository.CreateClassSubject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create the meeting
	go func() {
		service.MeetingService.Create(ctxData, &dataMeeting.MeetingRoomRequest{
			SchoolID:       result.SchoolID,
			ClassSubjectID: result.ID,
		})
	}()
	return
}

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.ClassRequest,
) (result *model.HighschoolClass, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.HighschoolClass
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
	item := &model.HighschoolClass{
		SchoolID:    newRequest.SchoolID,
		SpecialtyID: newRequest.SpecialtyID,

		Name:         newRequest.Name,
		Description:  newRequest.Description,
		Fees:         newRequest.Fees,
		Program:      newRequest.Program,
		Requirements: newRequest.Requirements,
		IsValid:      newRequest.IsValid,
	}

	// Check if the school type is highschool
	if foundItem.School.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
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

	// Check invalid date
	if !item.IsValid && foundItem.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	} else if item.IsValid && !foundItem.IsValid {
		item.InvalidDate = nil
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

func (service *Service) UpdateClassSubject(
	ctxData *types.ContextData,
	id int64,
	request *data.ClassSubjectRequest,
) (result *model.HighschoolClassSubject, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.HighschoolClassSubject
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetClassSubjectByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetClassSubjectByID(id)
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
	item := &model.HighschoolClassSubject{
		SchoolID:  newRequest.SchoolID,
		ClassID:   newRequest.ClassID,
		SubjectID: newRequest.SubjectID,

		Coefficient:  newRequest.Coefficient,
		Program:      newRequest.Program,
		Requirements: newRequest.Requirements,
		IsValid:      newRequest.IsValid,
	}

	// Check if the school type is highschool
	if foundItem.School.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetClassSubjectUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameClassSubjectUniqueObjects(foundUnique, item) && !service.Repository.AreSameClassSubjectUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Check invalid date
	if !item.IsValid && foundItem.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	} else if item.IsValid && !foundItem.IsValid {
		item.InvalidDate = nil
	}

	// Update
	result, err = service.Repository.UpdateClassSubjectByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create the meeting
	go func() {
		service.MeetingService.Create(ctxData, &dataMeeting.MeetingRoomRequest{
			SchoolID:       result.SchoolID,
			ClassSubjectID: result.ID,
		})
	}()
	return
}

func (service *Service) Delete(
	ctxData *types.ContextData,
	id int64) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.HighschoolClass
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

func (service *Service) DeleteClassSubject(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.HighschoolClassSubject
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetClassSubjectByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetClassSubjectByID(id)
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
	affectedRows, err = service.Repository.DeleteClassSubjectByID(id)
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

func (service *Service) DeleteMultipleClassSubject(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleClassSubjectByID(list, ctxData.Jwt.SchoolID)
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
) (result *model.HighschoolClass, errCode int, err error) {
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

func (service *Service) GetClassSubject(
	ctxData *types.ContextData,
	id int64,
) (result *model.HighschoolClassSubject, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetClassSubjectByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetClassSubjectByID(id)
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
) (result []model.HighschoolClass, errCode int, err error) {
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

func (service *Service) GetAllClassSubject(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllClassSubjectRequest,
) (result []model.HighschoolClassSubject, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllClassSubject(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
