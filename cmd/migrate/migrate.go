package migrate

import (
	"api/common/helpers"
	"api/config"
	communicationModel "api/services/communication/model"
	contactModel "api/services/contact/model"
	historyModel "api/services/history/model"
	directorModel "api/services/school/common/director/model"
	examModel "api/services/school/common/exam/model"
	schoolModel "api/services/school/common/school/model"
	yearModel "api/services/school/common/year/model"
	classModel "api/services/school/highschool/class/model"
	pupilModel "api/services/school/highschool/pupil/model"
	sectionModel "api/services/school/highschool/section/model"
	specialtyModel "api/services/school/highschool/specialty/model"
	subjectModel "api/services/school/highschool/subject/model"
	testModel "api/services/school/highschool/test/model"
	departmentModel "api/services/school/university/department/model"
	domainModel "api/services/school/university/domain/model"
	facultyModel "api/services/school/university/faculty/model"
	levelModel "api/services/school/university/level/model"
	studentModel "api/services/school/university/student/model"
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

		// School
		&yearModel.Year{},
		&schoolModel.School{},
		&schoolModel.SchoolInfo{},
		&schoolModel.SchoolConfig{},
		&examModel.Exam{},

		// Director
		&directorModel.Director{},

		// Highschool
		&sectionModel.HighschoolSection{},
		&specialtyModel.HighschoolSpecialty{},
		&classModel.HighschoolClass{},
		&subjectModel.Subject{},
		&subjectModel.SubjectProfessor{},
		&pupilModel.Pupil{},
		&testModel.Test{},

		// University
		&facultyModel.UniversityFaculty{},
		&departmentModel.UniversityDepartment{},
		&domainModel.UniversityDomain{},
		&levelModel.UniversityLevel{},
		&tuModel.TeachingUnit{},
		&tuModel.TeachingUnitProfessor{},
		&studentModel.Student{},
	)
	helpers.LogMigrations(
		err,
	)

	return err
}
