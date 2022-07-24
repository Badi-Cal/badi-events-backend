package models

import (
	"badi-cal/badi-events-backend/orm"
	"time"

	"gorm.io/gorm"
)

type INotifications interface {
	Create(payload NotificationCreatePayload) (*orm.Notification, error)
	GetAll() ([]*orm.Notification, error)
}

type Notifications struct {
	db *gorm.DB
}

type NotificationCreatePayload struct {
	Event_id       string `json:"event_id"`
	Body           string `json:"body"`
	Scheduled_time string `json:"scheduled_time"` // JS Date#toISOString
}

func MakeNotifications(db *gorm.DB) *Notifications {
	return &Notifications{
		db: db,
	}
}

func (n *Notifications) Create(payload NotificationCreatePayload) (*orm.Notification, error) {
	scheduled_time, _ := time.Parse(time.RFC3339, payload.Scheduled_time)
	notification := &orm.Notification{
		Event_id:       payload.Event_id,
		Body:           payload.Body,
		Scheduled_time: orm.JSONTime(scheduled_time),
		Status:         orm.Unsent,
		Attempts:       0,
	}
	result := n.db.Create(notification)
	if result.Error != nil {
		return nil, result.Error
	}

	return notification, nil
}

func (n *Notifications) GetAll() ([]*orm.Notification, error) {
	notifications := []*orm.Notification{}
	results := n.db.Find(&notifications)
	if results.Error != nil {
		return nil, results.Error
	}

	return notifications, nil
}
