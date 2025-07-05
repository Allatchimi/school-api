package model

import (
	"time"

	"gorm.io/gorm"

	"api/common/types"
	"api/common/utils/security"
	"api/services/user/role/model"
	"api/services/user/user/data"
)

type User struct {
	types.BaseGormModel
	RoleID int64       `gorm:"default:null"`
	Role   *model.Role `gorm:"default:null;foreignKey:RoleID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserInfoID int64     `gorm:"default:null"`
	Info       *UserInfo `gorm:"default:null;foreignKey:UserInfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserConfigID int64       `gorm:"default:null"`
	Config       *UserConfig `gorm:"default:null;foreignKey:UserConfigID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Email       string `gorm:"default:null"`
	PhoneNumber uint64 `gorm:"default:null"`
	Password    string `gorm:"default:null"`

	LoginMethod    string     `gorm:"default:null"`
	Provider       string     `gorm:"default:null"`
	ProviderUserID string     `gorm:"default:null"`
	IsActivated    bool       `gorm:"default:false"`
	ActivatedAt    *time.Time `gorm:"default:null"`
}

func (item *User) BeforeCreate(db *gorm.DB) (err error) {
	item.Password, err = security.EncodeArgon2id(item.Password)
	return
}

func (item *User) BeforeUpdate(db *gorm.DB) (err error) {
	item.Password, err = security.EncodeArgon2id(item.Password)
	return
}

func (item *User) ToResponse() *data.UserResponse {
	if item == nil {
		return nil
	}
	resp := &data.UserResponse{}
	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber

	resp.LoginMethod = item.LoginMethod
	resp.Provider = item.Provider
	resp.ProviderUserID = item.ProviderUserID
	resp.IsActivated = item.IsActivated
	resp.ActivatedAt = item.ActivatedAt

	resp.Role = item.Role.ToResponse()
	resp.Info = item.Info.ToResponse()
	resp.Config = item.Config.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *User) ToPublicResponse() *data.UserPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.UserPublicResponse{}
	resp.Email = item.Email

	resp.Role = item.Role.ToPublicResponse()
	resp.Info = item.Info.ToPublicResponse()
	return resp
}

func (item *User) FromGoogleUser(googleUser *types.GoogleUserProfileResponse) {
	item.ProviderUserID = googleUser.ID
	item.Email = googleUser.Email
	item.Info = &UserInfo{
		Username:  googleUser.FullName,
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
		Image:     googleUser.Picture,
	}
}
func (item *User) FromFacebookUser(facebookUser *types.FacebookUserProfileResponse) {
	item.ProviderUserID = facebookUser.ID
	item.Email = facebookUser.Email
	item.Info = &UserInfo{
		Username:  facebookUser.FullName,
		FirstName: facebookUser.FirstName,
		LastName:  facebookUser.LastName,
		Image:     facebookUser.PictureSmall.Data.Url,
	}
}

func ToResponseList(itemList []User) []data.UserResponse {
	resp := make([]data.UserResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
