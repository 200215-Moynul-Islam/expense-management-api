package errors

import "errors"

var (
	ErrEmailExists = errors.New("Email already exists.")

	ErrCategoryExists = errors.New("Category already exists.")

	ErrCategoryNotFound = errors.New("Category not found.")
	ErrForbiddenCategory = errors.New("You do not have access to this category.")
	ErrInvalidExpenseDate = errors.New("Invalid expense date. Use YYYY-MM-DD format.")

	ErrInvalidPage       = errors.New("Page must be greater than 0.")
	ErrInvalidLimit      = errors.New("Limit must be greater than 0.")
	ErrInvalidPagination = errors.New("Page and limit must be provided together.")
	ErrInvalidDateRange  = errors.New("From date cannot be after to date.")
	ErrInvalidSortBy     = errors.New("Invalid sort field.")
	ErrInvalidSortOrder  = errors.New("Invalid sort order.")
)
