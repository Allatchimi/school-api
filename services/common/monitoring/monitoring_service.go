package monitoring

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/common/monitoring/data"
	"api/services/school/common/director"
	"api/services/school/common/parent"
	"api/services/school/common/report"
	"api/services/school/common/school"
	"api/services/school/common/student"
	"api/services/school/common/teacher"
	"api/services/user/user"
)

type Service struct {
	UserService     *user.Service
	SchoolService   *school.Service
	DirectorService *director.Service
	TeacherService  *teacher.Service
	StudentService  *student.Service
	ParentService   *parent.Service
	ReportService   *report.Service
}

func NewService(
	userService *user.Service,
	schoolService *school.Service,
	directorService *director.Service,
	teacherService *teacher.Service,
	studentService *student.Service,
	parentService *parent.Service,
	reportService *report.Service,
) *Service {
	return &Service{
		UserService:     userService,
		SchoolService:   schoolService,
		DirectorService: directorService,
		TeacherService:  teacherService,
		StudentService:  studentService,
		ParentService:   parentService,
		ReportService:   reportService,
	}
}

const MODEL_NAME = "monitoring"
const DEFAULT_ERROR_MESSAGE = "interact with monitoring model"

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result *data.MonitoringResponseList, errCode int, err error) {
	// Count directors
	directorCount, err := service.DirectorService.Repository.CountAll(request.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count teachers
	teacherCount, err := service.TeacherService.Repository.CountAll(request.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count students
	studentCount, err := service.StudentService.Repository.CountAll(request.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count parents
	parentCount, err := service.ParentService.Repository.CountAll(request.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count users by year and schools
	var usersByYear = make([]data.UsersByYearResponse, 0)
	var schoolsCount int64
	if request.SchoolID < 1 {
		// Count users by year
		err = service.UserService.Repository.CountAllGroupByYear(&usersByYear)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		// Count schools
		schoolsCount, err = service.SchoolService.Repository.CountAll()
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}
	// Count success by school
	var successBySchool = make([]data.SuccessBySchoolResponse, 0)
	if request.SchoolID < 1 {
		// Count success by school
		// err = service.UserService.Repository.CountAllGroupBySchool(&successBySchool)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}
	// Count success by school gender
	var successBySchoolGender = make([]data.SuccessBySchoolGenderResponse, 0)
	if request.SchoolID < 1 {
		// Count success by school gender
		// err = service.UserService.Repository.CountAllGroupBySchoolGender(&successBySchoolGender)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	result = &data.MonitoringResponseList{
		Data: &data.MonitoringResponse{
			Count: &data.CountResponse{
				Schools:   schoolsCount,
				Directors: directorCount,
				Teachers:  teacherCount,
				Students:  studentCount,
				Parents:   parentCount,
			},
			UsersByYear:           usersByYear,
			SuccessBySchool:       successBySchool,
			SuccessBySchoolGender: successBySchoolGender,
		},
	}
	return
}
