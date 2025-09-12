package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	Ticker          string `gorm:"type:varchar(30); not null" json:"ticker"`
	Mnemonic        string `gorm:"type:text; not null" json:"mnemonic"`
	PublicKey       string `gorm:"type:varchar(255); not null" json:"public_key"`
	PrivateKey      string `gorm:"type:text; not null" json:"private_key"`
	Address         string `gorm:"type:varchar(255); not null" json:"address"`
	DerivationIndex uint   `gorm:"type:int; not null" json:"derivation_index"`
	PassphraseHash  string `gorm:"type:varchar(64); not null" json:"passphrase_hash"`
	CreatedBy       string `gorm:"type:varchar(255); not null" json:"created_by"`
}
