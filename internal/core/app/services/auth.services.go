package services

import (
	"errors"

	"github.com/aliftech/locksmith/internal/core/app/dto"
	"github.com/aliftech/locksmith/internal/core/app/models"
	"github.com/aliftech/locksmith/internal/core/app/repositories/interfaces"
	"github.com/aliftech/locksmith/internal/core/helpers"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo     interfaces.UserInterface
	UserAuthRepo interfaces.UserAuth
	Validate     *validator.Validate
}

func NewAuthService(userRepo interfaces.UserInterface, userAuth interfaces.UserAuth) *AuthService {
	return &AuthService{
		UserRepo:     userRepo,
		UserAuthRepo: userAuth,
		Validate:     validator.New(),
	}
}

func (s *AuthService) Signup(input dto.SignupreqBody) (*models.User, error) {
	// Validate input
	if err := s.Validate.Struct(input); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if email exists
	existingUser, _ := s.UserRepo.FindUserByEmail(input.Email)
	if existingUser.Email != "" {
		return nil, errors.New(helpers.Message("email_exist", ""))
	}

	// Hash password
	hashedPassword, err := s.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Create domain model
	user := &models.User{
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  hashedPassword,
	}

	// Persist user
	if err := s.UserAuthRepo.Signup(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
