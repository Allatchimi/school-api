package quiz

import (
	"api/common/helpers"
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
		helpers.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateAttempt(
	ctx *context.Context,
	input *struct {
		Body data.QuizAttemptRequest
	},
) (result *model.QuizAttempt, errCode int, err error) {
	result, errCode, err = controller.Service.CreateAttempt(
		helpers.GetJwtContext(ctx),
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
		helpers.GetJwtContext(ctx),
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
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteMultiple(helpers.GetJwtContext(ctx), input.Body.List)
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
	quiz, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
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
	},
) (result *data.QuizResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	quizList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
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

func (controller *Controller) GetAllAttempts(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.QuizID
	},
) (result *data.QuizAttemptResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	quizAttemptList, errCode, err := controller.Service.GetAllAttempts(helpers.GetJwtContext(ctx), newFilter, newPagination, input.QuizID.ID)
	if err != nil {
		return
	}
	result = &data.QuizAttemptResponseList{
		Data: model.ToQuizAttemptResponseList(quizAttemptList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
