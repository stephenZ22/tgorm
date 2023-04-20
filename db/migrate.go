package db

import "tgorm/models"

func DataBaseMigrate() {
	DB.AutoMigrate(&models.User{})
}
