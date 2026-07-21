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
	LoginUser(request dto.LoginRequest) (string, error)
	GetByID(
		userID int,
	) (*models.User, error)
	UpdateUser(
		userID int,
		request dto.UpdateUserRequest,
	) error
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

func (s *userService) LoginUser(
	request dto.LoginRequest,
) (string, error) {

	err := utils.ValidateLoginRequest(request)
	if err != nil {
		return "", err
	}

	user, err := s.userRepository.GetByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", appErrors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)
	if err != nil {
		return "", appErrors.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetByID(
	userID int,
) (*models.User, error) {

	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, appErrors.ErrUserNotFound
	}

	return user, nil
}

func (s *userService) UpdateUser(
	userID int,
	request dto.UpdateUserRequest,
) error {

	err := utils.ValidateUpdateUserRequest(request)
	if err != nil {
		return err
	}

	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return appErrors.ErrUserNotFound
	}

	if user.Email != request.Email {

		existingUser, err := s.userRepository.GetByEmail(request.Email)
		if err != nil {
			return err
		}

		if existingUser != nil {
			return appErrors.ErrEmailExists
		}
	}

	user.Name = request.Name
	user.Email = request.Email

	return s.userRepository.Update(user)
}
