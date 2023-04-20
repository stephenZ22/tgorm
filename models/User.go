package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Email  string `gorm:"unique"`
	Age    uint   `gorm:"default:1"`
	Gender uint   `gorm:"default:0"`
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetGender() (gender string) {
	if u.Gender == 1 {
		gender = "男"
	} else {
		gender = "女"
	}

	return
}
