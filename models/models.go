package models

import (
	"gorm.io/gorm"
)

type Models struct {
	Notifications INotifications
	Users         IUsers
}

func GetModels(db *gorm.DB) *Models {
	return &Models{
		Notifications: MakeNotifications(db),
		Users:         MakeUsers(db),
	}
}
