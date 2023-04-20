package main

import (
	"fmt"
	"tgorm/initializers"
	"tgorm/services"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.DataBaseInitMigrate()
}

func main() {
	fmt.Println("Hello Test Gorm 2222 ~~~")

	user := services.CreateUser("李明", "liming3@tgorm.com", 22, 1)
	fmt.Println(user.GetName(), user.GetGender())
}
