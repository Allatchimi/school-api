package school

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/school/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "school"
const DEFAULT_ERROR_MESSAGE = "interact with school model"

func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.School) (result *model.School, errCode int, err error) {
	// Create info
	newInfo, err := service.Repository.CreateInfo(&model.SchoolInfo{
		FullName:    item.Info.FullName,
		Description: item.Info.Description,
		Slogan:      item.Info.Slogan,

		PhoneNumber1: item.Info.PhoneNumber1,
		PhoneNumber2: item.Info.PhoneNumber2,
		PhoneNumber3: item.Info.PhoneNumber3,

		Email1: item.Info.Email1,
		Email2: item.Info.Email2,
		Email3: item.Info.Email3,

		Founder:   item.Info.Founder,
		FoundedAt: item.Info.FoundedAt,

		Address:           item.Info.Address,
		LocationLongitude: item.Info.LocationLongitude,
		LocationLatitude:  item.Info.LocationLatitude,

		Logo: item.Info.Logo,

		Image1: item.Info.Image1,
		Image2: item.Info.Image2,
		Image3: item.Info.Image3,
		Image4: item.Info.Image4,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create config
	newConfig, err := service.Repository.CreateConfig(&model.SchoolConfig{
		EmailDomain:  item.Config.EmailDomain,
		ColorPrimary: item.Config.ColorPrimary,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create school
	result, err = service.Repository.Create(&model.School{
		Name: item.Name,
		Type: item.Type,

		SchoolConfigID: newConfig.ID,
		SchoolInfoID:   newInfo.ID,
	})
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage(MODEL_NAME)
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.School) (result *model.School, errCode int, err error) {
	// Check if school exists
	foundItem, err := service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update school
	result, err = service.Repository.Update(id, item)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage(MODEL_NAME)
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update config
	newConfig, _, err := service.UpdateConfig(inputJwtToken, foundItem.SchoolConfigID, item.Config)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	result.Config = newConfig

	// Update info
	newInfo, errCode, err := service.UpdateInfo(inputJwtToken, foundItem.SchoolInfoID, item.Info)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	result.Info = newInfo
	return
}

func (service *Service) UpdateInfo(inputJwtToken *types.JwtToken, id int64, item *model.SchoolInfo) (result *model.SchoolInfo, errCode int, err error) {
	// Check if school info already exists
	foundItem, err := service.Repository.GetInfoById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateInfo(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateConfig(inputJwtToken *types.JwtToken, id int64, item *model.SchoolConfig) (result *model.SchoolConfig, errCode int, err error) {
	// Check if school config already exists
	foundItem, err := service.Repository.GetConfigById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateConfig(id, item)
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
		err = constants.Http404ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.School, errCode int, err error) {
	result, err = service.Repository.GetById(id)
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, typeName string) (result []model.School, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, typeName)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
