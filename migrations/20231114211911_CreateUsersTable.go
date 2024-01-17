package main

import (
	"badi-cal/badi-events-backend/orm"

	"gorm.io/gorm"
)

var migration20231114211911_CreateUsersTable = Migration{
	name:    "CreateUsersTable",
	version: "20231114211911",
	up: func(db *gorm.DB) {
		db.Migrator().CreateTable(&orm.User{})
	},
	down: func(db *gorm.DB) {
		db.Migrator().DropTable(&orm.User{})
	},
}

func init() {
	registerMigration(migration20231114211911_CreateUsersTable)
}
