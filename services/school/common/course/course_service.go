package course

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	serviceHelperFeature "api/services/helper/feature"
	serviceHelperMessage "api/services/helper/message"
	serviceHelperUser "api/services/helper/user"
	"api/services/school/common/course/data"
	"api/services/school/common/course/model"
	dataStudent "api/services/school/common/student/data"
)

type Service struct {
	Repository *Repository
}

func NewService(
	repository *Repository,
) *Service {
	return &Service{
		Repository: repository,
	}
}

const MODEL_NAME = "course"
const DEFAULT_ERROR_MESSAGE = "interact with course model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.CourseRequest,
) (result *model.Course, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.Course{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		UnitID:         newRequest.UnitID,

		Title:       newRequest.Title,
		Description: newRequest.Description,
		Content:     newRequest.Content,
	}
	if newRequest.ClassSubjectID < 1 && newRequest.UnitID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("class subject id or unit id(you should provide one of these fields)")
		return
	}

	// Insert the course
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert every document and video
	if len(newRequest.Documents) > 0 {
		for _, tempIterator := range newRequest.Documents {
			tempItem, tempErr := service.Repository.CreateCourseDocument(
				&model.CourseDocument{
					CourseID: result.ID,

					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					Url:         tempIterator.Url,
				},
			)
			if tempErr != nil || tempItem == nil || tempItem.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}
	if len(newRequest.Videos) > 0 {
		for _, tempIterator := range newRequest.Videos {
			tempItem, tempErr := service.Repository.CreateCourseVideo(
				&model.CourseVideo{
					CourseID: result.ID,

					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					Url:         tempIterator.Url,
				},
			)
			if tempErr != nil || tempItem == nil || tempItem.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}

	// Send message
	go func() {
		stdEnrollReq := &dataStudent.GetAllStudentEnrollRequest{}
		stdEnrollReq.SchoolID = result.SchoolID
		stdEnrollReq.YearID = result.YearID
		stdEnrollReq.ClassSubjectID = result.ClassSubjectID
		stdEnrollReq.UnitID = result.UnitID
		users, errUsers := serviceHelperUser.GetAllUserForStudentEnroll(stdEnrollReq)
		if errUsers != nil || len(users) < 1 {
			return
		}
		var title, message string
		title = "New published course"
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.ClassSubject != nil && result.ClassSubject.Subject != nil {
				message = fmt.Sprintf("%s %s: %s", result.ClassSubject.Subject.Name, result.ClassSubject.Class.Name, result.Title)
			} else {
				message = result.Title
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.Unit != nil {
				message = fmt.Sprintf("%s: %s", result.Unit.Name, result.Title)
			} else {
				message = result.Title
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
			fmt.Sprintf("/dashboard/common/courses/%d", result.ID),
			users,
		)
	}()
	return
}

func (service *Service) CreateCourseComment(
	ctxData *types.ContextData,
	courseID int64,
	request *data.CourseCommentRequest,
) (result *model.CourseComment, errCode int, err error) {
	// Check school
	newRequest := *request

	// Check if the item exists
	var foundItem *model.Course
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(courseID, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(courseID)
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
	item := &model.CourseComment{
		CourseID: courseID,
		UserID:   ctxData.Jwt.UserID,

		Message: newRequest.Message,
		Rate:    newRequest.Rate,
	}

	// Insert course
	result, err = service.Repository.CreateCourseComment(item)
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
	request *data.CourseRequest,
) (result *model.Course, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.Course
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

	// Format item
	if newRequest.ClassSubjectID < 1 && newRequest.UnitID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("class subject id or unit id(you should provide one of these fields)")
		return
	}

	// Delete all course documents and videos
	_, errDelete := service.Repository.DeleteCourseDocumentByCourseID(id)
	if errDelete != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	_, errDelete = service.Repository.DeleteCourseVideoByCourseID(id)
	if errDelete != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert every document and video
	if len(newRequest.Documents) > 0 {
		for _, tempIterator := range newRequest.Documents {
			tempItem, tempErr := service.Repository.CreateCourseDocument(
				&model.CourseDocument{
					CourseID: id,

					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					Url:         tempIterator.Url,
				},
			)
			if tempErr != nil || tempItem == nil || tempItem.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}
	if len(newRequest.Videos) > 0 {
		for _, tempIterator := range newRequest.Videos {
			tempItem, tempErr := service.Repository.CreateCourseVideo(
				&model.CourseVideo{
					CourseID: id,

					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					Url:         tempIterator.Url,
				},
			)
			if tempErr != nil || tempItem == nil || tempItem.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}

	// Update the course
	result, err = service.Repository.UpdateByID(id, &model.Course{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		UnitID:         newRequest.UnitID,

		Title:       newRequest.Title,
		Description: newRequest.Description,
		Content:     newRequest.Content,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateCourseComment(
	ctxData *types.ContextData,
	id int64,
	request *data.CourseCommentRequest,
) (result *model.CourseComment, errCode int, err error) {
	// Check school
	newRequest := *request

	// Check if the item exists
	var foundItem *model.CourseComment
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetCourseCommentByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetCourseCommentByID(id)
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

	// Check if the comment is not deleted
	if foundItem.IsDeleted {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Format request
	item := &model.CourseComment{
		CourseID: id,
		UserID:   ctxData.Jwt.UserID,
		Message:  newRequest.Message,
		Rate:     newRequest.Rate,
	}

	// Update
	result, err = service.Repository.UpdateCourseCommentByID(id, item)
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
	// Get the course
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Perform delete
	affectedRows, err = service.Repository.DeleteByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteCourseComment(
	ctxData *types.ContextData,
	courseID int64,
	commentID int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.CourseComment
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetCourseCommentByIDSchoolID(commentID, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetCourseCommentByID(commentID)
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

	// Check macthing with course id
	if foundItem.CourseID != courseID {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Perform soft delete
	_, err = service.Repository.UpdateCourseCommentByID(
		commentID,
		&model.CourseComment{
			Message:   "",
			Rate:      foundItem.Rate,
			IsDeleted: true,
		},
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	affectedRows = 1
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
) (result *model.Course, errCode int, err error) {
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
) (result []model.Course, errCode int, err error) {
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
	}

	// Get
	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllCourseComment(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	courseID int64,
	request *data.GetAllCourseCommentRequest,
) (result []model.CourseComment, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}
	newRequest.CourseID = courseID

	// Get
	result, err = service.Repository.GetAllCourseComment(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
