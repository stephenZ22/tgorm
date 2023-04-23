package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Address string
	School  string
	UserId  uint
}
