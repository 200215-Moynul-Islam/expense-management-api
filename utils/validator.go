package utils

import (
	"expense-management-api/dto"
	appErrors "expense-management-api/errors"

	"github.com/beego/beego/v2/core/validation"
)

func ValidateRegisterRequest(
	request dto.RegisterRequest,
) error {

	validationEngine := validation.Validation{}

	_, err := validationEngine.Valid(&request)
	if err != nil {
		return err
	}

	if !validationEngine.HasErrors() {
		return nil
	}

	return mapValidationError(validationEngine.Errors[0])
}

func mapValidationError(
	validationError *validation.Error,
) error {

	switch validationError.Field {

	case "Name":
		switch validationError.Name {
		case "Required":
			return appErrors.ErrNameRequired
		case "MaxSize":
			return appErrors.ErrNameTooLong
		}

	case "Email":
		switch validationError.Name {
		case "Required":
			return appErrors.ErrEmailRequired
		case "Email":
			return appErrors.ErrInvalidEmail
		}

	case "Password":
		switch validationError.Name {
		case "Required":
			return appErrors.ErrPasswordRequired
		case "MinSize":
			return appErrors.ErrPasswordTooShort
		}
	}

	return appErrors.ErrInvalidRequest
}
