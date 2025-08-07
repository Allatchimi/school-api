package model

import (
	"time"

	"gorm.io/gorm"

	"api/common/types"
	securityUtil "api/common/utils/security"
	modelSchool "api/services/school/common/school/model"
	modelRole "api/services/user/role/model"
	"api/services/user/user/data"
)

type User struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	RoleID int64           `gorm:"default:null"`
	Role   *modelRole.Role `gorm:"default:null;foreignKey:RoleID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	InfoID int64     `gorm:"default:null"`
	Info   *UserInfo `gorm:"default:null;foreignKey:InfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ConfigID int64       `gorm:"default:null"`
	Config   *UserConfig `gorm:"default:null;foreignKey:ConfigID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Email          string     `gorm:"default:null"`
	PhoneNumber    uint64     `gorm:"default:null"`
	Password       string     `gorm:"default:null"`
	Status         string     `gorm:"default:null"`
	LoginMethod    string     `gorm:"default:null"`
	Provider       string     `gorm:"default:null"`
	ProviderUserID string     `gorm:"default:null"`
	IsActivated    bool       `gorm:"default:false"`
	ActivatedAt    *time.Time `gorm:"default:null"`
}

func (item *User) BeforeCreate(db *gorm.DB) (err error) {
	item.Password, err = securityUtil.EncodeArgon2id(item.Password)
	return
}

func (item *User) BeforeUpdate(db *gorm.DB) (err error) {
	item.Password, err = securityUtil.EncodeArgon2id(item.Password)
	return
}

func (u *User) BeforeDelete(db *gorm.DB) (err error) {
	if u == nil {
		return
	}
	if u.InfoID > 0 {
		db.Unscoped().Delete(&UserInfo{}, u.InfoID)
	}
	if u.ConfigID > 0 {
		db.Unscoped().Delete(&UserConfig{}, u.ConfigID)
	}
	return nil
}

func (item *User) ToResponse() *data.UserResponse {
	if item == nil {
		return &data.UserResponse{}
	}
	resp := &data.UserResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Role = item.Role.ToResponse()
	resp.Info = item.Info.ToResponse()
	resp.Config = item.Config.ToResponse()

	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber
	resp.Status = item.Status
	resp.LoginMethod = item.LoginMethod
	resp.Provider = item.Provider
	resp.ProviderUserID = item.ProviderUserID
	resp.IsActivated = item.IsActivated
	resp.ActivatedAt = item.ActivatedAt
	return resp
}

func (item *User) ToPublicResponse() *data.UserPublicResponse {
	if item == nil {
		return &data.UserPublicResponse{}
	}
	resp := &data.UserPublicResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Role = item.Role.ToResponse()
	resp.Info = item.Info.ToPublicResponse()

	resp.Email = item.Email
	resp.Status = item.Status
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
