package serviceHelperUser

import (
	"api/services/school/common/parent"
	dataParent "api/services/school/common/parent/data"
	"api/services/school/common/student"
	dataStudent "api/services/school/common/student/data"
	"api/services/school/common/teacher"
	dataTeacher "api/services/school/common/teacher/data"
	"api/services/user/user"
	dataUser "api/services/user/user/data"
	"api/services/user/user/model"
)

var UserService *user.Service
var TeacherService *teacher.Service
var StudentService *student.Service
var ParentService *parent.Service

func InjectServices(
	userService *user.Service,
	teacherService *teacher.Service,
	studentService *student.Service,
	parentService *parent.Service,
) {
	UserService = userService
	TeacherService = teacherService
	StudentService = studentService
	ParentService = parentService
}

func GetAllUserForTeacherClassSubjectUnit(request *dataTeacher.GetAllTeacherClassSubjectUnitRequest) (users []model.User, err error) {
	foundItems, err := TeacherService.Repository.GetAllTeacherClassSubjectUnit(nil, nil, request)
	if err != nil {
		return
	}
	users = make([]model.User, 0, len(foundItems))
	if len(foundItems) < 1 {
		return
	}
	for _, item := range foundItems {
		if item.Teacher != nil && item.Teacher.User != nil && item.Teacher.User.ID > 0 {
			users = append(users, *item.Teacher.User)
		}
	}
	return
}

func GetAllUserForStudentEnroll(request *dataStudent.GetAllStudentEnrollRequest) (users []model.User, err error) {
	foundItems, err := StudentService.Repository.GetAllStudentEnroll(nil, nil, request)
	if err != nil {
		return
	}
	users = make([]model.User, 0, len(foundItems))
	if len(foundItems) < 1 {
		return
	}
	for _, item := range foundItems {
		if item.Student != nil && item.Student.User != nil && item.Student.User.ID > 0 {
			users = append(users, *item.Student.User)
		}
	}
	return
}

func GetAllUserForParentStudent(request *dataParent.GetAllParentStudentRequest) (users []model.User, err error) {
	foundItems, err := ParentService.Repository.GetAllParentStudent(nil, nil, request)
	if err != nil {
		return
	}
	users = make([]model.User, 0, len(foundItems))
	if len(foundItems) < 1 {
		return
	}
	for _, item := range foundItems {
		if item.Parent != nil && item.Parent.User != nil && item.Parent.User.ID > 0 {
			users = append(users, *item.Parent.User)
		}
	}
	return
}

func GetAllUserByFeature(feature string) (users []model.User, err error) {
	users, err = UserService.Repository.GetAll(nil, nil, &dataUser.GetAllRequest{Feature: feature})
	return
}
