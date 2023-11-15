package models

import (
	"badi-cal/badi-events-backend/orm"

	"gorm.io/gorm"
)

type IUsers interface {
	Create(payload UserCreatePayload) (*orm.User, error)
	GetAll() ([]*orm.User, error)
}

type Users struct {
	db *gorm.DB
}

type UserCreatePayload struct {
	Email            string `json:"email"`
	Name             string `json:"name"`
	GoogleId         string `json:"google_id"`
	GoogleToken      string `json:"google_token"`
	FromEmail        string `json:"from_email"`
	IsSharedCalendar int    `json:"is_shared_calendar"`
	TwitterId        string `json:"twitter_id"`
	TwitterToken     string `json:"twitter_token"`
	InstagramId      string `json:"instagram_id"`
	InstagramToken   string `json:"instagram_token"`
	WhatsAppId       string `json:"whats_app_id"`
	WhatAppToken     string `json:"whats_app_token"`
	Timezone         string `json:"timezone"`
	FirstDayOfWeek   string `json:"first_day_of_week"`
}

var firstDayOfWeekMap = map[orm.FirstDayOfWeek]string{
	0: "Sunday",
	1: "Monday",
}

func FirstDayOfWeekString(dayOfWeek orm.FirstDayOfWeek) string {
	return firstDayOfWeekMap[dayOfWeek]
}

func FirstDayOfWeekFromString(str string) orm.FirstDayOfWeek {
	for key, val := range firstDayOfWeekMap {
		if val == str {
			return key
		}
	}

	return orm.Sunday
}

func MakeUsers(db *gorm.DB) *Users {
	return &Users{
		db: db,
	}
}

func (n *Users) Create(payload UserCreatePayload) (*orm.User, error) {
	user := &orm.User{
		Email:            payload.Email,
		Name:             payload.Name,
		GoogleId:         payload.GoogleId,
		GoogleToken:      payload.GoogleToken,
		FromEmail:        payload.FromEmail,
		IsSharedCalendar: payload.IsSharedCalendar,
		TwitterId:        payload.TwitterId,
		TwitterToken:     payload.TwitterToken,
		InstagramId:      payload.InstagramId,
		InstagramToken:   payload.InstagramToken,
		WhatsAppId:       payload.WhatsAppId,
		WhatAppToken:     payload.WhatAppToken,
		Timezone:         payload.Timezone,
		FirstDayOfWeek:   FirstDayOfWeekFromString(payload.FirstDayOfWeek),
	}
	result := n.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (n *Users) GetAll() ([]*orm.User, error) {
	users := []*orm.User{}
	results := n.db.Find(&users)
	if results.Error != nil {
		return nil, results.Error
	}

	return users, nil
}
