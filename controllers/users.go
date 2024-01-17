package controllers

import (
	"badi-cal/badi-events-backend/models"
	"badi-cal/badi-events-backend/orm"
	"badi-cal/badi-events-backend/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserPayload struct {
	ID               uint   `json:"id"`
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

func toUserPayload(user orm.User) UserPayload {
	return UserPayload{
		ID:               user.ID,
		Email:            user.Email,
		Name:             user.Name,
		GoogleId:         user.GoogleId,
		GoogleToken:      user.GoogleToken,
		FromEmail:        user.FromEmail,
		IsSharedCalendar: user.IsSharedCalendar,
		TwitterId:        user.TwitterId,
		TwitterToken:     user.TwitterToken,
		InstagramId:      user.InstagramId,
		InstagramToken:   user.InstagramToken,
		WhatsAppId:       user.WhatsAppId,
		WhatAppToken:     user.WhatAppToken,
		Timezone:         user.Timezone,
		FirstDayOfWeek:   util.FirstDayOfWeekString(user.FirstDayOfWeek),
	}
}

func UsersCreate(w http.ResponseWriter, r *http.Request, app *App) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Body: %v", string(body))
	var data models.UserCreatePayload
	json.Unmarshal([]byte(body), &data)
	user, err0 := app.Models.Users.Create(data)
	if err0 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err0))
		return
	}
	payload := toUserPayload(*user)

	b, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
}

func UsersIndex(w http.ResponseWriter, r *http.Request, app *App) {
	users, err0 := app.Models.Users.GetAll()
	if err0 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err0))
		return
	}
	payload := []UserPayload{}
	for _, user := range users {
		payload = append(payload, toUserPayload(*user))
	}

	b, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
}
