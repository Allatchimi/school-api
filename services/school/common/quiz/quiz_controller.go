package quiz

import (
	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/quiz/data"
	"api/services/school/common/quiz/model"
	"context"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Create(
	ctx *context.Context,
	input *struct {
		Body data.QuizRequest
	},
) (result *model.Quiz, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateAnswer(
	ctx *context.Context,
	input *struct {
		data.QuizID
		Body data.QuizAnswerRequest
	},
) (errCode int, err error) {
	errCode, err = controller.Service.CreateAnswer(
		httpHelper.GetContextData(ctx),
		input.QuizID.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.QuizID
		Body data.QuizRequest
	},
) (result *model.Quiz, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateSolution(
	ctx *context.Context,
	input *struct {
		data.QuizID
		Body data.QuizSolutionRequest
	},
) (result *model.Quiz, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateSolution(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.QuizID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteMultiple(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultiple(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.QuizID
	},
) (result *model.Quiz, errCode int, err error) {
	quiz, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = quiz
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.QuizResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	quizList, errCode, err := controller.Service.GetAll(
		httpHelper.GetContextData(ctx), newFilter, newPagination,
		&input.GetAllRequest,
	)
	if err != nil {
		return
	}
	result = &data.QuizResponseList{
		Data: model.ToQuizResponseList(quizList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllQuizResult(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.QuizID
		data.GetAllQuizAnswerRequest
	},
) (result *data.QuizResultResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	quizAnswerList, errCode, err := controller.Service.GetAllQuizAnswer(
		httpHelper.GetContextData(ctx), newFilter, newPagination,
		&input.GetAllQuizAnswerRequest,
	)
	if err != nil {
		return
	}
	result = &data.QuizResultResponseList{
		Data: model.ToQuizAnswerResultResponseList(quizAnswerList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
