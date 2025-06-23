package data

type PermissionID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Permission id"`
}

type PermissionRoleID struct {
	RoleID int64 `json:"roleID" path:"roleID" required:"true" doc:"Role id"`
}

type UpdatePermissionRequest struct {
	TableName string `json:"tableName" required:"true" minLength:"2" doc:"Table name"`
	Create    bool   `json:"create" required:"true" doc:"Create permission"`
	Read      bool   `json:"read" required:"true" doc:"Read permission"`
	Update    bool   `json:"update" required:"true" doc:"Update permission"`
	Delete    bool   `json:"delete" required:"true" doc:"Delete permission"`
}

type GetAllRequest struct {
	RoleID    int64  `json:"roleID" query:"roleID" required:"false" doc:"Role id"`
	TableName string `json:"tableName" query:"tableName" required:"false" minLength:"2" doc:"Table name"`
}
