package errors

import "errors"

var (
	ErrEmailExists = errors.New("Email already exists.")
)
