package orm

import "time"

type NotificationStatus int

const (
	Unsent NotificationStatus = 0
	Sent                      = 1
	Failed                    = 2
)

type Notification struct {
	ID             uint `gorm:"primaryKey" json:"id"`
	Event_id       string
	Body           string
	Scheduled_time time.Time
	Status         NotificationStatus
	Attempts       int
	CreatedAt      time.Time // magic gorm name
	UpdatedAt      time.Time // magic gorm name
}
