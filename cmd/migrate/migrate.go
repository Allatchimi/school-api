package migrate

import (
	"api/common/helpers"
	"api/config"
	communicationModel "api/services/communication/model"
	contactModel "api/services/contact/model"
	historyModel "api/services/history/model"
	directorModel "api/services/school/common/director/model"
	examModel "api/services/school/common/exam/model"
	parentModel "api/services/school/common/parent/model"
	schoolModel "api/services/school/common/school/model"
	studentModel "api/services/school/common/student/model"
	teacherModel "api/services/school/common/teacher/model"
	yearModel "api/services/school/common/year/model"
	classModel "api/services/school/highschool/class/model"
	sectionModel "api/services/school/highschool/section/model"
	specialtyModel "api/services/school/highschool/specialty/model"
	subjectModel "api/services/school/highschool/subject/model"
	departmentModel "api/services/school/university/department/model"
	domainModel "api/services/school/university/domain/model"
	facultyModel "api/services/school/university/faculty/model"
	levelModel "api/services/school/university/level/model"
	tuModel "api/services/school/university/tu/model"
	permissionModel "api/services/user/permission/model"
	roleModel "api/services/user/role/model"
	userModel "api/services/user/user/model"
)

// Start Loads and applies all migrations.
func Start() error {
	err := config.DB.AutoMigrate(
		// Others
		&communicationModel.Communication{},
		&contactModel.Contact{},
		&historyModel.History{},

		// User
		&permissionModel.Permission{},
		&userModel.User{},
		&roleModel.Role{},
		&userModel.UserMfa{},
		&userModel.UserInfo{},

		// Director
		&directorModel.Director{},

		// School
		&schoolModel.School{},
		&schoolModel.SchoolInfo{},
		&schoolModel.SchoolConfig{},
		&yearModel.Year{},
		// Teacher
		&teacherModel.Teacher{},
		&teacherModel.TeacherLevelClass{},
		// Student
		&studentModel.Student{},
		&studentModel.StudentLevelClass{},
		// Parent
		&parentModel.Parent{},
		&parentModel.ParentStudent{},

		&examModel.Exam{},

		// Highschool
		&sectionModel.HighschoolSection{},
		&specialtyModel.HighschoolSpecialty{},
		&classModel.HighschoolClass{},
		&subjectModel.Subject{},
		&subjectModel.SubjectProfessor{},

		// University
		&facultyModel.UniversityFaculty{},
		&departmentModel.UniversityDepartment{},
		&domainModel.UniversityDomain{},
		&levelModel.UniversityLevel{},
		&tuModel.TeachingUnit{},
		&tuModel.TeachingUnitProfessor{},
	)
	helpers.LogMigrations(
		err,
	)

	return err
}
