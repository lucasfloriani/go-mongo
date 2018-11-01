package helper

import (
	"strconv"

	"github.com/labstack/echo"
)

const (
	// DefaultPageSize set default size of pagination
	DefaultPageSize int = 10
	// MaxPageSize set max size of pagination
	MaxPageSize int = 15
)

// PaginatedList represents a paginated list of data items.
type PaginatedList struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

// NewPaginatedList creates a new Paginated instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func NewPaginatedList(page, perPage, total int) *PaginatedList {
	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &PaginatedList{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// Offset returns the OFFSET value that can be used in a SQL statement.
func (p *PaginatedList) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the LIMIT value that can be used in a SQL statement.
func (p *PaginatedList) Limit() int {
	return p.PerPage
}

// GetPaginatedListFromRequest query by pagination parameters
// and returns a list
func GetPaginatedListFromRequest(c echo.Context, count int) *PaginatedList {
	page := parseInt(c.QueryParam("page"), 1)
	perPage := parseInt(c.QueryParam("per_page"), DefaultPageSize)
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	return NewPaginatedList(page, perPage, count)
}

// parseInt check string value and try to convert to integer,
// if ok returns converted value, else returns the defaultValue
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
