package initialize

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/others/initialize/data"
	"api/services/school/common/school"
	"api/services/school/common/year"
	dataYear "api/services/school/common/year/data"
	modelYear "api/services/school/common/year/model"
)

type Service struct {
	SchoolService *school.Service
	YearService   *year.Service
}

func NewService(schoolService *school.Service, yearService *year.Service) *Service {
	return &Service{SchoolService: schoolService, YearService: yearService}
}

func (service *Service) GetInitial(
	ctxData *types.ContextData,
) (result *data.InitializeResponse, errCode int, err error) {
	school, err := service.SchoolService.Repository.GetByID(ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("")
		return
	}
	tempYears, err := service.YearService.Repository.GetAll(nil, &types.Pagination{
		CurrentPage: 1,
		Limit:       10,
		Offset:      0,
	}, &dataYear.GetAllRequest{
		SchoolID: ctxData.Jwt.SchoolID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("")
		return
	}
	years := modelYear.ToResponseList(tempYears)
	result = &data.InitializeResponse{School: school.ToPublicResponse(), Years: years}
	return
}
