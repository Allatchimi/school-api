package helpers

import (
	"fmt"
	"strings"

	"api/common/constants"
	"api/common/types"

	"gorm.io/gorm"
)

// PaginationScopeV2 Returns *gorm.DB pointer with applied search, filter and pagination
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
func PaginationScopeV2(
	db *gorm.DB,
	selection string,
	where string,
	pagination *types.Pagination,
	filter *types.Filter,
	args ...any,
) func(*gorm.DB) *gorm.DB {
	// Count total records using the WHERE clause and bound arguments
	if pagination != nil {
		var count int64
		countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s %s) AS subquery", selection, where)
		db.Raw(countQuery, args...).Count(&count)
		pagination.UpdateFields(count)
	}

	// Build pagination (ORDER BY + LIMIT + OFFSET)
	var paginationFilter string
	if pagination != nil && filter != nil {
		paginationFilter = fmt.Sprintf(
			"ORDER BY %s %s LIMIT %d OFFSET %d",
			filter.OrderBy,
			filter.Sort,
			pagination.Limit,
			pagination.Offset,
		)
	} else if filter != nil {
		paginationFilter = fmt.Sprintf(
			"ORDER BY %s %s",
			filter.OrderBy,
			filter.Sort,
		)
	} else if pagination != nil {
		paginationFilter = fmt.Sprintf(
			"ORDER BY %s %s LIMIT %d OFFSET %d",
			"updated_at",
			"desc",
			pagination.Limit,
			pagination.Offset,
		)
	}

	// Return scoped function to apply raw SQL with parameters
	return func(tx *gorm.DB) *gorm.DB {
		rawQuery := fmt.Sprintf("%s %s %s", selection, where, paginationFilter)
		return tx.Raw(rawQuery, args...)
	}
}

// GetPaginationFiltersFromQuery Checks the entries and return the corrected ones.
func GetPaginationFiltersFromQuery(filter *types.Filter, pagination *types.PaginationRequest) (pageResult *types.Pagination, filterResult *types.Filter) {
	if filter == nil && pagination == nil {
		return
	}
	if filter != nil {
		if len(strings.TrimSpace(filter.OrderBy)) <= 0 {
			filter.OrderBy = constants.FilterOrderByDefault
		}
		if filter.Sort != "asc" {
			filter.Sort = constants.FilterSortDefault
		}
		filterResult = filter
	}

	if pagination != nil {
		page := pagination.Page
		limit := pagination.Limit

		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = constants.PaginationLimitDefault
		}
		pageResult = NewPaginationData(page, limit)
	}

	return
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
