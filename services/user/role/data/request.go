package data

type RoleID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Role id" example:"1"`
}

type RoleRequest struct {
	Name        string `json:"name" required:"true" min:"2" max:"30" doc:"Role name" example:"client"`
	Feature     string `json:"feature" required:"true" enum:"feature-admin,feature-director,feature-teacher,feature-student,feature-parent,feature-default" doc:"Feature name" example:"feature-admin"`
	Description string `json:"description" required:"false" doc:"Role description" example:"Client role used to allow users to access your services"`
}

type GetAllRequest struct {
	Feature string `json:"feature" query:"feature" required:"false" doc:"Feature name" example:"feature-admin"`
}
