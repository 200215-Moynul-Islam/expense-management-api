package errors

import "errors"

var (
	ErrInvalidRequest   = errors.New("Invalid request.")
	ErrNameRequired     = errors.New("Name is required.")
	ErrNameTooLong      = errors.New("Name must not exceed 100 characters.")
	ErrEmailRequired    = errors.New("Email is required.")
	ErrInvalidEmail     = errors.New("Invalid email format.")
	ErrPasswordRequired = errors.New("Password is required.")
	ErrPasswordTooShort = errors.New("Password must be at least 6 characters.")
)
