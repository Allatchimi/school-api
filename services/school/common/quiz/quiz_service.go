package quiz

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
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
	var canCreate = false

	if !canCreate {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}
	// Insert the quiz
	createdQuiz, _ := service.Repository.Create(
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
	if createdQuiz == nil || createdQuiz.ID <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Insert every question
	if len(request.Questions) > 0 {
		for _, question := range request.Questions {
			// Insert question
			createdQuestion, _ := service.Repository.CreateQuizQuestion(
				&model.QuizQuestion{
					Title:       question.Title,
					Description: question.Description,
					QuizID:      createdQuiz.ID,
				},
			)
			if createdQuestion == nil || createdQuestion.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
			if len(question.Options) > 0 {
				for _, option := range question.Options {
					// Insert option
					createdOption, _ := service.Repository.CreateQuizQuestionOption(
						&model.QuizQuestionOption{
							Title:          option.Title,
							Description:    option.Description,
							QuizQuestionID: createdQuestion.ID,
						},
					)
					if createdOption == nil || createdOption.ID <= 0 {
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

func (service *Service) CreateAttempt(inputJwtToken *types.JwtToken, request *data.QuizAttemptRequest) (result *model.QuizAttempt, errCode int, err error) {
	// TODO
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.QuizRequest) (result *model.Quiz, errCode int, err error) {
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

	// Update quiz
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
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
	affectedRows, err = service.Repository.DeleteMultiple(list)
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
	return
}

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Quiz, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllAttempts(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, quizID int64) (result []model.QuizAttempt, errCode int, err error) {
	result, err = service.Repository.GetAllAttempts(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
