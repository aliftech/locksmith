package migrations

import (
	"github.com/aliftech/locksmith/src/app/models"
	"github.com/aliftech/locksmith/src/config"
)

func BtcKeyMigration() {
	config.DB.AutoMigrate(&models.BTCKeyManager{})
}
