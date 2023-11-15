package orm

import "time"

type FirstDayOfWeek int

const (
	Sunday FirstDayOfWeek = 0
	Monday                = 1
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
	FirstDayOfWeek   FirstDayOfWeek

	CreatedAt time.Time // magic gorm name
	UpdatedAt time.Time // magic gorm name
}
