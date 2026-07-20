package errors

import "errors"

var (
	ErrInvalidRequest   = errors.New("Invalid request.")
	ErrNameRequired     = errors.New("Name is required.")
	ErrNameTooLong      = errors.New("Name must not exceed 100 characters.")
	ErrEmailRequired    = errors.New("Email is required.")
	ErrInvalidEmail     = errors.New("Invalid email format.")
	ErrEmailTooLong     = errors.New("Email must not exceed 255 characters.")
	ErrPasswordRequired = errors.New("Password is required.")
	ErrPasswordTooShort = errors.New("Password must be at least 6 characters.")
	ErrPasswordTooLong  = errors.New("Password must not exceed 255 characters.")
)
