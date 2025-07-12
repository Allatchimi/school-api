package types

import "api/common/constants"

type PaginationRequest struct {
	Page  int `json:"page" query:"page" required:"false" doc:"Current page"`
	Limit int `json:"limit" query:"limit" minimum:"1" required:"false" doc:"Max items per page"`
}

type Pagination struct {
	CurrentPage  int   `json:"currentPage" required:"false"`
	NextPage     int   `json:"nextPage" required:"false"`
	PreviousPage int   `json:"previousPage" required:"false"`
	TotalPages   int64 `json:"totalPages" required:"false"`
	Count        int64 `json:"count" required:"false"`
	Limit        int   `json:"limit" required:"false"`
	Offset       int   `json:"offset" required:"false"`
}

// UpdateFields Updates pagination parameters based on total item count.
func (p *Pagination) UpdateFields(count int64) {
	if p.Limit < 1 {
		p.Limit = constants.PaginationLimitDefault
	}
	p.Count = count                              // Update count
	DivUp(&count, &p.Limit, &p.TotalPages)       // Update total pages
	NextPage(&p.NextPage, &p.TotalPages)         // Update next page
	PreviousPage(&p.PreviousPage, &p.TotalPages) // Update previous page
	// Update current page
	if int64(p.CurrentPage) > p.TotalPages {
		p.CurrentPage = int(p.TotalPages)
	} else if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}
	// Update offset
	if p.CurrentPage > 1 {
		var tempOffset = p.CurrentPage - 1
		p.Offset = tempOffset * p.Limit
	} else {
		p.Offset = 0
	}
}

// DivUp Calculates the ceiling of the division between numerator and denominator.
func DivUp(numerator *int64, denominator *int, result *int64) {
	*result = 1 + (*numerator-1)/int64(*denominator)
}

// NextPage Calculates the next page number based on current page and total pages.
func NextPage(page *int, totalPages *int64) {
	if int64(*page) <= 0 {
		*page = 1
	}
	if int64(*page) < *totalPages {
		*page++
		return
	}
	*page = int(*totalPages)
}

// PreviousPage Calculates the previous page number based on current page and total pages.
func PreviousPage(page *int, totalPages *int64) {
	if int64(*page) > *totalPages {
		*page = int(*totalPages)
	}
	if int64(*page) > 1 {
		*page--
		return
	}
	*page = 1
}
