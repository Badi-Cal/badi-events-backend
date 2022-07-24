package controllers

import (
	"badi-cal/badi-events-backend/models"
	"badi-cal/badi-events-backend/orm"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NotificationPayload struct {
	ID             uint     `json:"id"`
	Event_id       string   `json:"event_id"`
	Body           string   `json:"body"`
	Scheduled_time JSONTime `json:"scheduled_time"`
	Status         string   `json:"status"`
	Attempts       int      `json:"attempts"`
}

var statusMap = map[int]string{
	0: "unsent",
	1: "sent",
	2: "failed",
}

func StatusString(status orm.NotificationStatus) string {
	return statusMap[int(status)]
}

func toNotificationPayload(notification orm.Notification) NotificationPayload {
	return NotificationPayload{
		ID:             notification.ID,
		Event_id:       notification.Event_id,
		Body:           notification.Body,
		Scheduled_time: JSONTime(notification.Scheduled_time),
		Status:         StatusString(notification.Status),
		Attempts:       notification.Attempts,
	}
}

func NotificationsCreate(w http.ResponseWriter, r *http.Request, app *App) {
	body, _ := ioutil.ReadAll(r.Body)
	var data models.NotificationCreatePayload
	json.Unmarshal([]byte(body), &data)
	b, _ := json.Marshal(data)
	notification, err0 := app.Models.Notifications.Create(data)
	if err0 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err0))
		return
	}
	payload := toNotificationPayload(*notification)

	b, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
}

func NotificationsIndex(w http.ResponseWriter, r *http.Request, app *App) {
	notifications, err0 := app.Models.Notifications.GetAll()
	if err0 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err0))
		return
	}
	payload := []NotificationPayload{}
	for _, notification := range notifications {
		payload = append(payload, toNotificationPayload(*notification))
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
