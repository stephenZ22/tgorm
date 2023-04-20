package initializers

import "tgorm/db"

func ConnectDb() {
	db.DataBaseConnection()
}
