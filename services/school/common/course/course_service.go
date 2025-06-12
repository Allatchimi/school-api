package course

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	common_svc_permission "api/services/common"
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

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.CourseRequest) (result *model.Course, errCode int, err error) {
	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.UserID,
		request.SchoolID,
		request.YearID,
		request.ClassSubjectID,
		request.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Insert the course
	result, err = service.Repository.Create(&model.Course{
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

func (service *Service) CreateComment(inputJwtToken *types.JwtToken, id int64, item *model.CourseComment) (result *model.CourseComment, errCode int, err error) {
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

	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.UserID,
		foundItem.SchoolID,
		foundItem.YearID,
		foundItem.ClassSubjectID,
		foundItem.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
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

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.CourseRequest) (result *model.Course, errCode int, err error) {
	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.UserID,
		request.SchoolID,
		request.YearID,
		request.ClassSubjectID,
		request.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

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

func (service *Service) UpdateComment(inputJwtToken *types.JwtToken, id int64, item *model.CourseComment) (result *model.CourseComment, errCode int, err error) {
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

	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.UserID,
		foundItem.Course.SchoolID,
		foundItem.Course.YearID,
		foundItem.Course.ClassSubjectID,
		foundItem.Course.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}
	if foundItem.UserID != inputJwtToken.UserID {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return

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

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
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

	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.UserID,
		foundItem.SchoolID,
		foundItem.YearID,
		foundItem.ClassSubjectID,
		foundItem.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
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

func (service *Service) DeleteComment(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
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

	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.UserID,
		foundItem.Course.SchoolID,
		foundItem.Course.YearID,
		foundItem.Course.ClassSubjectID,
		foundItem.Course.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}
	if foundItem.UserID != inputJwtToken.UserID {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Course, errCode int, err error) {
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

	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.UserID,
		result.SchoolID,
		result.YearID,
		result.ClassSubjectID,
		result.UnitID,
	)
	if !canAccess {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}
	return
}

func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
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
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllCourseCommentRequest,
) (result []model.CourseComment, errCode int, err error) {
	result, err = service.Repository.GetAllCourseComment(
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
