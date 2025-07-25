package quiz

import (
	"api/common/constants"
	"api/common/types"
	"api/services/school/common/quiz/data"
	"api/services/school/common/quiz/model"
	"net/http"
)

type Service struct {
	Repository *Repository
}

const MODEL_NAME = "quiz"
const DEFAULT_ERROR_MESSAGE = "interact with quiz model"

func NewService(repository *Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.QuizRequest,
) (result *model.Quiz, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format item
	item := &model.Quiz{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.SchoolID,
		ClassSubjectID: newRequest.ClassSubjectID,
		UnitID:         newRequest.UnitID,

		Title:       newRequest.Title,
		Description: newRequest.Description,
		Status:      newRequest.Status,
	}
	if newRequest.ClassSubjectID < 1 && newRequest.UnitID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("class subject id or unit id(you should provide one of these fields)")
		return
	}

	// Insert the quiz
	createdQuiz, err := service.Repository.Create(item)
	if err != nil || createdQuiz == nil || createdQuiz.ID <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Insert every question and option
	if len(newRequest.Questions) > 0 {
		for _, question := range newRequest.Questions {
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

func (service *Service) CreateAnswer(
	ctxData *types.ContextData,
	id int64,
	request *data.QuizAnswerRequest,
) (errCode int, err error) {
	// Check school
	newRequest := *request

	// Check if the item exists
	var foundItem *model.Quiz
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

	// Check if the student have already submitted and answer
	questionsIDs := make([]int64, len(newRequest.Answers))
	for i := range newRequest.Answers {
		questionsIDs[i] = newRequest.Answers[i].QuestionID
	}
	foundAnswer, tempErrFoundAnswer := service.Repository.GetAllQuizAnswerByStudentIDQuizQuestionIDs(
		newRequest.StudentID,
		questionsIDs,
	)
	if tempErrFoundAnswer != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundAnswer != nil && foundAnswer.ID > 0 {
		errCode = http.StatusLocked
		err = constants.Http423LockedErrorMessage()
		return
	}

	// Add answers
	for _, answer := range newRequest.Answers {
		_, tempAddErr := service.Repository.CreateQuizAnswer(
			&model.QuizAnswer{
				StudentID:            newRequest.StudentID,
				QuizQuestionID:       answer.QuestionID,
				QuizQuestionOptionID: answer.OptionID,
			},
		)
		if tempAddErr != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}
	return
}

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.QuizRequest,
) (result *model.Quiz, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.Quiz
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

	// Update the quiz
	updatedItem, err := service.Repository.UpdateByID(id, &model.Quiz{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		UnitID:         newRequest.UnitID,

		Title:       newRequest.Title,
		Description: newRequest.Description,
		Status:      newRequest.Status,
	})
	if err != nil || updatedItem == nil || updatedItem.ID <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Delete all quiz questions and options
	var errDelete error
	questionsIDs := make([]int64, len(foundItem.Questions))
	for i := range questionsIDs {
		questionsIDs[i] = foundItem.Questions[i].ID

		// Delete options for this question
		optionsIDs := make([]int64, len(foundItem.Questions[i].Options))
		for j := range optionsIDs {
			optionsIDs[i] = foundItem.Questions[i].Options[j].ID
		}
		_, errDelete = service.Repository.DeleteMultipleQuizQuestionOptionByID(optionsIDs)
	}
	if errDelete != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	_, errDeleteQt := service.Repository.DeleteMultipleQuizQuestionByID(questionsIDs)
	errDelete = errDeleteQt
	if errDelete != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert every question and option
	if len(newRequest.Questions) > 0 {
		for _, question := range newRequest.Questions {
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

func (service *Service) UpdateSolution(
	ctxData *types.ContextData,
	quizID int64,
	request *data.QuizSolutionRequest,
) (result *model.Quiz, errCode int, err error) {
	// Check school
	newRequest := *request

	// Check if the item exists
	var foundItem *model.Quiz
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(quizID, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(quizID)
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

	// Update the quiz question solution
	if len(newRequest.Solutions) > 0 {
		for _, solution := range newRequest.Solutions {
			// Update question solution
			updatedQuestion, errUpdate := service.Repository.UpdateQuizQuestionSolutionByID(
				solution.QuestionID,
				&model.QuizQuestion{
					SolutionID: solution.OptionID,
				},
			)
			if errUpdate != nil || updatedQuestion == nil || updatedQuestion.ID <= 0 {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
				return
			}
		}
	}

	// Reload the quiz
	result, err = service.Repository.GetByID(quizID)

	return
}

func (service *Service) Delete(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.Quiz
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
) (result *model.Quiz, errCode int, err error) {
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
) (result []model.Quiz, errCode int, err error) {
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

func (service *Service) GetAllQuizAnswer(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllQuizAnswerRequest,
) (result []model.QuizAnswer, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllQuizAnswer(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
