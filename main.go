package main

import (
	"fmt"
	"tgorm/initializers"
	"tgorm/models"
	"tgorm/services"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.DataBaseInitMigrate()
}

func main() {
	fmt.Println("Hello Test Gorm 2222 ~~~")

	// user := services.CreateUser("王小明", "wangxiaoming@tgorm.com", 27, 1)
	// users := createUsers()
	// for _, user := range users {
	// 	fmt.Println(user.GetName(), user.GetGender())
	// }

	// user := services.FindUserByName("黎明")
	// fmt.Println(user.GetName(), user.GetGender())

	// user2 := services.FindUserById(3)
	// fmt.Println(user2.GetName(), user2.GetGender())
	// user3 := &models.User{
	// 	Name:   "雷军",
	// 	Age:    59,
	// 	Gender: 1,
	// 	Email:  "leijun@tgorm.com",
	// 	UserInfo: models.UserInfo{
	// 		School:  "武汉大学",
	// 		Address: "湖北省",
	// 	},
	// }
	// users := []*models.User{}
	// users = append(users, user3)

	// users = services.MultipleCreateUser(users)
	// fmt.Println(user)
	fmt.Println(services.GetInfoSchoolByUserName("雷军"))

}

func createUsers() []*models.User {
	users := []*models.User{
		{Name: "刘德华", Age: 61, Gender: 1, Email: "liudehua@tgorm.com"},
		{Name: "张学友", Age: 62, Gender: 1, Email: "zhangxueyou@tgorm.com"},
		{Name: "郭富城", Age: 60, Gender: 1, Email: "guofucheng@tgorm.com"},
		{Name: "黎明", Age: 63, Gender: 1, Email: "liminghk@tgorm.com"},
	}

	newUsers := services.MultipleCreateUser(users)
	return newUsers
}
