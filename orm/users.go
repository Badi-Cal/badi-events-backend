package orm

import (
	"badi-cal/badi-events-backend/util"
	"time"
)

type User struct {
	ID               uint `gorm:"primaryKey" json:"id"`
	Email            string
	Name             string
	GoogleId         string
	GoogleToken      string
	FromEmail        string
	IsSharedCalendar int
	TwitterId        string
	TwitterToken     string
	InstagramId      string
	InstagramToken   string
	WhatsAppId       string
	WhatAppToken     string
	Timezone         string
	FirstDayOfWeek   util.FirstDayOfWeek

	CreatedAt time.Time // magic gorm name
	UpdatedAt time.Time // magic gorm name
}
