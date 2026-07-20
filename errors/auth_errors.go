package errors

import "errors"

var (
	ErrInvalidToken       = errors.New("Invalid or expired token.")
	ErrInvalidCredentials = errors.New("Invalid email or password.")
)
