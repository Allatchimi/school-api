package di

import (
	"api/cmd/api"
	"api/config"
	serviceHelper "api/services/helper"
	"api/services/others/communication"
	"api/services/others/contact"
	"api/services/others/health"
	"api/services/others/monitoring"
	"api/services/others/notification"
	"api/services/others/telegram"
	"api/services/school/common/course"
	"api/services/school/common/director"
	"api/services/school/common/exam"
	"api/services/school/common/manager"
	"api/services/school/common/meeting"
	"api/services/school/common/parent"
	"api/services/school/common/payment"
	"api/services/school/common/quiz"
	"api/services/school/common/report"
	"api/services/school/common/request"
	"api/services/school/common/result"
	"api/services/school/common/schedule"
	"api/services/school/common/school"
	"api/services/school/common/student"
	"api/services/school/common/teacher"
	"api/services/school/common/year"
	"api/services/school/highschool/class"
	"api/services/school/highschool/quarter"
	"api/services/school/highschool/section"
	"api/services/school/highschool/sequence"
	"api/services/school/highschool/specialty"
	"api/services/school/highschool/subject"
	"api/services/school/university/department"
	"api/services/school/university/domain"
	"api/services/school/university/faculty"
	"api/services/school/university/level"
	"api/services/school/university/semester"
	"api/services/school/university/unit"
	"api/services/user/auth"
	"api/services/user/permission"
	"api/services/user/profile"
	"api/services/user/role"
	"api/services/user/user"
)

// InjectDependencies Inject all dependencies
func InjectDependencies() {
	// Common
	var communicationRepo = communication.NewRepository(config.DB)
	var contactRepo = contact.NewRepository(config.DB)
	var notificationRepo = notification.NewRepository(config.DB)
	var healthRepo = health.NewRepository(config.DB)
	api.AllControllers.CommunicationController = communication.NewController(
		communication.NewService(
			communicationRepo,
		),
	)
	api.AllControllers.ContactController = contact.NewController(
		contact.NewService(
			contactRepo,
		),
	)
	api.AllControllers.NotificationController = notification.NewController(
		notification.NewService(
			notificationRepo,
		),
	)
	api.AllControllers.HealthController = health.NewController(
		health.NewService(
			healthRepo,
		),
	)

	// User
	var userRepo = user.NewRepository(config.DB)
	var roleRepo = role.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)
	api.AllControllers.UserController = user.NewController(
		user.NewService(
			userRepo,
		),
	)
	api.AllControllers.RoleController = role.NewController(
		role.NewService(
			roleRepo,
		),
	)
	api.AllControllers.PermissionController = permission.NewController(
		permission.NewService(
			permissionRepo,
		),
	)
	api.AllControllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			api.AllControllers.UserController.Service,
			api.AllControllers.RoleController.Service,
		),
	)
	api.AllControllers.ProfileController = profile.NewController(
		profile.NewService(
			api.AllControllers.UserController.Service,
		),
	)

	// School
	var schoolRepo = school.NewRepository(config.DB)
	var directorRepo = director.NewRepository(config.DB)
	var managerRepo = manager.NewRepository(config.DB)
	var teacherRepo = teacher.NewRepository(config.DB)
	var studentRepo = student.NewRepository(config.DB)
	var parentRepo = parent.NewRepository(config.DB)
	var yearRepo = year.NewRepository(config.DB)
	var courseRepo = course.NewRepository(config.DB)
	var examRepo = exam.NewRepository(config.DB)
	var meetingRepo = meeting.NewRepository(config.DB)
	var quizRepo = quiz.NewRepository(config.DB)
	var resultRepo = result.NewRepository(config.DB)
	var reportRepo = report.NewRepository(config.DB)
	var scheduleRepo = schedule.NewRepository(config.DB)
	var requestRepo = request.NewRepository(config.DB)
	var paymentRepo = payment.NewRepository(config.DB)
	api.AllControllers.SchoolController = school.NewController(
		school.NewService(
			schoolRepo,
		),
	)
	api.AllControllers.DirectorController = director.NewController(
		director.NewService(
			directorRepo,
			api.AllControllers.RoleController.Service,
			api.AllControllers.UserController.Service,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.ManagerController = manager.NewController(
		manager.NewService(
			managerRepo,
			api.AllControllers.RoleController.Service,
			api.AllControllers.UserController.Service,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.TeacherController = teacher.NewController(
		teacher.NewService(
			teacherRepo,
			api.AllControllers.RoleController.Service,
			api.AllControllers.UserController.Service,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.StudentController = student.NewController(
		student.NewService(
			studentRepo,
			api.AllControllers.RoleController.Service,
			api.AllControllers.UserController.Service,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.ParentController = parent.NewController(
		parent.NewService(
			parentRepo,
			api.AllControllers.RoleController.Service,
			api.AllControllers.UserController.Service,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.YearController = year.NewController(
		year.NewService(
			yearRepo,
		),
	)
	api.AllControllers.CourseController = course.NewController(
		course.NewService(
			courseRepo,
		),
	)
	api.AllControllers.ExamController = exam.NewController(
		exam.NewService(
			examRepo,
		),
	)
	api.AllControllers.MeetingController = meeting.NewController(
		meeting.NewService(
			meetingRepo,
			api.AllControllers.UserController.Service,
			api.AllControllers.TeacherController.Service,
		),
	)
	api.AllControllers.QuizController = quiz.NewController(
		quiz.NewService(
			quizRepo,
		),
	)
	api.AllControllers.RequestController = request.NewController(
		request.NewService(
			requestRepo,
			api.AllControllers.StudentController.Service,
		),
	)
	api.AllControllers.ResultController = result.NewController(
		result.NewService(
			resultRepo,
		),
	)
	api.AllControllers.ScheduleController = schedule.NewController(
		schedule.NewService(
			scheduleRepo,
		),
	)
	api.AllControllers.PaymentController = payment.NewController(
		payment.NewService(
			paymentRepo,
		),
	)
	api.AllControllers.ReportController = report.NewController(
		report.NewService(
			reportRepo,
			api.AllControllers.SchoolController.Service,
			api.AllControllers.ResultController.Service,
			api.AllControllers.ExamController.Service,
		),
	)
	api.AllControllers.MonitoringController = monitoring.NewController(
		monitoring.NewService(
			api.AllControllers.UserController.Service,
			api.AllControllers.SchoolController.Service,
			api.AllControllers.DirectorController.Service,
			api.AllControllers.TeacherController.Service,
			api.AllControllers.StudentController.Service,
			api.AllControllers.ParentController.Service,
			api.AllControllers.ReportController.Service,
		),
	)

	// Highschool
	var sequenceRepo = sequence.NewRepository(config.DB)
	var quarterRepo = quarter.NewRepository(config.DB)
	var sectionRepo = section.NewRepository(config.DB)
	var specialtyRepo = specialty.NewRepository(config.DB)
	var classRepo = class.NewRepository(config.DB)
	var subjectRepo = subject.NewRepository(config.DB)
	api.AllControllers.SequenceController = sequence.NewController(
		sequence.NewService(
			sequenceRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.QuarterController = quarter.NewController(
		quarter.NewService(
			quarterRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.SectionController = section.NewController(
		section.NewService(
			sectionRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.SpecialtyController = specialty.NewController(
		specialty.NewService(
			specialtyRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.ClassController = class.NewController(
		class.NewService(
			classRepo,
			api.AllControllers.SchoolController.Service,
			api.AllControllers.MeetingController.Service,
		),
	)
	api.AllControllers.SubjectController = subject.NewController(
		subject.NewService(
			subjectRepo,
			api.AllControllers.SchoolController.Service,
		),
	)

	// University
	var semesterRepo = semester.NewRepository(config.DB)
	var facultyRepo = faculty.NewRepository(config.DB)
	var departmentRepo = department.NewRepository(config.DB)
	var domainRepo = domain.NewRepository(config.DB)
	var levelRepo = level.NewRepository(config.DB)
	var unitRepo = unit.NewRepository(config.DB)
	api.AllControllers.SemesterController = semester.NewController(
		semester.NewService(
			semesterRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.FacultyController = faculty.NewController(
		faculty.NewService(
			facultyRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.DepartmentController = department.NewController(
		department.NewService(
			departmentRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.DomainController = domain.NewController(
		domain.NewService(
			domainRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.LevelController = level.NewController(
		level.NewService(
			levelRepo,
			api.AllControllers.SchoolController.Service,
		),
	)
	api.AllControllers.UnitController = unit.NewController(
		unit.NewService(
			unitRepo,
			api.AllControllers.SchoolController.Service,
			api.AllControllers.MeetingController.Service,
		),
	)

	// Telegram
	api.AllControllers.TelegramController = telegram.NewController(
		telegram.NewService(
			api.AllControllers.SchoolController.Service,
		),
	)

	// Permissions checker
	serviceHelper.InjectServices(
		api.AllControllers.UserController.Service,
		api.AllControllers.NotificationController.Service,
	)
}
