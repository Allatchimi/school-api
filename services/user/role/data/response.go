package data

import (
	"api/common/types"
)

type RoleResponse struct {
	types.BaseGormModelResponse
	Name        string `json:"name" required:"false" doc:"Role name"`
	Feature     string `json:"feature" required:"false" doc:"Feature name"`
	Description string `json:"description" required:"false" doc:"Role description"`
}

type RoleResponseList struct {
	types.PaginatedResponse
	Data []RoleResponse `json:"data" required:"false" doc:"List of roles" example:"[]"`
}
