package initializers

import "tgorm/db"

func DataBaseInitMigrate() {
	db.DataBaseMigrate()
}
