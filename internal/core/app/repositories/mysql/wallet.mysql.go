package mysql

import (
	"github.com/aliftech/locksmith/internal/core/app/models"
	"github.com/aliftech/locksmith/internal/core/app/repositories/interfaces"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) interfaces.WalletInterface {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) StoreWallet(wallet *models.Wallet) error {
	return r.db.Create(wallet).Error
}
