package common_svc_permission

import (
	"api/common/constants"
	"api/services/school/common/director"
	"api/services/school/common/parent"
	"api/services/school/common/parent/model"
	"api/services/school/common/school"
	"api/services/school/common/student"
	"api/services/school/common/teacher"
	"api/services/user/user"
)

var userRepo *user.Repository
var schoolRepo *school.Repository
var directorRepo *director.Repository
var teacherRepo *teacher.Repository
var studentRepo *student.Repository
var parentRepo *parent.Repository

func InjectRepositories(
	userRepository *user.Repository,
	schoolRepository *school.Repository,
	directorRepository *director.Repository,
	teacherRepository *teacher.Repository,
	studentRepository *student.Repository,
	parentRepository *parent.Repository,
) {
	userRepo = userRepository
	schoolRepo = schoolRepository
	directorRepo = directorRepository
	teacherRepo = teacherRepository
	studentRepo = studentRepository
	parentRepo = parentRepository
}

func CanAccessBySchool(
	userID int64,
	schoolID int64,
) bool {
	// Get role
	foundUser, err := userRepo.GetByID(userID)
	if err != nil || foundUser == nil || foundUser.ID < 1 {
		return false
	}

	// Always return true for admin
	if foundUser.Role.Feature == constants.FeatureAdmin {
		return true
	}
	// For directors, check if the user is a director of the provided school id
	if foundUser.Role.Feature == constants.FeatureDirector {
		foundItem, _ := directorRepo.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}
	// For teacher, check if the user is a teacher of the provided school id
	if foundUser.Role.Feature == constants.FeatureTeacher {
		foundItem, _ := teacherRepo.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}
	// For student, check if the user is a student of the provided school id
	if foundUser.Role.Feature == constants.FeatureStudent {
		foundItem, _ := studentRepo.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}
	// For parent, check if the user is a parent of any student of the provided school id
	if foundUser.Role.Feature == constants.FeatureStudent {
		foundItem, _ := parentRepo.GetParentStudentByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 && foundItem.Parent.UserID == userID {
			if foundItem.Student.SchoolID == schoolID {
				return true
			}
		}
	}

	return false
}

func CanAccessBySchoolYearClassSubjectUnit(
	userID int64,
	schoolID int64,
	yearID int64,
	classSubjectID int64,
	unitID int64,
) bool {
	// Get role
	foundUser, err := userRepo.GetByID(userID)
	if err != nil || foundUser == nil || foundUser.Role.ID <= 0 {
		return false
	}

	// Always return true for admin
	if foundUser.Role.Feature == constants.FeatureAdmin {
		return true
	}
	// For directors, check if the user is a director of the provided school id
	if foundUser.Role.Feature == constants.FeatureDirector {
		foundItem, _ := directorRepo.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}

	// Retrieve school information
	foundSchool, err := schoolRepo.GetByID(schoolID)
	if err != nil || foundSchool == nil || foundSchool.ID <= 0 {
		return false
	}

	// For teacher, check if the user is a teacher of the provided school id, year id and class subject/unit id
	if foundUser.Role.Feature == constants.FeatureTeacher {
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			foundItem, _ := teacherRepo.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDClassSubjectID(
				userID, schoolID, yearID, classSubjectID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := teacherRepo.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDUnitID(
				userID, schoolID, yearID, unitID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		}
	}
	// For student, check if the user is a student of the provided school id year id and class subject/unit id
	if foundUser.Role.Feature == constants.FeatureStudent {
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			foundItem, _ := studentRepo.GetStudentEnrollByUserIDSchoolIDYearIDClassSubjectID(
				userID, schoolID, yearID, classSubjectID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := studentRepo.GetStudentEnrollByUserIDSchoolIDYearIDUnitID(
				userID, schoolID, yearID, unitID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		}
	}
	// For parent, check if the user is a parent of any student of the provided school id year id and class subject/unit id
	if foundUser.Role.Feature == constants.FeatureStudent {
		var foundItem *model.ParentStudent
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
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
	return false
}

func CanAccessBySchoolYearClassLevelDomain(
	userID int64,
	schoolID int64,
	yearID int64,
	classID int64,
	levelDomainID int64,
) bool {
	// Get role
	foundUser, err := userRepo.GetByID(userID)
	if err != nil || foundUser == nil || foundUser.Role.ID <= 0 {
		return false
	}

	// Always return true for admin
	if foundUser.Role.Feature == constants.FeatureAdmin {
		return true
	}
	// For directors, check if the user is a director of the provided school id
	if foundUser.Role.Feature == constants.FeatureDirector {
		foundItem, _ := directorRepo.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}

	// Retrieve school information
	foundSchool, err := schoolRepo.GetByID(schoolID)
	if err != nil || foundSchool == nil || foundSchool.ID <= 0 {
		return false
	}

	// For teacher, check if the user is a teacher of the provided school id year id and clas/level domain
	if foundUser.Role.Feature == constants.FeatureTeacher {
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			foundItem, _ := teacherRepo.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDClassID(
				userID, schoolID, yearID, classID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := teacherRepo.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDLevelDomainID(
				userID, schoolID, yearID, levelDomainID)
			if foundItem != nil && foundItem.ID > 0 && foundItem.Teacher.UserID == userID && foundItem.Teacher.SchoolID == schoolID {
				if foundItem.YearID == yearID && foundItem.Unit.LevelDomainID == levelDomainID {
					return true
				}
			}
		}
	}
	// For student, check if the user is a student of the provided school id year id and clas/level domain
	if foundUser.Role.Feature == constants.FeatureStudent {
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			foundItem, _ := studentRepo.GetStudentEnrollByUserIDSchoolIDYearIDClassID(
				userID, schoolID, yearID, classID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := studentRepo.GetStudentEnrollByUserIDSchoolIDYearIDLevelDomainID(
				userID, schoolID, yearID, levelDomainID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		}
	}
	// For parent, check if the user is a parent of any student of the provided school id year id and clas/level domain
	if foundUser.Role.Feature == constants.FeatureStudent {
		foundItem, _ := parentRepo.GetParentStudentByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 && foundItem.Parent.UserID == userID {
			if foundItem.Student.SchoolID == schoolID {
				return true
			}
		}
	}
	return false
}
