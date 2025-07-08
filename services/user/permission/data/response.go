package data

import (
	"api/common/types"
	"api/services/user/role/data"
)

type PermissionResponse struct {
	types.BaseGormModelResponse
	TableName string `json:"tableName" required:"false" doc:"Table name"`
	Create    bool   `json:"create" required:"false" doc:"Create permission"`
	Read      bool   `json:"read" required:"false" doc:"Read permission"`
	Update    bool   `json:"update" required:"false" doc:"Update permission"`
	Delete    bool   `json:"delete" required:"false" doc:"Delete permission"`

	Role *data.RoleResponse `json:"role" required:"false" doc:"Role" `
}

type PermissionListResponse struct {
	types.PaginatedResponse
	Data []PermissionResponse `json:"data" required:"false" doc:"List of all permissions" example:"[]"`
}
