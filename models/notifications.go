package models

import (
	"badi-cal/badi-events-backend/orm"

	"gorm.io/gorm"
)

type INotifications interface {
	GetAll() ([]*orm.Notification, error)
}

type Notifications struct {
	db *gorm.DB
}

func MakeNotifications(db *gorm.DB) *Notifications {
	return &Notifications{
		db: db,
	}
}

func (n *Notifications) GetAll() ([]*orm.Notification, error) {
	notifications := []*orm.Notification{}
	results := n.db.Find(&notifications)
	if results.Error != nil {
		return nil, results.Error
	}

	return notifications, nil
}
