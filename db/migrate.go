package db

import "tgorm/models"

func DataBaseMigrate() {
	Connection.AutoMigrate(&models.User{})
	Connection.AutoMigrate(&models.UserInfo{})
}
