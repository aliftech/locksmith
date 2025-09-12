package interfaces

import "github.com/aliftech/locksmith/internal/core/app/models"

type WalletInterface interface {
	StoreWallet(wallet *models.Wallet) error
}
