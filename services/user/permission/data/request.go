package data

type PermissionID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Permission id" example:"1"`
}

type PermissionRoleID struct {
	RoleID int64 `json:"roleID" path:"roleID" required:"true" doc:"Role id" example:"1"`
}

type UpdatePermissionRequest struct {
	TableName string `json:"tableName" required:"true" minLength:"2" doc:"Table name" example:"users"`
	Create    bool   `json:"create" required:"true" doc:"Create permission" example:"false"`
	Read      bool   `json:"read" required:"true" doc:"Read permission" example:"false"`
	Update    bool   `json:"update" required:"true" doc:"Update permission" example:"false"`
	Delete    bool   `json:"delete" required:"true" doc:"Delete permission" example:"false"`
}

type GetAllRequest struct {
	RoleID    int64  `json:"roleID" query:"roleID" required:"false" doc:"Role id" example:"1"`
	TableName string `json:"tableName" query:"tableName" required:"false" minLength:"2" doc:"Table name" example:"users"`
}
