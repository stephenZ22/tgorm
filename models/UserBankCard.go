package models

import "gorm.io/gorm"

type UserBankCard struct {
	gorm.Model
	UserID   uint
	Number   string `gorm:"unique"`
	BankName string
	User     User
}
