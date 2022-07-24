package controllers

import (
	"badi-cal/badi-events-backend/models"
	"fmt"
	"net/http"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}

type App struct {
	Models *models.Models
}
type Handler func(w http.ResponseWriter, r *http.Request, app *App)
type HandlerDict map[string]Handler

var Routes = map[string]HandlerDict{
	"/notifications": {
		http.MethodGet: NotificationsIndex,
		http.MethodPut: NotificationsCreate,
	},
}

func (app *App) Route(w http.ResponseWriter, r *http.Request) {
	_, pathPresent := Routes[r.RequestURI]
	if pathPresent {
		handler, methodPresent := Routes[r.RequestURI][r.Method]
		if methodPresent {
			handler(w, r, app)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
