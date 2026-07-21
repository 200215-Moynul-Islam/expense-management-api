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
