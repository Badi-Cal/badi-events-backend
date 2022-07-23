package models

import (
	"gorm.io/gorm"
)

type Models struct {
	Notifications INotifications
}

func GetModels(db *gorm.DB) *Models {
	return &Models{
		Notifications: MakeNotifications(db),
	}
}
