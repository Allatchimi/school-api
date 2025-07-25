package fixture

import (
	"api/common/constants"
	"api/config"
	"api/services/user/permission"
	permissionModel "api/services/user/permission/model"
	"api/services/user/role"
	roleModel "api/services/user/role/model"
	"api/services/user/user"
	userModel "api/services/user/user/model"
	"time"
)

// Load loads initial database values
func Load() (err error) {
	var roleRepo = role.NewRepository(config.DB)
	var userRepo = user.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)

	// Add role admin
	roleAdmin, err := roleRepo.GetByName(config.Env.FixtureRoleAdmin)
	if err != nil {
		return
	}
	if !(roleAdmin != nil && roleAdmin.Name == config.Env.FixtureRoleAdmin) {
		roleAdmin, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleAdmin,
			Feature:     constants.FeatureAdmin,
			Description: "Administrator role",
		})
	}
	// Add role default
	roleDefault, err := roleRepo.GetByName(config.Env.FixtureRoleDefault)
	if err != nil {
		return
	}
	if !(roleDefault != nil && roleDefault.Name == config.Env.FixtureRoleDefault) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleDefault,
			Feature:     constants.FeatureDefault,
			Description: "Default role",
		})
	}
	// Add role director
	roleDirector, err := roleRepo.GetByName(config.Env.FixtureRoleDirector)
	if err != nil {
		return
	}
	if !(roleDirector != nil && roleDirector.Name == config.Env.FixtureRoleDirector) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleDirector,
			Feature:     constants.FeatureDirector,
			Description: "Director role",
		})
	}
	// Add role teacher
	roleTeacher, err := roleRepo.GetByName(config.Env.FixtureRoleTeacher)
	if err != nil {
		return
	}
	if !(roleTeacher != nil && roleTeacher.Name == config.Env.FixtureRoleTeacher) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleTeacher,
			Feature:     constants.FeatureTeacher,
			Description: "Teacher role",
		})
	}
	// Add role student
	roleStudent, err := roleRepo.GetByName(config.Env.FixtureRoleStudent)
	if err != nil {
		return
	}
	if !(roleStudent != nil && roleStudent.Name == config.Env.FixtureRoleStudent) {
		_, err = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleStudent,
			Feature:     constants.FeatureStudent,
			Description: "Student role",
		})
		if err != nil {
			return
		}
	}
	// Add role parent
	roleParent, err := roleRepo.GetByName(config.Env.FixtureRoleParent)
	if err != nil {
		return
	}
	if !(roleParent != nil && roleParent.Name == config.Env.FixtureRoleParent) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleParent,
			Feature:     constants.FeatureParent,
			Description: "Parent role",
		})
	}

	// Add user admin
	userAdmin, err := userRepo.GetByEmailSchoolID(config.Env.FixtureUserAdminEmail, 0)
	if err != nil {
		return
	}
	if !(userAdmin != nil && userAdmin.Email == config.Env.FixtureUserAdminEmail) {
		userInfoAdmin, _ := userRepo.CreateUserInfo(&userModel.UserInfo{
			Username: "Admin",
			Language: "en",
		})
		userConfigAdmin, _ := userRepo.CreateUserConfig(&userModel.UserConfig{})

		tmpActivatedAt := time.Now()
		userAdmin, _ = userRepo.Create(&userModel.User{
			Email:    config.Env.FixtureUserAdminEmail,
			Password: config.Env.FixtureUserAdminPassword,

			Status: constants.USER_STATUS_ENABLED,

			LoginMethod: constants.AuthLoginMethodDefault,
			IsActivated: true,
			ActivatedAt: &tmpActivatedAt,

			RoleID:   roleAdmin.ID,
			InfoID:   userInfoAdmin.ID,
			ConfigID: userConfigAdmin.ID,
		})
	}

	if userAdmin == nil || roleAdmin == nil {
		return
	}

	// Add permissions for admin
	foundPermission, err := permissionRepo.GetByRoleIDTableName(roleAdmin.ID, "*")
	if err != nil {
		return
	}
	if foundPermission == nil || foundPermission.RoleID != roleAdmin.ID {
		_, err = permissionRepo.Create(&permissionModel.Permission{
			RoleID:    roleAdmin.ID,
			TableName: "*",
			Create:    true,
			Read:      true,
			Update:    true,
			Delete:    true,
		})
	}

	return
}
