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
	roleAdmin, _ := roleRepo.GetByName(config.Env.FixtureRoleAdmin)
	if !(roleAdmin != nil && roleAdmin.Name == config.Env.FixtureRoleAdmin) {
		roleAdmin, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleAdmin,
			Feature:     constants.FeatureAdmin,
			Description: "Administrator role",
		})
	}
	// Add role default
	roleDefault, _ := roleRepo.GetByName(config.Env.FixtureRoleDefault)
	if !(roleDefault != nil && roleDefault.Name == config.Env.FixtureRoleDefault) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleDefault,
			Feature:     constants.FeatureDefault,
			Description: "Default role",
		})
	}
	// Add role director
	roleDirector, _ := roleRepo.GetByName(config.Env.FixtureRoleDirector)
	if !(roleDirector != nil && roleDirector.Name == config.Env.FixtureRoleDirector) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleDirector,
			Feature:     constants.FeatureDirector,
			Description: "Director role",
		})
	}
	// Add role teacher
	roleTeacher, _ := roleRepo.GetByName(config.Env.FixtureRoleTeacher)
	if !(roleTeacher != nil && roleTeacher.Name == config.Env.FixtureRoleTeacher) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleTeacher,
			Feature:     constants.FeatureTeacher,
			Description: "Teacher role",
		})
	}
	// Add role student
	roleStudent, _ := roleRepo.GetByName(config.Env.FixtureRoleStudent)
	if !(roleStudent != nil && roleStudent.Name == config.Env.FixtureRoleStudent) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleStudent,
			Feature:     constants.FeatureStudent,
			Description: "Student role",
		})
	}
	// Add role parent
	roleParent, _ := roleRepo.GetByName(config.Env.FixtureRoleParent)
	if !(roleParent != nil && roleParent.Name == config.Env.FixtureRoleParent) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.FixtureRoleParent,
			Feature:     constants.FeatureParent,
			Description: "Parent role",
		})
	}

	// Add user admin
	userAdmin, _ := userRepo.GetByEmail(config.Env.FixtureUserAdminEmail)
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

			LoginMethod: constants.AuthLoginMethodDefault,
			IsActivated: true,
			ActivatedAt: &tmpActivatedAt,

			RoleID:       roleAdmin.ID,
			UserInfoID:   userInfoAdmin.ID,
			UserConfigID: userConfigAdmin.ID,
		})
	}

	if userAdmin == nil || roleAdmin == nil {
		return
	}

	// Add permissions for admin
	foundPermission, _ := permissionRepo.GetByRoleIDTableName(roleAdmin.ID, "*")
	if foundPermission == nil || foundPermission.RoleID != roleAdmin.ID {
		_, _ = permissionRepo.Create(&permissionModel.Permission{
			RoleID:    roleAdmin.ID,
			TableName: "*",
			Create:    true,
			Read:      true,
			Update:    true,
			Delete:    true,
		})
		return
	}

	return
}
