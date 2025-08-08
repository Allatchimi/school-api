package api

import (
	"github.com/danielgtaylor/huma/v2"

	"api/services/others/communication"
	"api/services/others/contact"
	"api/services/others/health"
	"api/services/others/initialize"
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

type Controllers struct {
	// Common service
	CommunicationController *communication.Controller
	ContactController       *contact.Controller
	NotificationController  *notification.Controller
	HealthController        *health.Controller

	// User service
	AuthController       *auth.Controller
	RoleController       *role.Controller
	PermissionController *permission.Controller
	UserController       *user.Controller
	ProfileController    *profile.Controller

	// School service
	SchoolController     *school.Controller
	DirectorController   *director.Controller
	ManagerController    *manager.Controller
	TeacherController    *teacher.Controller
	StudentController    *student.Controller
	ParentController     *parent.Controller
	YearController       *year.Controller
	CourseController     *course.Controller
	ExamController       *exam.Controller
	MeetingController    *meeting.Controller
	QuizController       *quiz.Controller
	ResultController     *result.Controller
	ReportController     *report.Controller
	ScheduleController   *schedule.Controller
	RequestController    *request.Controller
	PaymentController    *payment.Controller
	MonitoringController *monitoring.Controller
	// Secondary
	SectionController   *section.Controller
	SpecialtyController *specialty.Controller
	ClassController     *class.Controller
	SubjectController   *subject.Controller
	SequenceController  *sequence.Controller
	QuarterController   *quarter.Controller
	// Faculty
	FacultyController    *faculty.Controller
	DepartmentController *department.Controller
	DomainController     *domain.Controller
	LevelController      *level.Controller
	SemesterController   *semester.Controller
	UnitController       *unit.Controller

	// Telegram
	TelegramController *telegram.Controller

	// Init
	InitController *initialize.Controller
}

var AllControllers = &Controllers{}

// Register all API endpoints
func registerEndpoints(humaApi *huma.API) {
	// Common service
	communication.RegisterEndpoints(humaApi, AllControllers.CommunicationController)
	contact.RegisterEndpoints(humaApi, AllControllers.ContactController)
	notification.RegisterEndpoints(humaApi, AllControllers.NotificationController)
	health.RegisterEndpoints(humaApi, AllControllers.HealthController)

	// User service
	auth.RegisterEndpoints(humaApi, AllControllers.AuthController)
	role.RegisterEndpoints(humaApi, AllControllers.RoleController)
	permission.RegisterEndpoints(humaApi, AllControllers.PermissionController)
	user.RegisterEndpoints(humaApi, AllControllers.UserController)
	profile.RegisterEndpoints(humaApi, AllControllers.ProfileController)

	// School service
	school.RegisterEndpoints(humaApi, AllControllers.SchoolController)
	director.RegisterEndpoints(humaApi, AllControllers.DirectorController)
	manager.RegisterEndpoints(humaApi, AllControllers.ManagerController)
	teacher.RegisterEndpoints(humaApi, AllControllers.TeacherController)
	student.RegisterEndpoints(humaApi, AllControllers.StudentController)
	parent.RegisterEndpoints(humaApi, AllControllers.ParentController)
	year.RegisterEndpoints(humaApi, AllControllers.YearController)
	course.RegisterEndpoints(humaApi, AllControllers.CourseController)
	exam.RegisterEndpoints(humaApi, AllControllers.ExamController)
	meeting.RegisterEndpoints(humaApi, AllControllers.MeetingController)
	quiz.RegisterEndpoints(humaApi, AllControllers.QuizController)
	report.RegisterEndpoints(humaApi, AllControllers.ReportController)
	result.RegisterEndpoints(humaApi, AllControllers.ResultController)
	schedule.RegisterEndpoints(humaApi, AllControllers.ScheduleController)
	request.RegisterEndpoints(humaApi, AllControllers.RequestController)
	payment.RegisterEndpoints(humaApi, AllControllers.PaymentController)
	monitoring.RegisterEndpoints(humaApi, AllControllers.MonitoringController)
	// Highschool
	sequence.RegisterEndpoints(humaApi, AllControllers.SequenceController)
	quarter.RegisterEndpoints(humaApi, AllControllers.QuarterController)
	section.RegisterEndpoints(humaApi, AllControllers.SectionController)
	specialty.RegisterEndpoints(humaApi, AllControllers.SpecialtyController)
	class.RegisterEndpoints(humaApi, AllControllers.ClassController)
	subject.RegisterEndpoints(humaApi, AllControllers.SubjectController)
	// University
	semester.RegisterEndpoints(humaApi, AllControllers.SemesterController)
	faculty.RegisterEndpoints(humaApi, AllControllers.FacultyController)
	department.RegisterEndpoints(humaApi, AllControllers.DepartmentController)
	domain.RegisterEndpoints(humaApi, AllControllers.DomainController)
	level.RegisterEndpoints(humaApi, AllControllers.LevelController)
	unit.RegisterEndpoints(humaApi, AllControllers.UnitController)

	// Telegram
	telegram.RegisterEndpoints(humaApi, AllControllers.TelegramController)

	// Init
	initialize.RegisterEndpoints(humaApi, AllControllers.InitController)
}
