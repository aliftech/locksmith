package interfaces

import "github.com/aliftech/locksmith/internal/core/app/models"

type UserAuth interface {
	// User signup interface
	Signup(user *models.User) error
}
