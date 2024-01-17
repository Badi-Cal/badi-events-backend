package controllers

import (
	"badi-cal/badi-events-backend/mock_models"
	"badi-cal/badi-events-backend/models"
	"badi-cal/badi-events-backend/orm"
	"badi-cal/badi-events-backend/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsersCreate(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := CreateMockApp(ctrl)
	now := time.Now()
	m := mock_models.NewMockIUsers(ctrl)
	m.
		EXPECT().
		Create(models.UserCreatePayload{
			Email:            "bob@bob.com",
			Name:             "Bob Bobberson",
			GoogleId:         "googleBob",
			GoogleToken:      "goobleBobToken",
			FromEmail:        "from@bob.com",
			IsSharedCalendar: 0,
			TwitterId:        "bobTwitter",
			TwitterToken:     "bobTwitterToken",
			InstagramId:      "bobInstagram",
			InstagramToken:   "bobInstagramToken",
			WhatsAppId:       "bobWhatsApp",
			WhatAppToken:     "bobWhatsAppToken",
			Timezone:         "America/New_York",
			FirstDayOfWeek:   "Monday",
		}).
		Return(&orm.User{
			ID:               1,
			Email:            "bob@bob.com",
			Name:             "Bob Bobberson",
			GoogleId:         "googleBob",
			GoogleToken:      "goobleBobToken",
			FromEmail:        "from@bob.com",
			IsSharedCalendar: 0,
			TwitterId:        "bobTwitter",
			TwitterToken:     "bobTwitterToken",
			InstagramId:      "bobInstagram",
			InstagramToken:   "bobInstagramToken",
			WhatsAppId:       "bobWhatsApp",
			WhatAppToken:     "bobWhatsAppToken",
			Timezone:         "America/New_York",
			FirstDayOfWeek:   util.FirstDayOfWeekFromString("Monday"),
			CreatedAt:        now,
			UpdatedAt:        now,
		}, nil).
		AnyTimes()
	app.Models.Users = m

	bodyReader := strings.NewReader(
		`{
			"email": "bob@bob.com", 
			"name": "Bob Bobberson",
			"google_id": "googleBob", 
			"google_token": "goobleBobToken", 
			"from_email": "from@bob.com", 
			"is_shared_calendar": 0, 
			"twitter_id": "bobTwitter", 
			"twitter_token": "bobTwitterToken", 
			"instagram_id": "bobInstagram", 
			"instagram_token": "bobInstagramToken", 
			"whats_app_id": "bobWhatsApp", 
			"whats_app_token": "bobWhatsAppToken", 
			"timezone": "America/New_York", 
			"first_day_of_week": "Monday"
		}`,
	)
	r := httptest.NewRequest(
		http.MethodPost,
		"/users",
		bodyReader,
	)
	w := httptest.NewRecorder()
	app.Route(w, r)

	assert.Equal(w.Code, http.StatusOK)
	assert.Equal(w.Body.String(), `{"id":1,"email":"bob@bob.com","name":"Bob Bobberson","google_id":"googleBob","google_token":"goobleBobToken","from_email":"from@bob.com","is_shared_calendar":0,"twitter_id":"bobTwitter","twitter_token":"bobTwitterToken","instagram_id":"bobInstagram","instagram_token":"bobInstagramToken","whats_app_id":"bobWhatsApp","whats_app_token":"bobWhatsAppToken","timezone":"America/New_York","first_day_of_week":"Monday"}`)
}

func TestUsersIndex(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := CreateMockApp(ctrl)
	now := time.Now()
	records := []*orm.User{{
		ID:               1,
		Email:            "bob@bob.com",
		Name:             "Bob Bobberson",
		GoogleId:         "googleBob",
		GoogleToken:      "goobleBobToken",
		FromEmail:        "from@bob.com",
		IsSharedCalendar: 0,
		TwitterId:        "bobTwitter",
		TwitterToken:     "bobTwitterToken",
		InstagramId:      "bobInstagram",
		InstagramToken:   "bobInstagramToken",
		WhatsAppId:       "bobWhatsApp",
		WhatAppToken:     "bobWhatsAppToken",
		Timezone:         "America/New_York",
		FirstDayOfWeek:   util.FirstDayOfWeekFromString("Monday"),
		CreatedAt:        now,
		UpdatedAt:        now,
	}}
	m := mock_models.NewMockIUsers(ctrl)
	m.
		EXPECT().
		GetAll().
		Return(records, nil).
		AnyTimes()
	app.Models.Users = m

	r := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)
	w := httptest.NewRecorder()
	app.Route(w, r)

	assert.Equal(w.Code, http.StatusOK)
	assert.Equal(w.Body.String(), `[{"id":1,"email":"bob@bob.com","name":"Bob Bobberson","google_id":"googleBob","google_token":"goobleBobToken","from_email":"from@bob.com","is_shared_calendar":0,"twitter_id":"bobTwitter","twitter_token":"bobTwitterToken","instagram_id":"bobInstagram","instagram_token":"bobInstagramToken","whats_app_id":"bobWhatsApp","whats_app_token":"bobWhatsAppToken","timezone":"America/New_York","first_day_of_week":"Monday"}]`)
}
