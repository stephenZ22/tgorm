package main

import (
	"fmt"
	"tgorm/db"
	"tgorm/initializers"
	"tgorm/models"
	"tgorm/services"

	"github.com/jaswdr/faker"
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
	// 	Name:   "王强",
	// 	Age:    34,
	// 	Gender: 1,
	// 	Email:  "wangqiang@tgorm.com",
	// 	UserInfo: models.UserInfo{
	// 		School:  "北京大学",
	// 		Address: "北京市",
	// 	},
	// 	UserBankCard: []models.UserBankCard{
	// 		{Number: "0011234567", BankName: "工商银行"},
	// 		{Number: "0022345678", BankName: "农业银行"},
	// 		{Number: "0022343345", BankName: "农业银行"},
	// 	},
	// }
	// users := []*models.User{}
	// users = append(users, user3)

	// users = services.MultipleCreateUser(users)
	// fmt.Println(users)

	fmt.Println(services.GetInfoSchoolByUserName("王强"))
	cards := services.GetBankCardsByUserName("王强")
	for _, card := range cards {
		fmt.Println(card.BankName, card.Number)
	}

}

func createUsers() {
	users := make([]*models.User, 200, 200)
	p := faker.New()
	for i := 0; i < 100; i++ {
		username := p.Person().Name()
		email := p.Person().Contact().Email

		user := &models.User{
			Name:   username,
			Email:  email,
			Age:    uint(23 + i),
			Gender: 1,
			UserInfo: models.UserInfo{
				School:  p.Person().Faker.Address().Address(),
				Address: p.Person().Faker.Address().Address() + "School",
			},
			UserBankCard: []models.UserBankCard{
				{Number: p.Person().Faker.Payment().CreditCardNumber(), BankName: "工商银行"},
				{Number: p.Person().Faker.Payment().CreditCardNumber(), BankName: "农业银行"},
			},
		}

		db.Connection.Create(user)
		users = append(users, user)
	}
}
