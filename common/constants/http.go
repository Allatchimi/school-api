package constants

import "fmt"

const (
	TokenKey    = "bearer"
	UserIDKey   = "userID"
	SchoolIDKey = "schoolID"
	IssuerKey   = "issuer"
	PlatformKey = "platform"
	DeviceKey   = "device"
	AppKey      = "app"
	CodeKey     = "code"
)

var Http500ErrorMessage = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("Error occurred when trying to %s! Please try again later.", message))
}

var Http400BadRequestErrorMessage = func() error {
	return fmt.Errorf("%s", "Bad request! Please enter valid information.")
}

var Http400BadRequestErrorMessageV2 = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("Invalid %s! Please enter valid information.", message))
}

var Http401InvalidTokenErrorMessage = func() error {
	return fmt.Errorf("%s", "Invalid or expired token! Please enter valid information.")
}

var Http401InvalidTokenErrorMessage2 = func(message string) error {
	return fmt.Errorf("%s", message)
}

var Http403InvalidPermissionErrorMessage = func() error {
	return fmt.Errorf("%s", "You don't have permission to access this resource! Please enter valid information.")
}

var Http409ConflictErrorMessage = func() error {
	return fmt.Errorf("%s", "Can't process now because of conflict! Please enter valid information.")
}

var Http422LockedErrorMessage = func() error {
	return fmt.Errorf("%s", "Invalid inputs! Please enter valid information.")
}

var Http423LockedErrorMessage = func() error {
	return fmt.Errorf("%s", "Can't process now because this operation is locked! Please try again later.")
}

var Http406ErrorMessage = func() error {
	return fmt.Errorf("%s", "The inputs are not acceptable! Please enter valid information.")
}

var Http404ErrorMessage = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("%s not found! Please enter valid information.", message))
}

var Http302ErrorMessage = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("This %s already exists! Please enter valid information.", message))
}

const DefaultBodySize = 10000 // 10Kb
