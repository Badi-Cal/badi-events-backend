package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func NotificationsCreate(w http.ResponseWriter, r *http.Request) {
// 	notifications, err0 := orm.NotificationGetAll()
// 	if err0 != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, fmt.Sprintf("%v", err0))
// 		return
// 	}

// 	b, err := json.Marshal(notifications)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, fmt.Sprintf("%v", err))
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, string(b))
// }

func NotificationsIndex(w http.ResponseWriter, r *http.Request, app *App) {
	notifications, err0 := app.Models.Notifications.GetAll()
	if err0 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err0))
		return
	}

	b, err := json.Marshal(notifications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("%v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
}
