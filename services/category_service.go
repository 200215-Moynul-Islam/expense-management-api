package services

import (
	"expense-management-api/dto"
	appErrors "expense-management-api/errors"
	"expense-management-api/models"
	"expense-management-api/repositories"
	"expense-management-api/utils"
)

type CategoryService interface {
	CreateCategory(
		request dto.CreateCategoryRequest,
		userID int,
	) error
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(
	categoryRepository repositories.CategoryRepository,
) CategoryService {

	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *categoryService) CreateCategory(
	request dto.CreateCategoryRequest,
	userID int,
) error {

	err := utils.ValidateCreateCategoryRequest(request)
	if err != nil {
		return err
	}

	existingCategory, err := s.categoryRepository.GetByNameAndUserID(
		request.Name,
		userID,
	)

	if err != nil {
		return err
	}

	if existingCategory != nil {
		return appErrors.ErrCategoryExists
	}

	category := &models.Category{
		Name: request.Name,
		User: &models.User{
			ID: userID,
		},
	}

	return s.categoryRepository.Create(category)
}
