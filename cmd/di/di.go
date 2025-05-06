package di

import (
	"api/cmd/api"
	"api/config"
	"api/services/communication"
	"api/services/contact"
	"api/services/history"
	"api/services/school/common/director"
	"api/services/school/common/exam"
	"api/services/school/common/parent"
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
	"api/services/school/university/tu"
	"api/services/user/auth"
	"api/services/user/permission"
	"api/services/user/profile"
	"api/services/user/role"
	"api/services/user/user"
)

// InjectDependencies Inject all dependencies
func InjectDependencies() {
	// Others
	var communicationRepo = communication.NewRepository(config.DB)
	var contactRepo = contact.NewRepository(config.DB)
	var historyRepo = history.NewRepository(config.DB)
	api.AllControllers.CommunicationController = communication.NewController(
		communication.NewService(
			communicationRepo,
		),
	)
	api.AllControllers.ContactControllerController = contact.NewController(
		contact.NewService(
			contactRepo,
		),
	)
	api.AllControllers.HistoryController = history.NewController(
		history.NewService(
			historyRepo,
		),
	)

	// User
	var userRepo = user.NewRepository(config.DB)
	var roleRepo = role.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)
	api.AllControllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			userRepo,
			roleRepo,
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
	api.AllControllers.UserController = user.NewController(
		user.NewService(
			userRepo,
		),
	)
	api.AllControllers.ProfileController = profile.NewController(
		profile.NewService(
			userRepo,
		),
	)

	// School
	var schoolRepo = school.NewRepository(config.DB)
	var yearRepo = year.NewRepository(config.DB)
	var examRepo = exam.NewRepository(config.DB)
	var directorRepo = director.NewRepository(config.DB)
	var teacherRepo = teacher.NewRepository(config.DB)
	var studentRepo = student.NewRepository(config.DB)
	var parentRepo = parent.NewRepository(config.DB)
	api.AllControllers.SchoolController = school.NewController(
		school.NewService(
			schoolRepo,
		),
	)
	api.AllControllers.YearController = year.NewController(
		year.NewService(
			yearRepo,
		),
	)
	api.AllControllers.ExamController = exam.NewController(
		exam.NewService(
			examRepo,
		),
	)
	api.AllControllers.DirectorController = director.NewController(
		director.NewService(
			directorRepo, userRepo,
		),
	)
	api.AllControllers.TeacherController = teacher.NewController(
		teacher.NewService(
			teacherRepo,
		),
	)
	api.AllControllers.StudentController = student.NewController(
		student.NewService(
			studentRepo,
		),
	)
	api.AllControllers.ParentController = parent.NewController(
		parent.NewService(
			parentRepo,
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
			schoolRepo,
		),
	)
	api.AllControllers.QuarterController = quarter.NewController(
		quarter.NewService(
			quarterRepo,
			schoolRepo,
		),
	)
	api.AllControllers.SectionController = section.NewController(
		section.NewService(
			sectionRepo,
			schoolRepo,
		),
	)
	api.AllControllers.SpecialtyController = specialty.NewController(
		specialty.NewService(
			specialtyRepo,
			schoolRepo,
		),
	)
	api.AllControllers.ClassController = class.NewController(
		class.NewService(
			classRepo,
			schoolRepo,
		),
	)
	api.AllControllers.SubjectController = subject.NewController(
		subject.NewService(
			subjectRepo,
			schoolRepo,
		),
	)

	// University
	var semesterRepo = semester.NewRepository(config.DB)
	var facultyRepo = faculty.NewRepository(config.DB)
	var departmentRepo = department.NewRepository(config.DB)
	var domainRepo = domain.NewRepository(config.DB)
	var levelRepo = level.NewRepository(config.DB)
	var tuRepo = tu.NewRepository(config.DB)
	api.AllControllers.SemesterController = semester.NewController(
		semester.NewService(
			semesterRepo,
			schoolRepo,
		),
	)
	api.AllControllers.FacultyController = faculty.NewController(
		faculty.NewService(
			facultyRepo,
			schoolRepo,
		),
	)
	api.AllControllers.DepartmentController = department.NewController(
		department.NewService(
			departmentRepo,
			schoolRepo,
		),
	)
	api.AllControllers.DomainController = domain.NewController(
		domain.NewService(
			domainRepo,
			schoolRepo,
		),
	)
	api.AllControllers.LevelController = level.NewController(
		level.NewService(
			levelRepo,
			schoolRepo,
		),
	)
	api.AllControllers.TUController = tu.NewController(
		tu.NewService(
			tuRepo,
		),
	)
}
