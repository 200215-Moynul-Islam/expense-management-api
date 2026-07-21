package errors

import "errors"

var (
	ErrInvalidRequest = errors.New("Invalid request.")

	ErrNameRequired = errors.New("Name is required.")
	ErrNameTooLong  = errors.New("Name must not exceed 100 characters.")

	ErrEmailRequired    = errors.New("Email is required.")
	ErrInvalidEmail     = errors.New("Invalid email format.")
	ErrEmailTooLong     = errors.New("Email must not exceed 255 characters.")

	ErrPasswordRequired = errors.New("Password is required.")
	ErrPasswordTooShort = errors.New("Password must be at least 6 characters.")
	ErrPasswordTooLong  = errors.New("Password must not exceed 255 characters.")

	ErrCategoryNameRequired = errors.New("Category name is required.")
	ErrCategoryNameTooLong  = errors.New("Category name must not exceed 100 characters.")

	ErrCategoryIDRequired   = errors.New("Category ID is required.")
	ErrTitleRequired      = errors.New("Title is required.")
	ErrTitleTooLong       = errors.New("Title must not exceed 255 characters.")
	ErrAmountRequired     = errors.New("Amount is required.")
	ErrNoteTooLong        = errors.New("Note must not exceed 1000 characters.")
	ErrExpenseDateRequired = errors.New("Expense date is required.")
)
