package common_svc_permission

import (
	"api/common/constants"
	"api/services/school/common/director"
	"api/services/school/common/parent"
	"api/services/school/common/parent/model"
	"api/services/school/common/school"
	"api/services/school/common/student"
	"api/services/school/common/teacher"
	"slices"
)

var schoolRepo *school.Repository
var directorRepo *director.Repository
var teacherRepo *teacher.Repository
var studentRepo *student.Repository
var parentRepo *parent.Repository

func InjectRepositories(
	schoolRepository *school.Repository,
	directorRepository *director.Repository,
	teacherRepository *teacher.Repository,
	studentRepository *student.Repository,
	parentRepository *parent.Repository,
) {
	schoolRepo = schoolRepository
	directorRepo = directorRepository
	teacherRepo = teacherRepository
	studentRepo = studentRepository
	parentRepo = parentRepository
}

func CanAccessBySchool(
	roleID int64,
	userID int64,
	schoolID int64,
) bool {
	if slices.Contains(acceptedFeatureList, userFeature) {
		// Always return true for admin
		if userFeature == constants.FeatureAdmin {
			return true
		}
		// For directors, check if the user is a director of the provided school id
		if userFeature == constants.FeatureDirector {
			foundItem, _ := directorRepo.GetByUserID(userID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}
		// For teacher, check if the user is a teacher of the provided school id
		if userFeature == constants.FeatureTeacher {
			foundItem, _ := teacherRepo.GetByUserID(userID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}
		// For student, check if the user is a teacher of the provided school id
		if userFeature == constants.FeatureStudent {
			foundItem, _ := studentRepo.GetByID(userID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}
		// For parent, check if the user is a parent of any student of the provided school id
		if userFeature == constants.FeatureStudent {
			foundItem, _ := parentRepo.GetParentStudentByUserIDSchoolID(userID, schoolID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.Parent.UserID == userID {
				if foundItem.Student.SchoolID == schoolID {
					return true
				}
			}
		}
	}
	return false
}

func CanAccessBySchoolYearUnitClassSubject(
	roleID int64,
	userID int64,
	schoolID int64,
	yearID int64,
	unitID int64,
	classSubjectID int64,
) bool {
	if slices.Contains(acceptedFeatureList, userFeature) {
		// Always return true for admin
		if userFeature == constants.FeatureAdmin {
			return true
		}
		// For directors, check if the user is a director of the provided school id
		if userFeature == constants.FeatureDirector {
			foundItem, _ := directorRepo.GetByUserID(userID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}

		// Retrieve school information
		foundSchool, err := schoolRepo.GetByID(schoolID)
		if err != nil || foundSchool == nil || foundSchool.ID <= 0 {
			return false
		}

		// For teacher, check if the user is a teacher of the provided school id, year id and level domain/class id
		if userFeature == constants.FeatureTeacher {
			if foundSchool.Type == constants.SCHOOL_TYPE_UNIVERSITY {
				foundItem, _ := teacherRepo.GetTeacherUnitSubjectByUserIDSchoolIDYearIDUnitID(
					userID, schoolID, yearID, unitClassSubjectID)
				if foundItem != nil && foundItem.ID > 0 && foundItem.Teacher.UserID == userID && foundItem.Teacher.SchoolID == schoolID {
					if foundItem.YearID == yearID && foundItem.UnitID == unitClassSubjectID {
						return true
					}
				}
			} else {
				foundItem, _ := teacherRepo.GetTeacherUnitSubjectByUserIDSchoolIDYearIDClassSubjectID(
					userID, schoolID, yearID, unitClassSubjectID)
				if foundItem != nil && foundItem.ID > 0 && foundItem.Teacher.UserID == userID && foundItem.Teacher.SchoolID == schoolID {
					if foundItem.YearID == yearID && foundItem.UnitID == unitClassSubjectID {
						return true
					}
				}
			}
		}
		// For student, check if the user is a teacher of the provided school id and year id
		if userFeature == constants.FeatureStudent {
			foundItem, _ := studentRepo.GetByID(userID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}
		// For parent, check if the user is a parent of any student of the provided school id and year id
		if userFeature == constants.FeatureStudent {
			var foundItem *model.ParentStudent
			if foundSchool.Type == constants.SCHOOL_TYPE_UNIVERSITY {
				foundItem, err = parentRepo.GetParentStudentByUserIDSchoolID(userID, schoolID)
			} else {
				foundItem, err = parentRepo.GetParentStudentByUserIDSchoolID(userID, schoolID)
			}
			if foundItem != nil && foundItem.ID > 0 && foundItem.Parent.UserID == userID {
				if foundItem.Student.SchoolID == schoolID {
					return true
				}
			}
		}
	}
	return false
}

func CanAccessBySchoolYearLevelDomainClass(
	roleID int64,
	userID int64,
	schoolID int64,
	yearID int64,
	domainID int64,
	classID int64,
) bool {
	if slices.Contains(acceptedFeatureList, userFeature) {
		// Always return true for admin
		if userFeature == constants.FeatureAdmin {
			return true
		}
		// For directors, check if the user is a director of the provided school id
		if userFeature == constants.FeatureDirector {
			foundItem, _ := directorRepo.GetByUserID(userID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}
		// For teacher, check if the user is a teacher of the provided school id and year id
		if userFeature == constants.FeatureTeacher {
			foundItem, _ := teacherRepo.GetTeacherUnitSubjectByUserIDSchoolIDYearID(userID, schoolID, yearID, unitClassSubjectID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}
		// For student, check if the user is a teacher of the provided school id and year id
		if userFeature == constants.FeatureStudent {
			foundItem, _ := studentRepo.GetByID(userID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.UserID == userID {
				if foundItem.SchoolID == schoolID {
					return true
				}
			}
		}
		// For parent, check if the user is a parent of any student of the provided school id and year id
		if userFeature == constants.FeatureStudent {
			foundItem, _ := parentRepo.GetParentStudentByUserIDSchoolID(userID, schoolID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.Parent.UserID == userID {
				if foundItem.Student.SchoolID == schoolID {
					return true
				}
			}
		}
	}
	return false
}
