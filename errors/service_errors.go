package errors

import "errors"

var (
	ErrEmailExists = errors.New("Email already exists.")

	ErrCategoryExists = errors.New("Category already exists.")
)
