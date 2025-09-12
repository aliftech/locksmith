package mysql

import (
	"github.com/aliftech/locksmith/internal/core/app/models"
	"github.com/aliftech/locksmith/internal/core/app/repositories/interfaces"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.UserAuth {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Signup(user *models.User) error {
	return r.db.Create(user).Error
}
