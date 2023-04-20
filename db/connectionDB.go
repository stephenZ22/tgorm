package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBaseConnection() {
	var err error

	dsn := os.Getenv("DataBaseDsn")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

}

func CurrentDB() {

}
