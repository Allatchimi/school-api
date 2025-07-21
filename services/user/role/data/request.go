package data

type RoleID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Role id" example:"1"`
}

type RoleRequest struct {
	Name        string `json:"name" required:"true" minLength:"2" maxLength:"30" doc:"Role name" example:"client"`
	Feature     string `json:"feature" required:"true" enum:"dashboard-admin,dashboard-director,dashboard-teacher,dashboard-student,dashboard-parent,dashboard-default" doc:"Feature name" example:"dashboard-admin"`
	Description string `json:"description" required:"false" doc:"Role description" example:"Client role used to allow users to access your services"`
}

type GetAllRequest struct {
	Feature string `json:"feature" query:"feature" required:"false" doc:"Feature name" example:"dashboard-admin"`
}
