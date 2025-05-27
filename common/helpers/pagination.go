package helpers

import (
	"fmt"
	"strings"

	"api/common/constants"
	"api/common/types"

	"gorm.io/gorm"
)

// PaginationScope Returns *gorm.DB pointer with applied search, filter and pagination
//
// - Search performs a full-text search on a specific rows of the table.
// searchColumns is the specific rows.
// E.g: whereSearch = "title || ' ' || description".
// The official documentation for full text search in PostgresSQL can be found here
// https://www.postgresql.org/docs/17/textsearch-tables.html#TEXTSEARCH-TABLES-SEARCH
//
// - Filter applies an ORDER BY clause to the specified field name,
// sorting the results in ascending or descending based on the Sort parameter.
//
// - Pagination applies an offset and limit to the results, determining which subset of data to display.
func PaginationScope(db *gorm.DB, selection string, where string, pagination *types.Pagination, filter *types.Filter) func(*gorm.DB) *gorm.DB {
	var count *int64 = new(int64)
	db.Raw(fmt.Sprintf("SELECT COUNT(*) FROM (%s %s) AS subquery;", selection, where)).Count(count)
	pagination.UpdateFields(*count)

	var paginationFilter = ""
	if pagination != nil && filter != nil {
		paginationFilter = fmt.Sprintf(
			"ORDER BY %s %s LIMIT %d OFFSET %d",
			filter.OrderBy,
			filter.Sort,
			pagination.Limit,
			pagination.Offset,
		)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Raw(fmt.Sprintf("%s %s %s;",
			selection,
			where,
			paginationFilter,
		))
	}
}

// GetPaginationFiltersFromQuery Checks the entries and return the corrected ones.
func GetPaginationFiltersFromQuery(filter *types.Filter, pagination *types.PaginationRequest) (*types.Pagination, *types.Filter) {
	page := pagination.Page
	limit := pagination.Limit

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = constants.PaginationLimitDefault
	}
	if len(strings.TrimSpace(filter.OrderBy)) <= 0 {
		filter.OrderBy = constants.FilterOrderByDefault
	}
	if filter.Sort != "asc" {
		filter.Sort = constants.FilterSortDefault
	}

	return NewPaginationData(page, limit), filter
}

// NewPaginationData The user passes a pagination request, specifying the desired page and limit.
// We validate the inputs and return a new pagination object with the applied settings:
// current page, next page, previous page, total pages, count, limit, and offset.
func NewPaginationData(page int, limit int) *types.Pagination {
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	return &types.Pagination{
		CurrentPage:  page,
		NextPage:     page,
		PreviousPage: page,
		TotalPages:   0,
		Count:        0,
		Limit:        limit,
		Offset:       offset,
	}
}
