package errors

import "errors"

var (
	ErrEmailExists = errors.New("Email already exists.")

	ErrCategoryExists = errors.New("Category already exists.")

	ErrCategoryNotFound = errors.New("Category not found.")
	ErrForbiddenCategory = errors.New("You do not have access to this category.")
	ErrInvalidExpenseDate = errors.New("Invalid expense date. Use YYYY-MM-DD format.")
)
