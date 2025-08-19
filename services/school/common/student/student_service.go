package student

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"api/common/constants"
	"api/common/helpers"
	googleMailHelper "api/common/helpers/message/mail/google"

	"api/common/types"
	"api/common/utils"
	"api/config"
	serviceHelperMessage "api/services/helper/message"
	"api/services/school/common/school"
	"api/services/school/common/student/data"
	"api/services/school/common/student/model"
	"api/services/user/role"
	"api/services/user/user"
	dataUser "api/services/user/user/data"
	modelUser "api/services/user/user/model"

	"go.uber.org/zap"
)

type Service struct {
	Repository    *Repository
	RoleService   *role.Service
	UserService   *user.Service
	SchoolService *school.Service
}

func NewService(
	repository *Repository,
	roleService *role.Service,
	userService *user.Service,
	schoolService *school.Service,
) *Service {
	return &Service{
		Repository:    repository,
		RoleService:   roleService,
		UserService:   userService,
		SchoolService: schoolService,
	}
}

const MODEL_NAME = "student"
const DEFAULT_ERROR_MESSAGE = "interact with student model"
const uidAcceptedLetters = "ABCEFGHIJKLMNOPQRSUVWXYZ"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.StudentRequest,
) (result *model.Student, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleStudent)
	if errRole != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Get the school
	foundSchool, errSchool := service.SchoolService.Repository.GetByID(newRequest.SchoolID)
	if errSchool != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundSchool == nil || foundSchool.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("school")
		return
	}

	// Check request
	if request == nil || newRequest.Info == nil {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Generate the email
	newEmail := newRequest.Email
	if newRequest.AutoGenerateEmail {
		if foundSchool.Config == nil {
			errCode = http.StatusNotFound
			err = constants.Http404ErrorMessage("school")
			return
		}
		newEmail = helpers.GenerateEmailFromFullName(
			newRequest.Info.FirstName,
			newRequest.Info.LastName,
			foundSchool.Config.UserEmailDomainName,
		)
	}

	// Format request
	var item = &dataUser.UserRequest{
		SchoolID:    newRequest.SchoolID,
		RoleID:      userRole.ID,
		Email:       newEmail,
		PhoneNumber: newRequest.PhoneNumber,
		IsActivated: true,
		Status:      newRequest.Status,
		Info: &dataUser.UserInfoRequest{
			Gender:        newRequest.Info.Gender,
			Username:      newRequest.Info.Username,
			FirstName:     newRequest.Info.FirstName,
			LastName:      newRequest.Info.LastName,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}

	// Generate password
	password := helpers.GeneratePasswordFromUser(
		item.Info.FirstName,
		item.Info.LastName,
		item.Info.Birthday,
	)

	// Create user
	createdUser, errCodeCreate, errCreate := service.UserService.Create(ctxData, item, &password)
	if errCreate != nil {
		errCode = errCodeCreate
		err = errCreate
		return
	}
	if createdUser == nil || createdUser.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Generate the uid if it is empty
	var newUID = newRequest.UID
	if len(newUID) < 1 {
		// Get all uids to exclude
		uids, errUids := service.Repository.GetAll(nil, nil, &data.GetAllRequest{SchoolID: newRequest.SchoolID})
		if errUids != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		formattedUids := make([]string, len(uids))
		for i, uid := range uids {
			formattedUids[i] = uid.UID
		}
		// Generate uid
		var errGenerateUID error
		newUID, errGenerateUID = utils.GenerateRandomUID(
			fmt.Sprintf("%02d", time.Now().Year()),
			uidAcceptedLetters,
			formattedUids,
		)
		if errGenerateUID != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	// Create
	result, err = service.Repository.Create(&model.Student{
		SchoolID: newRequest.SchoolID,
		UserID:   createdUser.ID,
		UID:      newUID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create google workspace user
	go func() {
		if result.School == nil || result.School.Config == nil ||
			result.User == nil || result.User.Info == nil || result.User.Config == nil {
			return
		}
		if !request.AutoGenerateEmail {
			return
		}
		userGoogle, errGoogle := googleMailHelper.CreateGoogleWorkspaceUser(
			context.Background(),
			result.School.Config.GoogleWorkspaceCredentials,
			result.School.Config.GoogleWorkspaceUserEmailDomain,
			result.User,
			password,
		)
		if errGoogle != nil || userGoogle == nil {
			helpers.Logger.Warn(
				"Failed to create Google Workspace user!",
				zap.Error(errGoogle))
			return
		}
		helpers.Logger.Info("Google Workspace user created!", zap.String("Email", userGoogle.PrimaryEmail))
	}()
	return
}

func (service *Service) CreateStudentEnroll(
	ctxData *types.ContextData,
	request *data.StudentEnrollRequest,
) (result *model.StudentEnroll, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.StudentEnroll{
		SchoolID:      newRequest.SchoolID,
		YearID:        newRequest.YearID,
		ClassID:       newRequest.ClassID,
		LevelDomainID: newRequest.LevelDomainID,
		StudentID:     newRequest.StudentID,

		Origin: constants.STUDENT_ENROLL_ORIGIN_DASHBOARD,
	}

	// Check unique
	foundUnique, err := service.Repository.GetStudentEnrollUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreStudentEnrollSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create
	result, err = service.Repository.CreateStudentEnroll(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateStudentPreEnroll(
	ctxData *types.ContextData,
	request *data.StudentPreEnrollRequest,
) (result *model.StudentPreEnroll, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.StudentPreEnroll{
		SchoolID:      newRequest.SchoolID,
		YearID:        newRequest.YearID,
		ClassID:       newRequest.ClassID,
		LevelDomainID: newRequest.LevelDomainID,
		UserID:        ctxData.Jwt.UserID,

		Status:        constants.STUDENT_PRE_ENROLL_STATUS_INITIATED,
		Message:       newRequest.Message,
		Gender:        newRequest.Gender,
		FirstName:     newRequest.FirstName,
		LastName:      newRequest.LastName,
		Birthday:      newRequest.Birthday,
		BirthLocation: newRequest.BirthLocation,
		Document1:     newRequest.Document1,
		Document2:     newRequest.Document2,
		Document3:     newRequest.Document3,
		Document4:     newRequest.Document4,
		Document5:     newRequest.Document5,
	}

	// Check unique
	foundUnique, err := service.Repository.GetStudentPreEnrollUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreStudentPreEnrollSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create
	result, err = service.Repository.CreateStudentPreEnroll(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil || result.ID < 1 {
		return
	}
	if result.School == nil || result.School.Info == nil {
		return
	}

	// Send message
	go func() {
		var msgTitle, msgBody, msgClassLevelDomain string
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.Class != nil {
				msgClassLevelDomain = fmt.Sprintf("class %s", result.Class.Name)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.LevelDomain != nil && result.LevelDomain.Level != nil && result.LevelDomain.Domain != nil {
				msgClassLevelDomain = fmt.Sprintf("level domain %s %s", result.LevelDomain.Level.Name, result.LevelDomain.Domain.Name)
			}
		}
		switch result.Status {
		case constants.STUDENT_PRE_ENROLL_STATUS_INITIATED:
			msgTitle = fmt.Sprintf("Enrollment request received for %s", msgClassLevelDomain)
			msgBody = fmt.Sprintf(`Dear %s %s,
		Thank you for your enrollment request to %s for %s.
		We have successfully received your application and our admissions team is currently reviewing it. 
		You will receive a notification with our decision as soon as the review process is complete.
		If you have any questions in the meantime, please don't hesitate to contact our support team.
		Best regards,
		The Admissions Team`,
				result.FirstName, result.LastName, result.School.Info.FullName, msgClassLevelDomain)

		}
		serviceHelperMessage.SendMessage(
			&serviceHelperMessage.MessageRequest{
				PusNotification: true,
				Mail:            true,
			},
			msgTitle,
			msgBody,
			result.School,
			"",
			[]modelUser.User{*result.User},
		)
	}()
	return
}

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.StudentRequest,
) (result *model.Student, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.Student
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Generate the uid if it is empty
	var newUID = newRequest.UID
	if len(newUID) < 1 {
		// Get all uids to exclude
		uids, errUids := service.Repository.GetAll(nil, nil, &data.GetAllRequest{SchoolID: newRequest.SchoolID})
		if errUids != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		formattedUids := make([]string, len(uids))
		for i, uid := range uids {
			formattedUids[i] = uid.UID
		}
		// Generate uid
		var errGenerateUID error
		newUID, errGenerateUID = utils.GenerateRandomUID(
			fmt.Sprintf("%02d", time.Now().Year()),
			uidAcceptedLetters,
			formattedUids,
		)
		if errGenerateUID != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	// Check unique by uid
	foundUnique, err := service.Repository.GetUniqueObjectByUID(&model.Student{
		SchoolID: foundItem.SchoolID,
		UID:      newUID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsByUID(
		foundUnique,
		&model.Student{
			SchoolID: foundItem.SchoolID,
			UID:      newUID,
		},
	) && !service.Repository.AreSameUniqueObjectsByUID(
		foundUnique,
		foundItem,
	) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Get the role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleStudent)
	if errRole != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Get the school
	foundSchool, errSchool := service.SchoolService.Repository.GetByID(newRequest.SchoolID)
	if errSchool != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundSchool == nil || foundSchool.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("school")
		return
	}

	// Update user
	userRequest := dataUser.UserRequest{
		SchoolID:    newRequest.SchoolID,
		RoleID:      userRole.ID,
		Email:       newRequest.Email,
		PhoneNumber: newRequest.PhoneNumber,
		IsActivated: true,
		Status:      newRequest.Status,
		Info: &dataUser.UserInfoRequest{
			Gender:        newRequest.Info.Gender,
			Username:      newRequest.Info.Username,
			FirstName:     newRequest.Info.FirstName,
			LastName:      newRequest.Info.LastName,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}
	_, errCodeUser, errUser := service.UserService.Update(ctxData, foundItem.UserID, &userRequest)
	if errUser != nil {
		errCode = errCodeUser
		err = errUser
		return
	}

	// Update student
	result, err = service.Repository.UpdateByID(id, &model.Student{
		SchoolID: foundItem.SchoolID,
		UserID:   foundItem.UserID,
		UID:      newUID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateStudentEnroll(
	ctxData *types.ContextData,
	id int64,
	request *data.StudentEnrollRequest,
) (result *model.StudentEnroll, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.StudentEnroll
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetStudentEnrollByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetStudentEnrollByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format request
	item := &model.StudentEnroll{
		SchoolID:      newRequest.SchoolID,
		YearID:        newRequest.YearID,
		ClassID:       newRequest.ClassID,
		LevelDomainID: newRequest.LevelDomainID,
		StudentID:     newRequest.StudentID,
	}

	// Check unique
	foundUnique, err := service.Repository.GetStudentEnrollUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreStudentEnrollSameUniqueObjects(foundUnique, item) && !service.Repository.AreStudentEnrollSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update
	result, err = service.Repository.UpdateStudentEnrollByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateStudentPreEnroll(
	ctxData *types.ContextData,
	id int64,
	request *data.StudentPreEnrollRequest,
) (result *model.StudentPreEnroll, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.StudentPreEnroll
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetStudentPreEnrollByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetStudentPreEnrollByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check if the status is ok
	if foundItem.Status != constants.STUDENT_PRE_ENROLL_STATUS_INITIATED {
		errCode = http.StatusLocked
		err = constants.Http423LockedErrorMessage()
		return
	}

	// Format request
	item := &model.StudentPreEnroll{
		SchoolID:      newRequest.SchoolID,
		YearID:        newRequest.YearID,
		ClassID:       newRequest.ClassID,
		LevelDomainID: newRequest.LevelDomainID,

		Message:       newRequest.Message,
		Gender:        newRequest.Gender,
		FirstName:     newRequest.FirstName,
		LastName:      newRequest.LastName,
		Birthday:      newRequest.Birthday,
		BirthLocation: newRequest.BirthLocation,
		Document1:     newRequest.Document1,
		Document2:     newRequest.Document2,
		Document3:     newRequest.Document3,
		Document4:     newRequest.Document4,
		Document5:     newRequest.Document5,
	}

	// Check unique
	foundUnique, err := service.Repository.GetStudentPreEnrollUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreStudentPreEnrollSameUniqueObjects(foundUnique, item) && !service.Repository.AreStudentPreEnrollSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update
	result, err = service.Repository.UpdateStudentPreEnrollByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateStudentPreEnrollStatus(
	ctxData *types.ContextData,
	id int64,
	request *data.StudentPreEnrollStatusRequest,
) (result *model.StudentPreEnroll, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.StudentPreEnroll
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetStudentPreEnrollByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetStudentPreEnrollByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check status
	if foundItem.Status == constants.STUDENT_PRE_ENROLL_STATUS_ENROLLED ||
		foundItem.Status == constants.STUDENT_PRE_ENROLL_STATUS_REJECTED {
		errCode = http.StatusConflict
		err = constants.Http409ConflictErrorMessage()
		return
	}

	// Format request
	item := &model.StudentPreEnroll{
		Status:         request.Status,
		StatusFeedback: request.StatusFeedback,
	}

	// Update
	result, err = service.Repository.UpdateStudentPreEnrollStatusByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Check the new status
	if !(result.Status == constants.STUDENT_PRE_ENROLL_STATUS_ENROLLED ||
		result.Status == constants.STUDENT_PRE_ENROLL_STATUS_REJECTED) {
		return
	}

	// Create new student if the status is enrolled
	var createdStudent *model.Student
	var createdStudentEnroll *model.StudentEnroll
	if result.Status == constants.STUDENT_PRE_ENROLL_STATUS_ENROLLED {
		createdStudent, errCode, err = service.Create(ctxData, &data.StudentRequest{
			SchoolID:          foundItem.SchoolID,
			AutoGenerateEmail: true,
			Status:            constants.USER_STATUS_ENABLED,
			Info: &dataUser.UserInfoRequest{
				FirstName:     foundItem.FirstName,
				LastName:      foundItem.LastName,
				Gender:        foundItem.Gender,
				Birthday:      foundItem.Birthday,
				BirthLocation: foundItem.BirthLocation,
			},
		})
		if !(err != nil || createdStudent == nil || createdStudent.ID < 1) {
			createdStudentEnroll, errCode, err = service.CreateStudentEnroll(ctxData, &data.StudentEnrollRequest{
				SchoolID:      foundItem.SchoolID,
				YearID:        foundItem.YearID,
				ClassID:       foundItem.ClassID,
				LevelDomainID: foundItem.LevelDomainID,
				StudentID:     createdStudent.ID,

				Origin:         constants.STUDENT_ENROLL_ORIGIN_PRE_ENROLL,
				OriginFeedback: "Enrolled form pre enrollment request.",
			})
		}
		if err != nil || createdStudent == nil || createdStudent.ID < 1 || createdStudentEnroll == nil || createdStudentEnroll.ID < 1 {
			var errMsg string
			if errCode == http.StatusFound {
				errMsg = "A similar student enroll already exists!"
			} else {
				errMsg = "Automatically rejected by the system! Please try again later."
			}
			var tempErr error
			result, tempErr = service.Repository.UpdateStudentPreEnrollStatusByID(id, &model.StudentPreEnroll{
				Status:         constants.STUDENT_PRE_ENROLL_STATUS_REJECTED,
				StatusFeedback: errMsg,
			})
			if tempErr != nil {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			}
			if err == nil {
				errCode = http.StatusInternalServerError
				err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			}
		} else {
			createdStudent, err = service.Repository.GetByID(createdStudent.ID)
		}
	}

	if err != nil {
		return
	}
	if createdStudent == nil || createdStudent.User == nil || createdStudent.User.Info == nil ||
		result.School == nil || result.School.Info == nil {
		return
	}

	// Send message
	go func() {
		var msgTitle, msgBody, msgClassLevelDomain string
		switch result.School.Type {
		case constants.SCHOOL_TYPE_HIGHSCHOOL:
			if result.Class != nil {
				msgClassLevelDomain = fmt.Sprintf("class %s", result.Class.Name)
			}
		case constants.SCHOOL_TYPE_UNIVERSITY:
			if result.LevelDomain != nil && result.LevelDomain.Level != nil && result.LevelDomain.Domain != nil {
				msgClassLevelDomain = fmt.Sprintf("level domain %s %s", result.LevelDomain.Level.Name, result.LevelDomain.Domain.Name)
			}
		}
		switch result.Status {
		case constants.STUDENT_PRE_ENROLL_STATUS_ENROLLED:
			msgTitle = fmt.Sprintf("Enrollment accepted for %s!", msgClassLevelDomain)
			msgBody = fmt.Sprintf(`Hi %s %s, welcome to %s! 
			Please visit our website and log in with your credentials. 
			Your new email is %s, and your default password is a combination of your first name, first last name, and birth year/enrolled year. 
			For example: For a user with first name "John Durand", last name "Carmack Benie" and birthday "2010/06/13", the default password would be JohnCarmack2010. 
			If this doesn't work, please contact our support team through the website. Thank you.`, createdStudent.User.Info.FirstName, createdStudent.User.Info.LastName, result.School.Info.FullName, createdStudent.User.Email)

		case constants.STUDENT_PRE_ENROLL_STATUS_REJECTED:
			msgTitle = fmt.Sprintf("Enrollment rejected for %s!", msgClassLevelDomain)
			msgBody = fmt.Sprintf("Your enrollment for %s has been rejected. Please check your account dashboard for more details.", msgClassLevelDomain)
		}

		serviceHelperMessage.SendMessage(
			&serviceHelperMessage.MessageRequest{
				PusNotification: true,
				Mail:            true,
			},
			msgTitle,
			msgBody,
			result.School,
			"",
			[]modelUser.User{*foundItem.User},
		)
	}()
	return
}

func (service *Service) Delete(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.Student
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteStudentEnroll(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.StudentEnroll
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetStudentEnrollByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetStudentEnrollByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteStudentEnrollByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteStudentPreEnroll(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.StudentPreEnroll
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetStudentPreEnrollByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetStudentPreEnrollByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete
	affectedRows, err = service.Repository.DeleteStudentPreEnrollByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultiple(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleByID(list, ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultipleStudentEnroll(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleStudentEnrollByID(list, ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultipleStudentPreEnroll(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleStudentPreEnrollByID(list, ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) Get(
	ctxData *types.ContextData,
	id int64,
) (result *model.Student, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetStudentEnroll(
	ctxData *types.ContextData,
	id int64,
) (result *model.StudentEnroll, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetStudentEnrollByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetStudentEnrollByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetStudentPreEnroll(
	ctxData *types.ContextData,
	id int64,
) (result *model.StudentPreEnroll, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetStudentPreEnrollByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetStudentPreEnrollByID(id)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Student, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllStudentEnroll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllStudentEnrollRequest,
) (result []model.StudentEnroll, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllStudentEnroll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllStudentPreEnroll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllStudentPreEnrollRequest,
) (result []model.StudentPreEnroll, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	if ctxData.User.Feature != constants.FeatureAdmin && ctxData.User.Feature != constants.FeatureDirector {
		newRequest.UserID = ctxData.Jwt.UserID
	}
	result, err = service.Repository.GetAllStudentPreEnroll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllPublic(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Student, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
