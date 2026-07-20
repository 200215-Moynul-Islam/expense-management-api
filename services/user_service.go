package services

import (
	"expense-management-api/dto"
	appErrors "expense-management-api/errors"
	"expense-management-api/models"
	"expense-management-api/repositories"
	"expense-management-api/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(request dto.RegisterRequest) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) RegisterUser(
	request dto.RegisterRequest,
) error {

	err := utils.ValidateRegisterRequest(request)
	if err != nil {
		return err
	}

	existingUser, err := s.userRepository.GetByEmail(request.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return appErrors.ErrEmailExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passwordHash),
	}

	return s.userRepository.Create(user)
}
