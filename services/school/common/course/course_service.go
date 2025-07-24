package course

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	serviceHelper "api/services/helper"
	"api/services/school/common/course/data"
	"api/services/school/common/course/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "course"
const DEFAULT_ERROR_MESSAGE = "interact with course model"

func (service *Service) Create(ctxData *types.ContextData, request *data.CourseRequest) (result *model.Course, errCode int, err error) {
	// Format request
	item := &model.Course{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,

		Title:       request.Title,
		Description: request.Description,
		Content:     request.Content,
	}

	// Insert the course
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert every document and video
	if len(request.Documents) > 0 {
		for _, tempIterator := range request.Documents {
			tempItem, tempErr := service.Repository.CreateCourseDocument(
				&model.CourseDocument{
					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					CourseID:    result.ID,
				},
			)
			if tempErr != nil || tempItem == nil || tempItem.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}
	if len(request.Videos) > 0 {
		for _, tempIterator := range request.Videos {
			tempItem, tempErr := service.Repository.CreateCourseDocument(
				&model.CourseDocument{
					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					CourseID:    result.ID,
				},
			)
			if tempErr != nil || tempItem == nil || tempItem.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}

	// Refetch the course
	foundItem, err := service.Repository.GetByID(result.ID)
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
	result = foundItem
	return
}

func (service *Service) CreateComment(ctxData *types.ContextData, id int64, request *data.CourseCommentRequest) (result *model.CourseComment, errCode int, err error) {
	// Check if the course exists
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

	// Format request
	item := &model.CourseComment{
		CourseID: id,
		UserID:   ctxData.Jwt.UserID,
		Message:  request.Message,
		Rate:     request.Rate,
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

func (service *Service) Update(ctxData *types.ContextData, id int64, request *data.CourseRequest) (result *model.Course, errCode int, err error) {
	// Check if course exists
	foundItem, err := service.Repository.GetByID(id)
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
	if len(request.Documents) > 0 {
		for _, tempIterator := range request.Documents {
			tempItem, tempErr := service.Repository.CreateCourseDocument(
				&model.CourseDocument{
					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					CourseID:    id,
				},
			)
			if tempErr != nil || tempItem == nil || tempItem.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}
	if len(request.Videos) > 0 {
		for _, tempIterator := range request.Videos {
			tempItem, tempErr := service.Repository.CreateCourseDocument(
				&model.CourseDocument{
					Title:       tempIterator.Title,
					Description: tempIterator.Description,
					CourseID:    id,
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
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,

		Title:       request.Title,
		Description: request.Description,
		Content:     request.Content,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateComment(ctxData *types.ContextData, id int64, request *data.CourseCommentRequest) (result *model.CourseComment, errCode int, err error) {
	// Get the course comment
	foundItem, err := service.Repository.GetCourseCommentByID(id)
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
		Message:  request.Message,
		Rate:     request.Rate,
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

func (service *Service) Delete(ctxData *types.ContextData, id int64) (affectedRows int64, errCode int, err error) {
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

func (service *Service) DeleteComment(ctxData *types.ContextData, id int64) (affectedRows int64, errCode int, err error) {
	// Get the course comment
	foundItem, err := service.Repository.GetCourseCommentByID(id)
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

	// Perform soft delete
	_, err = service.Repository.UpdateCourseCommentByID(
		id,
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

func (service *Service) DeleteMultiple(ctxData *types.ContextData, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleByID(list)
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

func (service *Service) Get(ctxData *types.ContextData, id int64) (result *model.Course, errCode int, err error) {
	// Get the course
	result, err = service.Repository.GetByID(id)
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
	result, err = service.Repository.GetAll(
		filter,
		pagination,
		request,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllComment(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllCourseCommentRequest,
) (result []model.CourseComment, errCode int, err error) {
	// Get user
	_, err = serviceHelper.GetUserByID(ctxData.Jwt.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Proceed by feature
	// if foundUser.Role.Feature != constants.FeatureAdmin {
	// 	newRequest := *request
	// 	newRequest.SchoolID = foundUser.SchoolID
	// 	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	// } else {
	// 	result, err = service.Repository.GetAllCourseComment(
	// 		filter,
	// 		pagination,
	// 		request,
	// 	)
	// }
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
