package permissionchecker

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

var userSvc *user.Service
var schoolSvc *school.Service
var directorSvc *director.Service
var teacherSvc *teacher.Service
var studentSvc *student.Service
var parentSvc *parent.Service

func InjectServices(
	userService *user.Service,
	schoolService *school.Service,
	directorService *director.Service,
	teacherService *teacher.Service,
	studentService *student.Service,
	parentService *parent.Service,
) {
	userSvc = userService
	schoolSvc = schoolService
	directorSvc = directorService
	teacherSvc = teacherService
	studentSvc = studentService
	parentSvc = parentService
}

func CanAccessBySchool(
	userID int64,
	schoolID int64,
) bool {
	// Get role
	foundUser, err := userSvc.Repository.GetByID(userID)
	if err != nil || foundUser == nil || foundUser.ID < 1 {
		return false
	}

	// Always return true for admin
	if foundUser.Role.Feature == constants.FeatureAdmin {
		return true
	}
	// For directors, check if the user is a director of the provided school id
	if foundUser.Role.Feature == constants.FeatureDirector {
		foundItem, _ := directorSvc.Repository.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}
	// For teacher, check if the user is a teacher of the provided school id
	if foundUser.Role.Feature == constants.FeatureTeacher {
		foundItem, _ := teacherSvc.Repository.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}
	// For student, check if the user is a student of the provided school id
	if foundUser.Role.Feature == constants.FeatureStudent {
		foundItem, _ := studentSvc.Repository.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}
	// For parent, check if the user is a parent of any student of the provided school id
	if foundUser.Role.Feature == constants.FeatureStudent {
		foundItem, _ := parentSvc.Repository.GetParentStudentByUserIDSchoolID(userID, schoolID)
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
	foundUser, err := userSvc.Repository.GetByID(userID)
	if err != nil || foundUser == nil || foundUser.Role.ID <= 0 {
		return false
	}

	// Always return true for admin
	if foundUser.Role.Feature == constants.FeatureAdmin {
		return true
	}
	// For directors, check if the user is a director of the provided school id
	if foundUser.Role.Feature == constants.FeatureDirector {
		foundItem, _ := directorSvc.Repository.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}

	// Retrieve school information
	foundSchool, err := schoolSvc.Repository.GetByID(schoolID)
	if err != nil || foundSchool == nil || foundSchool.ID <= 0 {
		return false
	}

	// For teacher, check if the user is a teacher of the provided school id, year id and class subject/unit id
	if foundUser.Role.Feature == constants.FeatureTeacher {
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			foundItem, _ := teacherSvc.Repository.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDClassSubjectID(
				userID, schoolID, yearID, classSubjectID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := teacherSvc.Repository.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDUnitID(
				userID, schoolID, yearID, unitID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		}
	}
	// For student, check if the user is a student of the provided school id year id and class subject/unit id
	if foundUser.Role.Feature == constants.FeatureStudent {
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			foundItem, _ := studentSvc.Repository.GetStudentEnrollByUserIDSchoolIDYearIDClassSubjectID(
				userID, schoolID, yearID, classSubjectID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := studentSvc.Repository.GetStudentEnrollByUserIDSchoolIDYearIDUnitID(
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
			foundItem, err = parentSvc.Repository.GetParentStudentByUserIDSchoolID(userID, schoolID)
		} else {
			foundItem, err = parentSvc.Repository.GetParentStudentByUserIDSchoolID(userID, schoolID)
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
	foundUser, err := userSvc.Repository.GetByID(userID)
	if err != nil || foundUser == nil || foundUser.Role.ID <= 0 {
		return false
	}

	// Always return true for admin
	if foundUser.Role.Feature == constants.FeatureAdmin {
		return true
	}
	// For directors, check if the user is a director of the provided school id
	if foundUser.Role.Feature == constants.FeatureDirector {
		foundItem, _ := directorSvc.Repository.GetByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 {
			return true
		}
	}

	// Retrieve school information
	foundSchool, err := schoolSvc.Repository.GetByID(schoolID)
	if err != nil || foundSchool == nil || foundSchool.ID <= 0 {
		return false
	}

	// For teacher, check if the user is a teacher of the provided school id year id and clas/level domain
	if foundUser.Role.Feature == constants.FeatureTeacher {
		if foundSchool.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			foundItem, _ := teacherSvc.Repository.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDClassID(
				userID, schoolID, yearID, classID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := teacherSvc.Repository.GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDLevelDomainID(
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
			foundItem, _ := studentSvc.Repository.GetStudentEnrollByUserIDSchoolIDYearIDClassID(
				userID, schoolID, yearID, classID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		} else {
			foundItem, _ := studentSvc.Repository.GetStudentEnrollByUserIDSchoolIDYearIDLevelDomainID(
				userID, schoolID, yearID, levelDomainID)
			if foundItem != nil && foundItem.ID > 0 {
				return true
			}
		}
	}
	// For parent, check if the user is a parent of any student of the provided school id year id and clas/level domain
	if foundUser.Role.Feature == constants.FeatureStudent {
		foundItem, _ := parentSvc.Repository.GetParentStudentByUserIDSchoolID(userID, schoolID)
		if foundItem != nil && foundItem.ID > 0 && foundItem.Parent.UserID == userID {
			if foundItem.Student.SchoolID == schoolID {
				return true
			}
		}
	}
	return false
}
