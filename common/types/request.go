package types

type DeleteMultipleRequest struct {
	List []int64 `json:"list" required:"true" doc:"List of ID of items to delete"`
}
