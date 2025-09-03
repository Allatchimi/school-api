package migrations

import (
	"api/common/constants"
	"api/config"
	"api/services/user/role/model"
)

// M1 is the first migration
func M1() error {
	updateRoleFeature("dashboard-admin", constants.FeatureAdmin)
	updateRoleFeature("dashboard-director", constants.FeatureDirector)
	updateRoleFeature("dashboard-teacher", constants.FeatureTeacher)
	updateRoleFeature("dashboard-student", constants.FeatureStudent)
	updateRoleFeature("dashboard-parent", constants.FeatureParent)
	updateRoleFeature("dashboard-default", constants.FeatureDefault)
	return nil
}

func updateRoleFeature(oldFeature string, newFeature string) (err error) {
	fields := map[string]any{
		"feature": newFeature,
	}
	err = config.DB.
		Model(&model.Role{}).
		Where("feature = ?", oldFeature).
		Updates(fields).Error
	return
}
