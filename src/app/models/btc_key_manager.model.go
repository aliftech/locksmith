package models

import "gorm.io/gorm"

type BTCKeyManager struct {
	gorm.Model
	PublicKey  string
	PrivateKey string
}
