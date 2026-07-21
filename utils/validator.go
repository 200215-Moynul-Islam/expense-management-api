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

func ValidateUpdateUserRequest(
	request dto.UpdateUserRequest,
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

func ValidateLoginRequest(
	request dto.LoginRequest,
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
		case "MaxSize":
			return appErrors.ErrEmailTooLong
		}

	case "Password":
		switch validationError.Name {
		case "Required":
			return appErrors.ErrPasswordRequired
		case "MinSize":
			return appErrors.ErrPasswordTooShort
		case "MaxSize":
			return appErrors.ErrPasswordTooLong
		}
	}

	return appErrors.ErrInvalidRequest
}

func ValidateCreateCategoryRequest(
	request dto.CreateCategoryRequest,
) error {

	validationEngine := validation.Validation{}

	_, err := validationEngine.Valid(&request)
	if err != nil {
		return err
	}

	if !validationEngine.HasErrors() {
		return nil
	}

	return mapCategoryValidationError(validationEngine.Errors[0])
}

func ValidateUpdateCategoryRequest(
	request dto.UpdateCategoryRequest,
) error {

	validationEngine := validation.Validation{}

	_, err := validationEngine.Valid(&request)
	if err != nil {
		return err
	}

	if !validationEngine.HasErrors() {
		return nil
	}

	return mapCategoryValidationError(validationEngine.Errors[0])
}

func mapCategoryValidationError(
	validationError *validation.Error,
) error {

	switch validationError.Field {

	case "Name":
		switch validationError.Name {
		case "Required":
			return appErrors.ErrCategoryNameRequired

		case "MaxSize":
			return appErrors.ErrCategoryNameTooLong
		}
	}

	return appErrors.ErrInvalidRequest
}

func ValidateCreateExpenseRequest(
	request dto.CreateExpenseRequest,
) error {

	validationEngine := validation.Validation{}

	_, err := validationEngine.Valid(&request)
	if err != nil {
		return err
	}

	if !validationEngine.HasErrors() {
		return nil
	}

	return mapExpenseValidationError(
		validationEngine.Errors[0],
	)
}

func ValidateUpdateExpenseRequest(
	request dto.UpdateExpenseRequest,
) error {

	validationEngine := validation.Validation{}

	_, err := validationEngine.Valid(&request)
	if err != nil {
		return err
	}

	if !validationEngine.HasErrors() {
		return nil
	}

	return mapExpenseValidationError(
		validationEngine.Errors[0],
	)
}

func ValidateGetExpensesRequest(
	request dto.GetExpensesRequest,
) error {

	if request.Page != nil && *request.Page < 1 {
		return appErrors.ErrInvalidPage
	}

	if request.Limit != nil && *request.Limit < 1 {
		return appErrors.ErrInvalidLimit
	}

	if (request.Page == nil) != (request.Limit == nil) {
		return appErrors.ErrInvalidPagination
	}

	if request.FromDate != nil &&
		request.ToDate != nil &&
		request.FromDate.After(*request.ToDate) {

		return appErrors.ErrInvalidDateRange
	}

	if request.SortBy != "" {
		switch request.SortBy {
		case "created_at", "expense_date", "amount":
		default:
			return appErrors.ErrInvalidSortBy
		}
	}

	if request.SortOrder != "" {
		switch request.SortOrder {
		case "asc", "desc":
		default:
			return appErrors.ErrInvalidSortOrder
		}
	}

	return nil
}

func mapExpenseValidationError(
	validationError *validation.Error,
) error {

	switch validationError.Field {

	case "CategoryID":
		if validationError.Name == "Required" {
			return appErrors.ErrCategoryIDRequired
		}

	case "Title":
		switch validationError.Name {
		case "Required":
			return appErrors.ErrTitleRequired
		case "MaxSize":
			return appErrors.ErrTitleTooLong
		}

	case "Amount":
		if validationError.Name == "Required" {
			return appErrors.ErrAmountRequired
		}

	case "Note":
		if validationError.Name == "MaxSize" {
			return appErrors.ErrNoteTooLong
		}

	case "ExpenseDate":
		if validationError.Name == "Required" {
			return appErrors.ErrExpenseDateRequired
		}
	}

	return appErrors.ErrInvalidRequest
}
