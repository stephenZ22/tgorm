package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func DataBaseConnection() {
	var err error

	dsn := os.Getenv("DataBaseDsn")
	Connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

}
