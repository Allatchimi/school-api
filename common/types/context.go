package types

type ContextUserData struct {
	RoleID  int64  `json:"roleID"`
	Feature string `json:"feature"`

	DirectorID int64 `json:"directorID"`
	TeacherID  int64 `json:"teacherID"`
	StudentID  int64 `json:"studentID"`
	ParentID   int64 `json:"parentID"`
}

type ContextData struct {
	Jwt  *JwtToken
	User *ContextUserData
}
