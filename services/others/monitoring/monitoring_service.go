package monitoring

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/others/monitoring/data"
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
	// Count users
	directorCount, err := service.UserService.Repository.CountAllByFeature(request.SchoolID, constants.FeatureDirector)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	teacherCount, err := service.UserService.Repository.CountAllByFeature(request.SchoolID, constants.FeatureTeacher)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	studentCount, err := service.UserService.Repository.CountAllByFeature(request.SchoolID, constants.FeatureStudent)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	parentCount, err := service.UserService.Repository.CountAllByFeature(request.SchoolID, constants.FeatureParent)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count users by feature
	var usersByFeature = make([]data.UsersByFeatureResponse, 0)
	err = service.UserService.Repository.CountAllGroupByFeature(request.SchoolID, &usersByFeature)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count users by gender
	var usersByGender = make([]data.UsersByGenderResponse, 0)
	err = service.UserService.Repository.CountAllGroupByGender(request.SchoolID, &usersByGender)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count users by month
	var usersByMonth = make([]data.UsersByMonthResponse, 0)
	err = service.UserService.Repository.CountAllGroupByMonth(request.SchoolID, &usersByMonth)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count users by year
	var usersByYear = make([]data.UsersByYearResponse, 0)
	err = service.UserService.Repository.CountAllGroupByYear(request.SchoolID, &usersByYear)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count schools
	var schoolsCount int64
	schoolsCount, err = service.SchoolService.Repository.CountAll()
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Count success by school
	var successBySchool = make([]data.SuccessBySchoolResponse, 0)
	// err = service.UserService.Repository.CountAllGroupBySchool(&successBySchool)
	// if err != nil {
	// 	errCode = http.StatusInternalServerError
	// 	err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	// 	return
	// }
	// Count success by school gender
	var successBySchoolGender = make([]data.SuccessBySchoolGenderResponse, 0)
	// err = service.UserService.Repository.CountAllGroupBySchoolGender(&successBySchoolGender)
	// if err != nil {
	// 	errCode = http.StatusInternalServerError
	// 	err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	// 	return
	// }

	result = &data.MonitoringResponseList{
		Data: &data.MonitoringResponse{
			Count: &data.CountResponse{
				Schools:   schoolsCount,
				Directors: directorCount,
				Teachers:  teacherCount,
				Students:  studentCount,
				Parents:   parentCount,
			},
			UsersByFeature:        usersByFeature,
			UsersByGender:         usersByGender,
			UsersByMonth:          usersByMonth,
			UsersByYear:           usersByYear,
			SuccessBySchool:       successBySchool,
			SuccessBySchoolGender: successBySchoolGender,
		},
	}
	return
}
