package api

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"

	"api/common/constants"
	"api/config"
	"api/middlewares"
	"api/services/common/communication"
	"api/services/common/contact"
	"api/services/common/health"
	"api/services/school/common/course"
	"api/services/school/common/director"
	"api/services/school/common/exam"
	"api/services/school/common/meeting"
	"api/services/school/common/parent"
	"api/services/school/common/payment"
	"api/services/school/common/quiz"
	"api/services/school/common/request"
	"api/services/school/common/result"
	"api/services/school/common/schedule"
	"api/services/school/common/school"
	"api/services/school/common/statistic"
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
	// Others service
	CommunicationController *communication.Controller
	ContactController       *contact.Controller
	HealthController        *health.Controller

	// User service
	AuthController       *auth.Controller
	RoleController       *role.Controller
	PermissionController *permission.Controller
	UserController       *user.Controller
	ProfileController    *profile.Controller

	// School service
	SchoolController    *school.Controller
	DirectorController  *director.Controller
	TeacherController   *teacher.Controller
	StudentController   *student.Controller
	ParentController    *parent.Controller
	YearController      *year.Controller
	CourseController    *course.Controller
	ExamController      *exam.Controller
	MeetingController   *meeting.Controller
	QuizController      *quiz.Controller
	RequestController   *request.Controller
	ResultController    *result.Controller
	ScheduleController  *schedule.Controller
	PaymentController   *payment.Controller
	StatisticController *statistic.Controller
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
}

var AllControllers = &Controllers{}

// Register all API endpoints
func registerEndpoints(humaApi *huma.API) {
	// Others service
	communication.RegisterEndpoints(humaApi, AllControllers.CommunicationController)
	contact.RegisterEndpoints(humaApi, AllControllers.ContactController)
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
	teacher.RegisterEndpoints(humaApi, AllControllers.TeacherController)
	student.RegisterEndpoints(humaApi, AllControllers.StudentController)
	parent.RegisterEndpoints(humaApi, AllControllers.ParentController)
	year.RegisterEndpoints(humaApi, AllControllers.YearController)
	course.RegisterEndpoints(humaApi, AllControllers.CourseController)
	exam.RegisterEndpoints(humaApi, AllControllers.ExamController)
	meeting.RegisterEndpoints(humaApi, AllControllers.MeetingController)
	quiz.RegisterEndpoints(humaApi, AllControllers.QuizController)
	request.RegisterEndpoints(humaApi, AllControllers.RequestController)
	result.RegisterEndpoints(humaApi, AllControllers.ResultController)
	schedule.RegisterEndpoints(humaApi, AllControllers.ScheduleController)
	payment.RegisterEndpoints(humaApi, AllControllers.PaymentController)
	statistic.RegisterEndpoints(humaApi, AllControllers.StatisticController)
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
}

// Start Set up and start the API: set up API documentation,
// configure middlewares, and security measures.
func Start() {
	// Set up gin for your API
	gin.SetMode(config.Env.GinMode)
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	// Register gin middlewares
	engine.Use(middlewares.GinContextRegister())

	// Group API
	ginGroup := engine.Group(config.Env.ApiGroup)

	// OpenAPI documentation based on huma
	humaConfig := huma.DefaultConfig(constants.OpenApiTitle, constants.OpenApiVersion)
	// Custom hook to remove schema links
	humaConfig.CreateHooks = []func(huma.Config) huma.Config{
		func(c huma.Config) huma.Config {
			return c
		},
	}
	humaConfig.DocsPath = ""
	humaConfig.Servers = []*huma.Server{
		{URL: config.Env.ApiGroup},
	}
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		constants.SecurityAuthName: {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
			Description:  "Bearer token used to access some resources",
		},
	}
	humaConfig.Info.Description = constants.OpenApiDescription
	humaApi := humagin.NewWithGroup(engine, ginGroup, humaConfig)
	// Register middlewares
	humaApi.UseMiddleware(
		middlewares.HeadersMiddleware(humaApi),
		middlewares.CorsMiddleware(humaApi),
		middlewares.AuthMiddleware(humaApi),
		middlewares.PermissionMiddleware(
			humaApi,
			AllControllers.UserController.Service.Repository,
			AllControllers.PermissionController.Service.Repository,
		),
	)

	// Register endpoints
	// Serve static files as favicon
	engine.StaticFS("/assets", http.Dir(constants.AssetAppPath))
	// Register endpoint for docs with support for custom template
	ginGroup.GET("/docs", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", []byte(*config.OpenAPITemplates.Scalar))
	})
	registerEndpoints(&humaApi)

	// Start to listen
	formattedPort := fmt.Sprintf(":%d", config.Env.AppPort)
	err = engine.Run(formattedPort)
	if err != nil {
		panic(err)
	}
}
