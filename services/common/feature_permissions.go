package service_permission

import (
	"api/common/constants"
	"api/services/school/common/director"
	"api/services/school/common/school"
	"api/services/school/common/student"
	"api/services/school/common/teacher"
	"slices"
)

func CanCreate(
	schoolRepo school.Repository,
	directorRepo director.Repository,
	teacherRepo teacher.Repository,
	studentRepo student.Repository,
	acceptedFeatureList []string,
	userFeature string,
	userID int64,
	schoolID int64,
) bool {
	if slices.Contains(acceptedFeatureList, userFeature) {
		if userFeature == constants.FeatureAdmin {
			return true
		}
		if userFeature == constants.FeatureDirector {
			// Check if the user is director of the provided school id
		}
		if userFeature == constants.FeatureTeacher {
			// Check if the user is teacher of the provided school id
		}
		if userFeature == constants.FeatureStudent {
			// Check if the user is student of the provided school id
		}
	}
	return false
}
