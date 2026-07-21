package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

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

func (c *ExpenseController) GetAll() {

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

	var request dto.GetExpensesRequest

	if categoryID, err := c.GetInt("category_id"); err == nil {
		request.CategoryID = &categoryID
	}

	if fromDate := c.GetString("from_date"); fromDate != "" {
		t, err := time.Parse("2006-01-02", fromDate)
		if err != nil {
			utils.SendJSONResponse(
				c.Ctx,
				http.StatusBadRequest,
				false,
				"Invalid from_date. Use YYYY-MM-DD.",
				nil,
			)
			return
		}
		request.FromDate = &t
	}

	if toDate := c.GetString("to_date"); toDate != "" {
		t, err := time.Parse("2006-01-02", toDate)
		if err != nil {
			utils.SendJSONResponse(
				c.Ctx,
				http.StatusBadRequest,
				false,
				"Invalid to_date. Use YYYY-MM-DD.",
				nil,
			)
			return
		}
		request.ToDate = &t
	}

	if page, err := c.GetInt("page"); err == nil {
		request.Page = &page
	}

	if limit, err := c.GetInt("limit"); err == nil {
		request.Limit = &limit
	}

	request.SortBy = c.GetString("sort_by")
	request.SortOrder = c.GetString("sort_order")

	expenses, err := c.expenseService.GetExpensesByUserID(
		userID,
		request,
	)
	if err != nil {
		c.handleError(err)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusOK,
		true,
		"Expenses retrieved successfully.",
		expenses,
	)
}

func (c *ExpenseController) GetByID() {

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

	id, err := c.GetInt(":id")
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid expense ID.",
			nil,
		)
		return
	}

	expense, err := c.expenseService.GetExpenseByID(
		id,
		userID,
	)
	if err != nil {
		c.handleError(err)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusOK,
		true,
		"Expense retrieved successfully.",
		expense,
	)
}

func (c *ExpenseController) Delete() {

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

	id, err := c.GetInt(":id")
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid expense ID.",
			nil,
		)
		return
	}

	err = c.expenseService.DeleteExpense(
		id,
		userID,
	)
	if err != nil {
		c.handleError(err)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusOK,
		true,
		"Expense deleted successfully.",
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
		errors.Is(err, appErrors.ErrInvalidExpenseDate),
		errors.Is(err, appErrors.ErrInvalidPage),
		errors.Is(err, appErrors.ErrInvalidLimit),
		errors.Is(err, appErrors.ErrInvalidPagination),
		errors.Is(err, appErrors.ErrInvalidDateRange),
		errors.Is(err, appErrors.ErrInvalidSortBy),
		errors.Is(err, appErrors.ErrInvalidSortOrder):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			err.Error(),
			nil,
		)

	case errors.Is(err, appErrors.ErrCategoryNotFound),
		errors.Is(err, appErrors.ErrExpenseNotFound):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusNotFound,
			false,
			err.Error(),
			nil,
		)

	case errors.Is(err, appErrors.ErrForbiddenCategory),
		errors.Is(err, appErrors.ErrForbiddenExpense):

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
