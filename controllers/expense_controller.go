package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"expense-management-api/dto"
	appErrors "expense-management-api/errors"
	"expense-management-api/repositories"
	"expense-management-api/services"
	"expense-management-api/utils"
)

type ExpenseController struct {
	BaseController
	expenseService services.ExpenseService
}

func (c *ExpenseController) Prepare() {
	c.expenseService = services.NewExpenseService(
		repositories.NewExpenseRepository(),
		repositories.NewCategoryRepository(),
	)
}

func (c *ExpenseController) Create() {

	var request dto.CreateExpenseRequest

	err := json.Unmarshal(
		c.Ctx.Input.RequestBody,
		&request,
	)
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid request body.",
			nil,
		)
		return
	}

	userID, ok := c.getUserID()
	if !ok {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusUnauthorized,
			false,
			"Unauthorized.",
			nil,
		)
		return
	}

	err = c.expenseService.CreateExpense(
		request,
		userID,
	)
	if err != nil {
		c.handleError(err)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusCreated,
		true,
		"Expense created successfully.",
		nil,
	)
}

func (c *ExpenseController) handleError(
	err error,
) {

	switch {

	case errors.Is(err, appErrors.ErrInvalidRequest),
		errors.Is(err, appErrors.ErrCategoryIDRequired),
		errors.Is(err, appErrors.ErrTitleRequired),
		errors.Is(err, appErrors.ErrTitleTooLong),
		errors.Is(err, appErrors.ErrAmountRequired),
		errors.Is(err, appErrors.ErrNoteTooLong),
		errors.Is(err, appErrors.ErrExpenseDateRequired),
		errors.Is(err, appErrors.ErrInvalidExpenseDate):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			err.Error(),
			nil,
		)

	case errors.Is(err, appErrors.ErrCategoryNotFound):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusNotFound,
			false,
			err.Error(),
			nil,
		)

	case errors.Is(err, appErrors.ErrForbiddenCategory):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusForbidden,
			false,
			err.Error(),
			nil,
		)

	default:

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusInternalServerError,
			false,
			"Internal server error.",
			nil,
		)
	}
}
