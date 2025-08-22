package initialize

import (
	"net/http"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	serviceHelperFeature "api/services/helper/feature"
	"api/services/others/initialize/data"
	"api/services/school/common/school"
	"api/services/school/common/year"
	dataYear "api/services/school/common/year/data"
	modelYear "api/services/school/common/year/model"

	"go.uber.org/zap"
)

type Service struct {
	SchoolService *school.Service
	YearService   *year.Service
}

func NewService(schoolService *school.Service, yearService *year.Service) *Service {
	return &Service{SchoolService: schoolService, YearService: yearService}
}

const DEFAULT_ERROR_MESSAGE = "initialize data"

func (service *Service) GetInitial(
	ctxData *types.ContextData,
) (result *data.InitializeResponse, errCode int, err error) {
	// Check school
	newRequest := &dataYear.GetAllRequest{}
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check feature
	var okCheck bool = false
	if ctxData.User.Feature != constants.FeatureAdmin {
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
	}

	// Get school
	school, err := service.SchoolService.Repository.GetByID(newRequest.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("")
		return
	}

	// Get years
	var tempYears []modelYear.Year = make([]modelYear.Year, 0)
	if okCheck {
		tempYears, err = service.YearService.Repository.GetAll(nil, nil, newRequest)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		helpers.Logger.Info("Initial years", zap.Int("count", len(tempYears)))
	}

	// Update response
	years := modelYear.ToResponseList(tempYears)
	result = &data.InitializeResponse{School: school.ToPublicResponse(), Years: years}
	return
}
