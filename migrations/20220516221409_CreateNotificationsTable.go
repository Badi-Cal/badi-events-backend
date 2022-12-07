package main

import (
	"badi-cal/badi-events-backend/orm"

	"gorm.io/gorm"
)

var migration20220516221409_CreateNotificationsTable = Migration{
	version: "20220516221409",
	up: func(db *gorm.DB) {
		db.Migrator().CreateTable(&orm.Notification{})
	},
	down: func(db *gorm.DB) {
		db.Migrator().DropTable(&orm.Notification{})
	},
}

func init() {
	registerMigration(migration20220516221409_CreateNotificationsTable)
}
