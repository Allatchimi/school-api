package data

import (
	"api/common/types"
)

type RoleResponse struct {
	types.BaseGormModelResponse
	Feature     string `json:"feature" required:"false" doc:"Feature name"`
	Name        string `json:"name" required:"false" doc:"Role name"`
	Description string `json:"description" required:"false" doc:"Role description"`
}

type RoleResponseList struct {
	types.PaginatedResponse
	Data []RoleResponse `json:"data" required:"false" doc:"List of role"`
}
