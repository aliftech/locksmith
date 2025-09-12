package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Firstname string `gorm:"type:varchar(255); not null" json:"firstname"`
	Lastname  string `gorm:"type:varchar(255); not null;" json:"lastname"`
	Email     string `gorm:"type:varchar(255); not null" json:"email"`
	Password  string `gorm:"type:varchar(255); not null" json:"password"`
}
