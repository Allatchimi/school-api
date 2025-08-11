package health

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/config"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

func (service *Service) HealthLive(ctxData *types.ContextData) (result bool, errCode int, err error) {
	result = true
	return
}

func (service *Service) HealthDepencencies(ctxData *types.ContextData) (result bool, errCode int, err error) {
	// Check postgres
	_, err = service.Repository.GetRole()
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("postgres")
		result = false

		// Retry to connect
		errConnect := config.ConnectDatabase()
		if errConnect != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage("reconnect postgres")
			result = false
			return
		}
		return
	}

	// Check redis
	err = config.CheckRedis()
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("redis")
		result = false

		// Retry to connect
		errConnect := config.ConnectRedis()
		if errConnect != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage("reconnect redis")
			result = false
			return
		}
		return
	}

	result = true
	return
}
