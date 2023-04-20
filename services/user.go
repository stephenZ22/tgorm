package services

import (
	"log"
	"tgorm/db"
	"tgorm/models"
)

//type User models.User

func CreateUser(name string, email string, age uint, gender uint) *models.User {

	user := &models.User{Name: name, Email: email, Age: age, Gender: gender}
	result := db.DB.Create(user)

	if result.Error != nil {
		log.Fatal("Create user error")
	}

	return user
}
