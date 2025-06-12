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

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.QuizRequest) (result *model.Quiz, errCode int, err error) {
	// Insert the quiz
	createdQuiz, err := service.Repository.Create(
		&model.Quiz{
			SchoolID:       request.SchoolID,
			YearID:         request.SchoolID,
			ClassSubjectID: request.ClassSubjectID,
			UnitID:         request.UnitID,

			Title:       request.Title,
			Description: request.Description,
			Status:      request.Status,
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

func (service *Service) CreateAnswer(inputJwtToken *types.JwtToken, id int64, request *data.QuizAnswerRequest) (errCode int, err error) {
	// Load the quiz
	foundQuiz, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundQuiz == nil || foundQuiz.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check if the student have already submitted and answer
	questionsIDs := make([]int64, len(request.Answers))
	for i := range request.Answers {
		questionsIDs[i] = request.Answers[i].QuestionID
	}
	foundAnswer, tempErrFoundAnswer := service.Repository.GetAllQuizAnswerByStudentIDQuizQuestionIDs(
		request.StudentID,
		questionsIDs,
	)
	if tempErrFoundAnswer != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundAnswer != nil && foundAnswer.ID > 0 {
		errCode = http.StatusForbidden
		err = constants.Http409ConflictErrorMessage()
		return
	}

	// Add answers
	for _, answer := range request.Answers {
		_, tempAddErr := service.Repository.CreateQuizAnswer(
			&model.QuizAnswer{
				StudentID:            request.StudentID,
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

	// Update the quiz
	updatedItem, err := service.Repository.UpdateByID(id, &model.Quiz{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,

		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
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

func (service *Service) UpdateSolution(inputJwtToken *types.JwtToken, id int64, request *data.QuizSolutionRequest) (result *model.Quiz, errCode int, err error) {
	// Load the quiz
	foundQuiz, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundQuiz == nil || foundQuiz.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Update the quiz question solution
	if len(request.Solutions) > 0 {
		for _, solution := range request.Solutions {
			// Update question solution
			updatedQuestion, errUpdate := service.Repository.UpdateQuizQuestionSolutionByID(
				solution.QuestionID,
				&model.QuizQuestion{
					SolutionID: solution.SolutionID,
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
	result, err = service.Repository.GetByID(id)

	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
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

func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, selection []int64) (affectedRows int64, errCode int, err error) {
	// Delete
	affectedRows, err = service.Repository.DeleteMultipleByID(selection)
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

func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	clause *data.GetAllRequest,
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

func (service *Service) GetAllQuizAnswer(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	clause *data.GetAllQuizAnswerRequest,
) (result []model.QuizAnswer, errCode int, err error) {
	var quizID, quizQuestionID, studentID int64
	if clause != nil {
		quizID = clause.QuizID
		quizQuestionID = clause.QuizQuestionID
		studentID = clause.StudentID
	}
	result, err = service.Repository.GetAllQuizAnswer(filter, pagination, quizID, quizQuestionID, studentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllQuizQuestionOption(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	clause *data.GetAllQuizQuestionOptionRequest,
) (result []model.QuizQuestionOption, errCode int, err error) {
	var quizQuestionID int64
	if clause != nil {
		quizQuestionID = clause.QuizQuestionID
	}
	result, err = service.Repository.GetAllQuizQuestionOption(filter, pagination, quizQuestionID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
