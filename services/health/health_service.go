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

func (service *Service) GetPostgresHealth(inputJwtToken *types.JwtToken) (result bool, errCode int, err error) {
	_, err = service.Repository.GetHistory()
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("")
		result = false
		return
	}
	result = true
	return
}

func (service *Service) GetRedisHealth(inputJwtToken *types.JwtToken) (result bool, errCode int, err error) {
	err = config.CheckRedis()
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("")
		result = false
		return
	}
	result = true
	return
}
