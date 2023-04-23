package services

import (
	"log"
	"tgorm/db"
	"tgorm/models"
)

//type User models.User

func CreateUser(name string, email string, age uint, gender uint) models.User {

	user := models.User{Name: name, Email: email, Age: age, Gender: gender}
	result := db.Connection.Create(&user)

	if result.Error != nil {
		log.Fatal("Create user error")
	}

	return user
}

func MultipleCreateUser(users []*models.User) []*models.User {
	result := db.Connection.Create(users)

	if result.Error != nil {
		log.Fatal("Create users error")
	}

	return users
}

func FindUserById(id int) (user models.User) {
	db.Connection.First(&user, id)

	return
}

func FindUserByName(name string) (user models.User) {
	db.Connection.Where("name = ?", name).First(&user)

	return
}

func GetInfoSchoolByUserName(n string) string {
	var info models.UserInfo
	user := FindUserByName(n)

	db.Connection.Model(&user).Association("UserInfo").Find(&info)
	return info.School
}

func GetBankCardsByUserName(n string) (cards []models.UserBankCard) {

	user := FindUserByName(n)

	db.Connection.Model(&user).Association("UserBankCard").Find(&cards)
	return
}
