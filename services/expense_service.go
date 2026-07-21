package services

import (
	"time"

	"expense-management-api/dto"
	appErrors "expense-management-api/errors"
	"expense-management-api/models"
	"expense-management-api/repositories"
	"expense-management-api/utils"
)

type ExpenseService interface {
	CreateExpense(
		request dto.CreateExpenseRequest,
		userID int,
	) error

	GetExpensesByUserID(
		userID int,
		request dto.GetExpensesRequest,
	) ([]*models.Expense, error)
}

type expenseService struct {
	expenseRepository  repositories.ExpenseRepository
	categoryRepository repositories.CategoryRepository
}

func NewExpenseService(
	expenseRepository repositories.ExpenseRepository,
	categoryRepository repositories.CategoryRepository,
) ExpenseService {

	return &expenseService{
		expenseRepository:  expenseRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *expenseService) CreateExpense(
	request dto.CreateExpenseRequest,
	userID int,
) error {

	err := utils.ValidateCreateExpenseRequest(request)
	if err != nil {
		return err
	}

	category, err := s.categoryRepository.GetByID(
		request.CategoryID,
	)
	if err != nil {
		return err
	}

	if category == nil {
		return appErrors.ErrCategoryNotFound
	}

	if category.User.ID != userID {
		return appErrors.ErrForbiddenCategory
	}

	expenseDate, err := time.Parse(
		"2006-01-02",
		request.ExpenseDate,
	)
	if err != nil {
		return appErrors.ErrInvalidExpenseDate
	}

	expense := &models.Expense{
		Title:       request.Title,
		Amount:      request.Amount,
		Note:        request.Note,
		ExpenseDate: expenseDate,
		User: &models.User{
			ID: userID,
		},
		Category: &models.Category{
			ID: category.ID,
		},
	}

	return s.expenseRepository.Create(expense)
}

func (s *expenseService) GetExpensesByUserID(
	userID int,
	request dto.GetExpensesRequest,
) ([]*models.Expense, error) {

	err := utils.ValidateGetExpensesRequest(request)
	if err != nil {
		return nil, err
	}

	filter := repositories.ExpenseFilter{
		CategoryID: request.CategoryID,
		FromDate:   request.FromDate,
		ToDate:     request.ToDate,
		Page:       request.Page,
		Limit:      request.Limit,
		SortBy:     request.SortBy,
		SortOrder:  request.SortOrder,
	}

	return s.expenseRepository.GetExpenses(
		userID,
		filter,
	)
}
