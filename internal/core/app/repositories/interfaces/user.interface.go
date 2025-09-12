package interfaces

import "github.com/aliftech/locksmith/internal/core/app/models"

type UserInterface interface {
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
}
