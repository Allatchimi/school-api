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
		return
	}

	// Check redis
	err = config.CheckRedis()
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("redis")
		result = false
		return
	}

	result = true
	return
}
