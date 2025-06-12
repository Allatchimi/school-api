package migrate

import (
	"api/common/helpers"
	"api/config"
	communicationModel "api/services/common/communication/model"
	contactModel "api/services/common/contact/model"
	notificationModel "api/services/common/notification/model"
	courseModel "api/services/school/common/course/model"
	directorModel "api/services/school/common/director/model"
	examModel "api/services/school/common/exam/model"
	meetingModel "api/services/school/common/meeting/model"
	parentModel "api/services/school/common/parent/model"
	quizModel "api/services/school/common/quiz/model"
	requestModel "api/services/school/common/request/model"
	resultModel "api/services/school/common/result/model"
	scheduleModel "api/services/school/common/schedule/model"
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

// Apply applies all models with default migration rule.
func Apply() error {
	err := config.DB.AutoMigrate(
		// ----------- Others models -----------
		// Notification
		&notificationModel.Notification{},
		// Communication
		&communicationModel.Communication{},
		// Contact
		&contactModel.Contact{},

		// ----------- User models -----------
		// Permission
		&permissionModel.Permission{},
		// User
		&userModel.User{},
		&userModel.UserMfa{},
		&userModel.UserInfo{},
		// Role
		&roleModel.Role{},

		// ----------- Common school models -----------
		// Director
		&directorModel.Director{},
		// School
		&schoolModel.School{},
		&schoolModel.SchoolInfo{},
		&schoolModel.SchoolConfig{},
		// Year
		&yearModel.Year{},
		// Teacher
		&teacherModel.Teacher{},
		&teacherModel.TeacherClassSubjectUnit{},
		// Student
		&studentModel.Student{},
		&studentModel.StudentEnroll{},
		// Parent
		&parentModel.Parent{},
		&parentModel.ParentStudent{},
		&parentModel.ParentAssign{},
		&parentModel.ParentAssignStudent{},
		// Course
		&courseModel.Course{},
		&courseModel.CourseDocument{},
		&courseModel.CourseVideo{},
		&courseModel.CourseComment{},
		// Exam
		&examModel.Exam{},
		&examModel.ExamType{},
		// Meeting
		&meetingModel.MeetingRoom{},
		// Quiz
		&quizModel.Quiz{},
		&quizModel.QuizQuestion{},
		&quizModel.QuizQuestionOption{},
		&quizModel.QuizAnswer{},
		// Request
		&requestModel.Request{},
		// Schedule
		&scheduleModel.Schedule{},
		// Result
		&resultModel.Result{},

		// ----------- Highschool models -----------
		// Section
		&sectionModel.HighschoolSection{},
		// Specialty
		&specialtyModel.HighschoolSpecialty{},
		// Class
		&classModel.HighschoolClass{},
		&classModel.HighschoolClassSubject{},
		// Quarter
		&quarterModel.HighschoolQuarter{},
		&quarterModel.HighschoolQuarterSequence{},
		// Sequence
		&sequenceModel.HighschoolSequence{},
		// Subject
		&subjectModel.HighschoolSubject{},

		// ----------- University models -----------
		// Faculty
		&facultyModel.UniversityFaculty{},
		// Department
		&departmentModel.UniversityDepartment{},
		// Domain
		&domainModel.UniversityDomain{},
		// Level
		&levelModel.UniversityLevel{},
		&levelModel.UniversityLevelDomain{},
		// Semester
		&semesterModel.UniversitySemester{},
		// Unit
		&tuModel.UniversityUnit{},
	)
	helpers.LogMigrations(
		err,
	)

	return err
}
