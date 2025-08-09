package data

type RoleID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Role id"`
}

type RoleRequest struct {
	Feature     string `json:"feature" required:"true" enum:"dashboard-admin,dashboard-director,dashboard-teacher,dashboard-student,dashboard-parent,dashboard-default" doc:"Feature name"`
	Name        string `json:"name" required:"true" doc:"Role name"`
	Description string `json:"description" required:"false" doc:"Role description"`
}

type GetAllRequest struct {
	Feature string `json:"feature" query:"feature" required:"false" doc:"Feature name"`
}
