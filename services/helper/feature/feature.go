package serviceHelperFeature

import (
	"api/common/constants"
	"api/common/types"
	"api/services/school/common/parent"
	"api/services/school/common/student"
	"api/services/school/common/teacher"
)

var TeacherService *teacher.Service
var StudentService *student.Service
var ParentService *parent.Service

func InjectServices(
	teacherService *teacher.Service,
	studentService *student.Service,
	parentService *parent.Service,
) {
	TeacherService = teacherService
	StudentService = studentService
	ParentService = parentService
}

func GetUserDataByFeatureName(ctxData *types.ContextData) (teacherID int64, studentID int64, parentID int64, ok bool, err error) {
	if ctxData == nil {
		ok = false
		return
	}

	// Director
	if ctxData.User.Feature == constants.FeatureDirector {
		ok = true
		return
	}

	// Teacher
	if ctxData.User.Feature == constants.FeatureTeacher {
		foundTeacher, errFound := TeacherService.Repository.GetByUserID(ctxData.Jwt.UserID)
		if errFound != nil {
			err = errFound
			return
		}
		if foundTeacher == nil || foundTeacher.ID < 1 || foundTeacher.SchoolID != ctxData.Jwt.SchoolID {
			return
		}
		teacherID = foundTeacher.ID
		ok = true
		return
	}

	// Student
	if ctxData.User.Feature == constants.FeatureStudent {
		foundStudent, errFound := StudentService.Repository.GetByUserID(ctxData.Jwt.UserID)
		if errFound != nil {
			err = errFound
			return
		}
		if foundStudent == nil || foundStudent.ID < 1 || foundStudent.SchoolID != ctxData.Jwt.SchoolID {
			return
		}
		teacherID = foundStudent.ID
		ok = true
		return
	}

	// Parent
	if ctxData.User.Feature == constants.FeatureParent {
		foundParent, errFound := ParentService.Repository.GetByUserID(ctxData.Jwt.UserID)
		if errFound != nil {
			err = errFound
			return
		}
		if foundParent == nil || foundParent.ID < 1 || foundParent.SchoolID != ctxData.Jwt.SchoolID {
			return
		}
		teacherID = foundParent.ID
		ok = true
		return
	}
	return
}
