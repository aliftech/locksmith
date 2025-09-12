package mysql

import (
	"github.com/aliftech/locksmith/internal/core/app/models"
	"github.com/aliftech/locksmith/internal/core/app/repositories/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindUserById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return &user, err
}
