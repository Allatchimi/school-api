package quiz

import (
	"net/http"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	common_svc_permission "api/services/common"
	"api/services/school/common/quiz/data"
	"api/services/school/common/quiz/model"
	"api/services/user/user"
)

type Service struct {
	Repository     *Repository
	UserRepository *user.Repository
}

const MODEL_NAME = "quiz"
const DEFAULT_ERROR_MESSAGE = "interact with quiz model"

func NewService(repository *Repository, userRepository *user.Repository) *Service {
	return &Service{
		Repository:     repository,
		UserRepository: userRepository,
	}
}

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.QuizRequest) (result *model.Quiz, errCode int, err error) {
	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.RoleID,
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

	// Insert the quiz
	createdQuiz, err := service.Repository.Create(
		&model.Quiz{
			Title:          request.Title,
			Description:    request.Description,
			StartDate:      request.StartDate,
			EndDate:        request.EndDate,
			SchoolID:       request.SchoolID,
			YearID:         request.SchoolID,
			UnitID:         request.UnitID,
			ClassSubjectID: request.ClassSubjectID,
		},
	)
	if err != nil || createdQuiz == nil || createdQuiz.ID <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Insert every question and option
	if len(request.Questions) > 0 {
		for _, question := range request.Questions {
			// Insert question
			createdQuestion, errQt := service.Repository.CreateQuizQuestion(
				&model.QuizQuestion{
					Title:       question.Title,
					Description: question.Description,
					QuizID:      createdQuiz.ID,
				},
			)
			if errQt != nil || createdQuestion == nil || createdQuestion.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
			if len(question.Options) > 0 {
				for _, option := range question.Options {
					// Insert option
					createdOption, errOpt := service.Repository.CreateQuizQuestionOption(
						&model.QuizQuestionOption{
							Title:          option.Title,
							Description:    option.Description,
							QuizQuestionID: createdQuestion.ID,
						},
					)
					if errOpt != nil || createdOption == nil || createdOption.ID <= 0 {
						errCode = http.StatusInternalServerError
						err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
						return
					}
				}
			}
		}
	}

	result, err = service.Repository.GetByID(createdQuiz.ID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateAttempt(inputJwtToken *types.JwtToken, quizID int64, request *data.QuizAttemptRequest) (result *model.QuizAttempt, errCode int, err error) {
	// TODO
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.QuizRequest) (result *model.Quiz, errCode int, err error) {
	// Check if the user can access
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.RoleID,
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

	// Check if quiz exists
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

	// Check if the quiz has'nt started
	if utils.AreDateEquals((foundItem.StartDate).UTC(), time.Now().UTC()) {
		errCode = http.StatusConflict
		err = constants.Http409ConflictErrorMessage()
		return
	}

	// Update the quiz
	updatedItem, err := service.Repository.UpdateByID(id, &model.Quiz{
		Title:       request.Title,
		Description: request.Description,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,

		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,
	})
	if err != nil || updatedItem == nil || updatedItem.ID <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Delete all quiz questions and options
	var errDelete error
	questionsIDs := make([]int64, len(foundItem.Questions))
	for i := 0; i < len(questionsIDs); i++ {
		questionsIDs[i] = foundItem.Questions[i].ID

		// Delete options for this question
		optionsIDs := make([]int64, len(foundItem.Questions[i].Options))
		for j := 0; j < len(optionsIDs); j++ {
			optionsIDs[i] = foundItem.Questions[i].Options[j].ID
		}
		_, errDeleteOpt := service.Repository.DeleteMultipleQuizQuestionOptionByID(optionsIDs)
		errDelete = errDeleteOpt
	}
	_, errDeleteQt := service.Repository.DeleteMultipleQuizQuestionByID(questionsIDs)
	errDelete = errDeleteQt
	if errDelete != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert every question and option
	if len(request.Questions) > 0 {
		for _, question := range request.Questions {
			// Insert question
			createdQuestion, errQt := service.Repository.CreateQuizQuestion(
				&model.QuizQuestion{
					Title:       question.Title,
					Description: question.Description,
					QuizID:      updatedItem.ID,
				},
			)
			if errQt != nil || createdQuestion == nil || createdQuestion.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
			if len(question.Options) > 0 {
				for _, option := range question.Options {
					// Insert option
					createdOption, errOpt := service.Repository.CreateQuizQuestionOption(
						&model.QuizQuestionOption{
							Title:          option.Title,
							Description:    option.Description,
							QuizQuestionID: createdQuestion.ID,
						},
					)
					if errOpt != nil || createdOption == nil || createdOption.ID <= 0 {
						errCode = http.StatusInternalServerError
						err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
						return
					}
				}
			}
		}
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	// Check if the user can access
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
	canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
		inputJwtToken.RoleID,
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

func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	// Check if the user can access
	for i := 0; i < len(list); i++ {
		foundItem, errCheck := service.Repository.GetByID(list[i])
		if errCheck != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		if foundItem == nil || foundItem.ID < 0 {
			errCode = http.StatusNotFound
			err = constants.Http404ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		canAccess := common_svc_permission.CanAccessBySchoolYearClassSubjectUnit(
			inputJwtToken.RoleID,
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
	}

	// Delete
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Quiz, errCode int, err error) {
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
		inputJwtToken.RoleID,
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
	clause *types.FilterlSchoolYearUnitClassSubjectRequest,
) (result []model.Quiz, errCode int, err error) {
	var schoolID, yearID, classSubjectID, unitID int64
	if clause != nil {
		schoolID = clause.SchoolID
		yearID = clause.YearID
		classSubjectID = clause.ClassSubjectID
		unitID = clause.UnitID
	}
	result, err = service.Repository.GetAll(filter, pagination, schoolID, yearID, classSubjectID, unitID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllQuizAttempt(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	quizID int64,
	clause *types.FilterlSchoolYearUnitClassSubjectRequest,
) (result []model.QuizAttempt, errCode int, err error) {
	var schoolID, yearID, classSubjectID, unitID int64
	if clause != nil {
		schoolID = clause.SchoolID
		yearID = clause.YearID
		classSubjectID = clause.ClassSubjectID
		unitID = clause.UnitID
	}
	result, err = service.Repository.GetAllQuizAttempt(filter, pagination, quizID, schoolID, yearID, classSubjectID, unitID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
