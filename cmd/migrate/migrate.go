package migrate

import (
	"api/common/helpers"
	"api/config"
	communicationModel "api/services/communication/model"
	contactModel "api/services/contact/model"
	directorModel "api/services/school/common/director/model"
	examModel "api/services/school/common/exam/model"
	meetingModel "api/services/school/common/meeting/model"
	parentModel "api/services/school/common/parent/model"
	quizModel "api/services/school/common/quiz/model"
	schoolModel "api/services/school/common/school/model"
	studentModel "api/services/school/common/student/model"
	teacherModel "api/services/school/common/teacher/model"
	yearModel "api/services/school/common/year/model"
	classModel "api/services/school/highschool/class/model"
	quarterModel "api/services/school/highschool/quarter/model"
	sectionModel "api/services/school/highschool/section/model"
	sequenceModel "api/services/school/highschool/sequence/model"
	specialtyModel "api/services/school/highschool/specialty/model"
	subjectModel "api/services/school/highschool/subject/model"
	departmentModel "api/services/school/university/department/model"
	domainModel "api/services/school/university/domain/model"
	facultyModel "api/services/school/university/faculty/model"
	levelModel "api/services/school/university/level/model"
	semesterModel "api/services/school/university/semester/model"
	tuModel "api/services/school/university/unit/model"
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
		&examModel.Exam{},
		// Teacher
		&teacherModel.Teacher{},
		&teacherModel.TeacherUnitSubject{},
		// Student
		&studentModel.Student{},
		&studentModel.StudentEnroll{},
		// Parent
		&parentModel.Parent{},
		&parentModel.ParentStudent{},
		&parentModel.ParentAssign{},
		&parentModel.ParentAssignStudent{},
		// Meeting
		&meetingModel.MeetingRoom{},
		// Quiz
		&quizModel.Quiz{},
		&quizModel.QuizQuestion{},
		&quizModel.QuizQuestionOption{},
		&quizModel.QuizAttempt{},
		&quizModel.QuizAttemptQuestion{},

		// Highschool
		&sequenceModel.HighschoolSequence{},
		&quarterModel.HighschoolQuarter{},
		&sectionModel.HighschoolSection{},
		&specialtyModel.HighschoolSpecialty{},
		&subjectModel.HighschoolSubject{},
		&classModel.HighschoolClass{},
		&classModel.HighschoolClassSubject{},

		// University
		&semesterModel.UniversitySemester{},
		&facultyModel.UniversityFaculty{},
		&departmentModel.UniversityDepartment{},
		&domainModel.UniversityDomain{},
		&levelModel.UniversityLevel{},
		&levelModel.UniversityLevelDomain{},
		&tuModel.UniversityUnit{},
	)
	helpers.LogMigrations(
		err,
	)

	return err
}
