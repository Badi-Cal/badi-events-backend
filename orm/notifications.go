package orm

import "time"

type NotificationStatus int

const (
	Unsent NotificationStatus = 0
	Sent                      = 1
	Failed                    = 2
)

type Notification struct {
	ID             uint `gorm:"primaryKey"`
	event_id       string
	body           string
	scheduled_time time.Time
	status         NotificationStatus
	attempts       int
	CreatedAt      time.Time // magic gorm name
	UpdatedAt      time.Time // magic gorm name
}
