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

	GetCategoriesByUserID(
		userID int,
	) ([]*models.Category, error)

	GetCategoryByID(
		id int,
		userID int,
	) (*models.Category, error)

	UpdateCategory(
		id int,
		userID int,
		request dto.UpdateCategoryRequest,
	) error
	DeleteCategory(
		id int,
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

func (s *categoryService) GetCategoriesByUserID(
	userID int,
) ([]*models.Category, error) {

	return s.categoryRepository.GetAllByUserID(userID)
}

func (s *categoryService) GetCategoryByID(
	id int,
	userID int,
) (*models.Category, error) {

	category, err := s.categoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, appErrors.ErrCategoryNotFound
	}

	if category.User.ID != userID {
		return nil, appErrors.ErrForbiddenCategory
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(
	id int,
	userID int,
	request dto.UpdateCategoryRequest,
) error {

	err := utils.ValidateUpdateCategoryRequest(request)
	if err != nil {
		return err
	}

	category, err := s.categoryRepository.GetByID(id)
	if err != nil {
		return err
	}

	if category == nil {
		return appErrors.ErrCategoryNotFound
	}

	if category.User.ID != userID {
		return appErrors.ErrForbiddenCategory
	}

	existingCategory, err := s.categoryRepository.GetByNameAndUserID(
		request.Name,
		userID,
	)
	if err != nil {
		return err
	}

	if existingCategory != nil && existingCategory.ID != category.ID {
		return appErrors.ErrCategoryExists
	}

	category.Name = request.Name

	return s.categoryRepository.Update(category)
}

func (s *categoryService) DeleteCategory(
	id int,
	userID int,
) error {

	category, err := s.categoryRepository.GetByID(id)
	if err != nil {
		return err
	}

	if category == nil {
		return appErrors.ErrCategoryNotFound
	}

	if category.User.ID != userID {
		return appErrors.ErrForbiddenCategory
	}

	return s.categoryRepository.Delete(category)
}
