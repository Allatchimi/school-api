package types

type DefaultResponse struct {
	Message string `json:"message" required:"false" doc:"Message"`
}

type ErrorResponse struct {
	Message string `json:"message" required:"false" doc:"Message"`
}

type PaginatedResponse struct {
	Filter     *Filter     `json:"filter" required:"false" doc:"Filter"`
	Pagination *Pagination `json:"pagination" required:"false" doc:"Pagination"`
}

type DeletedResponse struct {
	AffectedRows int64 `json:"affectedRows" required:"false" doc:"Number of row affected with this delete"`
}
